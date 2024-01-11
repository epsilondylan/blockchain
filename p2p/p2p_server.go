package p2p;

import (
    "encoding/json"
    "fmt"
    "strconv"
    "time"
    "context"
    "io/ioutil"

    dhash "github.com/epsilondylan/blockchain/hash"

    "github.com/epsilondylan/blockchain/common"
    "github.com/epsilondylan/blockchain/models"
    pto "github.com/epsilondylan/blockchain/proto"
)

type P2P_Server struct {
    pto.UnimplementedP2PServer
    PeerList *pto.PeerList
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
        return &pto.NewBlockResponse{ChainNeedUpdate: false}, nil
    }
    
    if in.Index > tailBlcok.Index+1 {
        return &pto.NewBlockResponse{ChainNeedUpdate: true}, nil
    }
    return &pto.NewBlockResponse{ChainNeedUpdate: false}, nil
}

func (s *P2P_Server) NewTransaction(ctx context.Context,in *pto.Trans)(*pto.NewTransactionResponse,error){
    return &pto.NewTransactionResponse{},nil
}

func (s *P2P_Server) ReadPeerListFromLocal() error {
    //读取以json格式存储的peerlist
    data, err := ioutil.ReadFile(./peerlist.json)
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
    err = ioutil.WriteFile(./peerlist.json, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func (s *P2P_Server) SetupConnections() ([]*pto.P2PClient, error) {
    //读取本地的peerlist
    err:=s.ReadPeerListFromLocal()
    if err!=nil{
        return nil,err
    }
    //建立连接
    var clients []*pto.P2PClient
    for _,peer:=range s.PeerList.PeerList{
        conn,err:=grpc.Dial(fmt.Sprintf("%s:%d",peer.Ip,peer.Port),grpc.WithInsecure())
        if err!=nil{
            continue
        }
        clients=append(clients,pto.NewP2PClient(conn))
    }
    return clients,nil
}

func (s *P2P_Server) Broadcast(block *pto.Block, clients []*pto.P2PClient) {
    for _, client := range clients {
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







