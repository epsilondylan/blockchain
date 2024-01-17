package main
 
import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"

    "blockchain/models"
    "google.golang.org/grpc"
)

var peerlist *models.PeerList
var clients []models.P2PClient

func main(){

    data, err := os.ReadFile("./peerlist.json")
    fmt.Println("读取peerlist")
    if err != nil {
        return
    }
    fmt.Println("解析peerlist")
    //解析peerlist
    peerlist = new(models.PeerList)
    err = json.Unmarshal(data, peerlist)
    if err != nil {
        return
    }
    fmt.Println("连接peer")
    //连接peerlist中的peer
    for _,peer:=range peerlist.Peers{
        conn,err:=grpc.Dial(fmt.Sprintf("%s:%d",peer.IP,peer.Port),grpc.WithInsecure())
        if err!=nil{
            continue
        }
        clients=append(clients,models.NewP2PClient(conn))
    }
    //发送交易
    //从控制台读取交易
   //错误的block
   block:=&models.Block{
    PVHash: "000000000000000000",
    Timestamp: time.Now().UnixNano(),// 时间戳
    Data: "You are fool",// 数据
    Index: 100,// 区块编号
    Nonce: 0,// 随机数
    Hash: "000000000000000000",// 当前区块哈希
   }

//广播
for _,client:=range clients{
    client.NewBlock(context.Background(),block)
}

}










