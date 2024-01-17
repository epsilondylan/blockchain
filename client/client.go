package main
 
import (
    "context"
    "encoding/json"
    "fmt"
    "os"

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
    for{
        var Account string
        var Cipher string
        var trans string
        fmt.Println("请输入账户名：")
        fmt.Scanln(&Account)
        fmt.Println("请输入密码：")
        fmt.Scanln(&Cipher)
        fmt.Println("请输入交易内容：")
        fmt.Scanln(&trans)
        //生成交易
        t:=&models.Trans{
            Account:Account,
            Cipher:Cipher,
            Transaction:trans,
        }
        //广播交易
        for _,client:=range clients{
            client.NewTransaction(context.Background(),t)
        }
    }
}










