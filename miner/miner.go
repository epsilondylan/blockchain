package main


import(
    "encoding/json"
    "fmt"
    "log"
    "blockchain/models"
    "net"

    p2p "blockchain/p2p"
    "google.golang.org/grpc"

)


func Mine(Server *p2p.P2P_Server){
    for{
        var transstring string
        //一次最多十条交易，也可以优化
        for i := 0; i < 10; i++ {
            select {

            case trans := <- Server.NewTrans:
                 temp,err:=json.Marshal(trans)
                 if err!=nil{
                    return
                 }
                transstring+=string(temp)
            default:
                break
            }
        }
        //生成区块
        block := models.GenerateBlock(models.GetChainTail().Hash, transstring, models.GetChainLen())
        err:=models.AppendChain(block)
        if err == nil{
            //广播区块
            Server.Broadcast(block)
            //将区块写入本地
            err1:=Server.SaveChainToLocal()
            if err1!=nil{
                fmt.Println(err1)
            }
        }
    }
}

func main(){
    //初始化
    Server:=new(p2p.P2P_Server)
    Server.Init()
    lis, err := net.Listen("tcp", ":1230")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    models.RegisterP2PServer(s, Server)
    go Server.SetupConnections()
    go Mine(Server)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

}

