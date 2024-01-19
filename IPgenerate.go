package main
import(
    "encoding/json"
    "blockchain/models"
    "os"
    "fmt"
)

var peerlist models.PeerList

func main(){
    Peer1:=&models.Peer{
        IP:"10.1.0.91",
        Port:1230,
        }
    Peer2:=&models.Peer{
        IP:"10.1.0.92",
        Port:1230,
        }
    Peer3:=&models.Peer{
        IP:"10.1.0.93",
        Port:1230,
        }
    Peer4:=&models.Peer{
        IP:"10.1.0.94",
        Port:1230,
        }
    Peer5:=&models.Peer{
        IP:"10.1.0.95",
        Port:1230,
        }
        peerlist.Peers=append(peerlist.Peers,Peer1)
        peerlist.Peers=append(peerlist.Peers,Peer2)
        peerlist.Peers=append(peerlist.Peers,Peer3)
        peerlist.Peers=append(peerlist.Peers,Peer4)
        peerlist.Peers=append(peerlist.Peers,Peer5)
    //写在本地
    temp,err:=json.Marshal(peerlist)
    if err!=nil{
        fmt.Println(err)
    }
    f,err:=os.Create("peerlist.json")
    if err!=nil{
        fmt.Println(err)
    }
    defer f.Close()
    f.Write(temp)
    f.Sync()
    
}
