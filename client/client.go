package main
 
import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "blockchain/models"
    "google.golang.org/grpc"
    keygen "blockchain/keygen"
    "time"
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
        var numInputs int
        var numOutputs int
        var txIns []string
        var pubkeys []string
        var payouts []int32
        var myPubkey string
        fmt.Println("请输入交易的输入数量")
        fmt.Scanln(&numInputs)
        fmt.Println("请输入交易的输出数量")
        fmt.Scanln(&numOutputs)
        for i:=0;i<numInputs;i++{
            var txIn string
            fmt.Println("请输入交易的第",i+1,"个输入")
            fmt.Scanln(&txIn)
            txIns=append(txIns,txIn)
        }
        for i:=0;i<numOutputs;i++{
            var amount int
            var account int
            var pubkey string
            fmt.Println("请输入交易的第",i+1,"个输出的金额")
            fmt.Scanln(&amount)
            fmt.Println("请输入交易的第",i+1,"个输出的账户")
            fmt.Scanln(&account)
            pubkey, cs, err := keygen.Signature(fmt.Sprintf("alice%d", account), []byte(""))
            if err != nil {
                fmt.Println("签名失败")
                fmt.Println(err)
                return
            }
            if (len(cs) < 0) {
                fmt.Println("签名失败")
                fmt.Println(err)
                return
            }
            pubkeys=append(pubkeys,pubkey)
            payouts = append(payouts,int32(amount))
        }
        fmt.Println("请输入本人账户")
        account := 0
        fmt.Scanln(&account)
        myPubkey, cs, err := keygen.Signature(fmt.Sprintf("alice%d", account), []byte(""))
        if err != nil {
            fmt.Println("签名失败")
            return
        }
        if (len(cs) < 0) {
            fmt.Println("签名失败")
            return
        }
        time := time.Now().Unix()
        t:=&models.Trans{NumInputs:int32(numInputs),NumOutputs:int32(numOutputs),Tx_Ins:txIns,Tx_Outs:pubkeys,PayOut:payouts,Locktime:time,Pubkey:myPubkey,Signature:""}
        //广播交易
        for _,client:=range clients{
            client.NewTransaction(context.Background(),t)
        }
    }
}




/*
syntax = "proto3";
package p2p;
option go_package = "/models";

message Block{
    int64 Index = 1;
    string PVHash = 2;
    int64 Timestamp = 3;
    string MerkleRoot = 4;
    int64 Nonce = 5;
    repeated string Hash = 6;
}

    
message Trans{
    int32 numInputs = 1; 
    int32 numOutputs = 2;
    repeated string Tx_Ins = 3;
    repeated string Tx_Outs = 4;
    int64 Locktime = 5;
    repeated string Unhashed_Txs = 6;
    repeated string pubkeys = 7;
}

message BlockChain{
    repeated Block Chain = 1;
}

message TailBlockRequest{}

message UpdateBlockChainResponse{}

message NewTransactionResponse{}

message NewBlockResponse{
    bool ChainNeedUpdate = 1;
}

message Peer{
    string IP = 1;
    int32 Port = 2;
}

message PeerList{
    repeated Peer Peers = 1;
}

message TransPool{
    repeated Trans TransPool = 1;
}


service P2P{
    rpc RequestTailBlock(TailBlockRequest) returns(Block){}
    rpc UpdateBlockChain(BlockChain) returns(UpdateBlockChainResponse){}
    rpc NewBlock(Block) returns (NewBlockResponse){}
    rpc NewTransaction(Trans) returns(NewTransactionResponse){}
    rpc NewPeer(Peer) returns(PeerList){}
}
*/





