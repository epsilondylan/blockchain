package main


import(
    "encoding/json"
    "fmt"
    "blockchain/models"
    "net"
    p2p "blockchain/p2p"
    "google.golang.org/grpc"
    "os"

)
/*
type Block struct {
	Index      int64
	PVHash     string
	Timestamp  int64 
    MerkleRoot string
	Nonce      int64
	TxHashes   []string
}
*/

type utxo struct {
    spent map[string]bool
    hash  string
}



func Mine(Server *p2p.P2P_Server){
    for{

        var TxHashesPool []string
        //一次最多十条交易，也可以优化
        for i:=0;i<10;i++ {
            select {
                case trans := <- Server.NewTrans:
                    println("now get new trans")
                    temp,err:=json.Marshal(trans)
                    if err!=nil{
                        println("error is in Mine")
                        return
                    }
                    println("ok to get trans")
                    TxHashesPool=append(TxHashesPool,fmt.Sprintf("%x", temp))
                default:
                    break
            }
        }
        //生成区块
        prevBlock:=models.GetChainTail()
        pvhash := prevBlock.Selfhash
        println("prevBlock is ok")
        block := models.GenerateBlock(models.GetChainTail().Index+1, pvhash , TxHashesPool)
        err:=models.AppendChain(block)
        if err == nil{
            //广播区块
            if (len(block.Hash) > 0){
                //record filepath
                data, err := os.ReadFile("utxoset.json")
                newset := make([]*utxo, 0)
                json.Unmarshal(data, &newset)
                for i := 0; i < len(block.Hash); i++ {
                    falsearray := make(map[string]bool)
                    trans := models.FromJson([]byte(block.Hash[i]))
                    fmt.Println("trans:",trans)
                    for j := 0; j < len(trans.Tx_Outs); j++ {
                        falsearray[trans.Tx_Outs[j]] = false
                    }
                    newutxo := utxo{spent: falsearray, hash: block.Hash[i]}
                    newset = append(newset, &newutxo)
                    data, err := json.Marshal(newset)
                    err = os.WriteFile("./utxoset.json", data, 0644)
                    if err != nil {
                       fmt.Println("failed to write to utxoset.json: %v", err)
                    }
                }
                if err != nil {
                    fmt.Println("failed to write to utxoset.json: %v", err)
                }
                writedata,err := json.Marshal(newset)
                os.WriteFile("./utxoset.json", writedata, 0644)
                fmt.Println("writedata:",writedata)
                fmt.Println("new block added")
            }
            fmt.Println("广播区块")
            Server.Broadcast(block)
            //将区块写入本地
            err1:=Server.SaveChainToLocal()
            if err1!=nil{
                fmt.Println(err1)
            }
        }
    }
}

/*
func GenerateBlock(index int32, pvhash string, hashes []string) (*Block){
	var metaData string// 元数据
	tree := NewMerkleTree(len(hashes))
	// Populate the tree with the hashes
	rootHashes := make([][]byte, len(hashes))
	for i, hash := range hashes {
		rootHashes[i] = []byte(hash)
	}
	tree.PopulateTree(nil, rootHashes)
	MerkleRoot := tree.Root()
	time := time.Now().UnixNano()
	// Append the byte slices to metaData
	metaData = metaData + string(index)
	metaData = metaData + pvhash
	metaData = metaData + string(time)
	metaData = metaData + string(MerkleRoot)
	hash, nonce := dhash.HashwithDifficulty([]byte(metaData), common.HashDifficulty)// 计算hash
	
	if (len(hash) == 0) {
		return nil
	}
	return &Block{
		Index:     int64(index),
		PVHash:    pvhash,
		Timestamp: time,
		MerkleRoot: string(MerkleRoot),
		Nonce:      int64(nonce),
		TxHashes:    hashes,
	}

}
*/
func main(){
    //初始化
    Server:=new(p2p.P2P_Server)
    Server.Init()
    lis, err := net.Listen("tcp", ":1230")
    if err != nil {
        fmt.Println("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    models.RegisterP2PServer(s, Server)
    go Server.SetupConnections()
    go Mine(Server)
    if err := s.Serve(lis); err != nil {
        fmt.Println("failed to serve: %v", err)
    }

}

