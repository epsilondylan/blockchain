package p2p;

import (
    "encoding/json"
    "fmt"
    "strconv"
    "time"
    "context"
    "os"

    dhash "github.com/epsilondylan/blockchain/hash"

    "github.com/epsilondylan/blockchain/common"
    "github.com/epsilondylan/blockchain/models"
    pto "github.com/epsilondylan/blockchain/proto"
    "google.golang.org/grpc"
)

type P2P_Server struct {
    pto.UnimplementedP2PServer
    PeerList *pto.PeerList
    TransPool *pto.TransPool
    NewTrans  chan *pto.Trans
    clients []*pto.P2PClient
}

func (s *P2P_Server) init(){
    s.PeerList=new(pto.PeerList)
    s.TransPool=new(pto.TransPool)
    s.clients=make([]*pto.P2PClient,0)
    s.NewTrans=make(chan *pto.Trans,1000)
    //从本地读取区块链，没有的话创建新的区块链
    models.CreateChain()
    data,err:=os.ReadFile(./blockchain.json)
    if err!=nil{
        return
    }
    //解析区块链
    chain,err:=models.FormatChain(data)
    models.ReplaceChain(chain)
    if err!=nil{
        fmt.Println(err)
        return
    }
    //从区块链中解析交易池
    for _,block :=range chain{
        transpool,err:=models.FormatTrans(block.TransPool)
        if err!=nil{
            fmt.Println(err)
            return
        }
        //将交易池中的交易加入到交易池中
        s.TransPool.TransPool=append(s.TransPool.TransPool,transpool.TransPool...)
    }
}

func (s *P2P_Server) SaveChainToLocal() error {
    //将区块链转换成json格式
    data, err := json.Marshal(models.FetchChain())
    if err != nil {
        return err
    }
    //将json格式的区块链写入到本地
    err = os.WriteFile(./blockchain.json, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func (s *P2P_Server) RequestTailBlock(ctx context.Context, in *pto.TailBlockRequest) (*pto.Block, error) {
    return models.GetChainTail(), nil
}

//这里是UpdateBlockChain的接收方，目的是当server接受到了一个更长的链上的block，它去请求整条链，然后发送消息的server把整条链给它，这里是
//在处理接受到整条链的时候。
func (s *P2P_Server) UpdateBlockChain(ctx context.Context, in *pto.BlockChain) (*pto.UpdateBlockChainResponse, error) {
    dhash.StopHash()
    defer dhash.StartHash()
    err:=models.ReplaceChain(in)
    if err!=nil{
        return &pto.UpdateBlockChainResponse{},common.Error(common.ErrInvalidChain)
    }
    err=s.SaveChainToLocal()//同样很蠢，值得优化
    if err!=nil{
        return &pto.UpdateBlockChainResponse{},err
    }
    go s.Broadcast(models.GetChainTail())
    return &pto.UpdateBlockChainResponse{},nil
}

func (s *P2P_Server) NewBlock(ctx context.Context, in *pto.Block) (*pto.NewBlockResponse, error) {
    dhash.StopHash()
    defer dhash.StartHash()
    tailBlcok := models.GetChainTail()
    if *tailBlcok == *in {
        return &pto.NewBlockResponse{ChainNeedUpdate: false}, nil
    }
    if !in.IsTempValid() || in.Index <= tailBlcok.Index {
        return &pto.NewBlockResponse{ChainNeedUpdate: false}, common.Error(common.ErrInvalidBlock)
    }
    if in.IsVaild(tailBlcok) {
        err := models.AppendChain(in)
        if err != nil {
            return &pto.NewBlockResponse{ChainNeedUpdate: false}, err
        }
        go s.Broadcast(in)
        err = s.SaveChainToLocal()//每更新一个块就保存一次，非常蠢，值得优化
        if err != nil {
            return &pto.NewBlockResponse{ChainNeedUpdate: false}, err
        }
        return &pto.NewBlockResponse{ChainNeedUpdate: false}, nil
    }
    
    if in.Index > tailBlcok.Index+1 {
        return &pto.NewBlockResponse{ChainNeedUpdate: true}, nil
    }
    return &pto.NewBlockResponse{ChainNeedUpdate: false}, nil
}

func (s *P2P_Server) NewTransaction(ctx context.Context,in *pto.Trans)(*pto.NewTransactionResponse,error){
    //检验交易是否已在链上存在
    //非常蠢，值得优化
    for _,pvtrans:=range s.TransPool.TransPool{
        if pvtrans==in{
            return &pto.NewTransactionResponse{},nil
        }
    }
    //以及因为管理账户太麻烦了，这里也不检查合法，反正没要求演示
    s.NewTrans<-in
    return &pto.NewTransactionResponse{},nil
}

func (s *P2P_Server) ReadPeerListFromLocal() error {
    //读取以json格式存储的peerlist
    data, err := os.ReadFile(./peerlist.json)
    if err != nil {
        return err
    }
    //解析peerlist
    err = json.Unmarshal(data, s.PeerList)
    if err != nil {
        return err
    }
    return nil
}

func (s *P2P_Server) NewPeer(ctx context.Context,in *pto.Peer)(*pto.PeerList,error){
    //读取本地的peerlist
    err:=s.ReadPeerListFromLocal()
    if err!=nil{
        return nil,err
    }
    //检验重复
    for _,peer:=range s.PeerList.PeerList{
        if peer==in{
            return s.PeerList,nil
        }
    }
    //将新的peer加入到peerlist中
    s.PeerList.PeerList=append(s.PeerList.PeerList,in)
    //将peerlist写入到本地
    err=s.WritePeerListToLocal()
    if err!=nil{
        return nil,err
    }
    return s.PeerList,nil 
}

func (s *P2P_Server) WritePeerListToLocal() error {
    //将peerlist转换成json格式
    data, err := json.Marshal(s.PeerList)
    if err != nil {
        return err
    }
    //将json格式的peerlist写入到本地
    err = os.WriteFile(./peerlist.json, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func (s *P2P_Server) SetupConnections()  error {
    time.Sleep(5*time.Second)
    //读取本地的peerlist
    err:=s.ReadPeerListFromLocal()
    if err!=nil{
        return err
    }
    //建立连接
    for _,peer:=range s.PeerList{
        conn,err:=grpc.Dial(fmt.Sprintf("%s:%d",peer.Ip,peer.Port),grpc.WithInsecure())
        if err!=nil{
            continue
        }
        s.clients=append(s.clients,pto.NewP2PClient(conn))
    }
    //启动监控
    go s.monitorConnections()
    return nil
}
//监控连接是否还在，不在重新连接，是否算一种优化呢？
func (s *P2P_Server) monitorConnections() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    for {
        <-ticker.C
        for i, client := range s.clients {
            if conn, ok := client.(*grpc.ClientConn); ok {
                state := conn.GetState()
                if state != connectivity.Ready {
                    // 尝试重新连接
                    peer := s.PeerList[i]
                    newConn, err := grpc.Dial(fmt.Sprintf("%s:%d", peer.Ip, peer.Port), grpc.WithInsecure())
                    if err == nil {
                        s.clients[i] = pto.NewP2PClient(newConn)
                    }
                }
            }
        }
    }
}


func (s *P2P_Server) Broadcast(block *pto.Block) {
    for _, client := range s.clients {
        go func(client *pto.P2PClient) {
            res,err:=client.NewBlock(context.Background(),block)
            if err!=nil{
                fmt.Println(err)
            }
            if res.ChainNeedUpdate{
                _,err:=client.UpdateBlockChain(context.Background(),models.FetchChain())
                if err!=nil{
                    fmt.Println(err)
                }
            }
        }(client)
    }
}







