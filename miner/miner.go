package miner


import(
    "encoding/json"
    "fmt"
    "sync"
    "time"
    "context"
    "io/ioutil"

    dhash "github.com/epsilondylan/blockchain/hash"
    "github.com/epsilondylan/blockchain/common"
    "github.com/epsilondylan/blockchain/models"
    pto "github.com/epsilondylan/blockchain/proto"
    p2p "github.com/epsilondylan/blockchain/p2p"
    "google.golang.org/grpc"

)


func Mine(Server *p2p.P2P_Server){
    for{
        var transstring string
        //一次最多十条交易，也可以优化
        for i := 0; i < 10; i++ {
            case trans <- Server.NewTrans:
                 temp,err:=json.Marshal(trans)
                 if err!=nil{
                    return
                 }
                transstring+=string(temp)
            default:
                break
        }
        //生成区块
        block := models.GenerateBlock(models.GetChainTail().Hash, transStr, models.GetChainLen())
        err:=models.AppendChain(block)
        if err == nil{
            //广播区块
            Server.BroadcastBlock(block)
            //将区块写入本地
            err1:=SaveChainToLocal()
            if err1!=nil{
                fmt.Println(err1)
            }
        }
    }
}

func main(){
    //初始化
    Server:=new(p2p.P2P_Server)
    lis, err := net.Listen("tcp", ":1230")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pto.RegisterP2PServer(s, Server)
    go Server.SetupConnections()
    go Mine(Server)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

}

