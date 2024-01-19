package models

import (
	"encoding/json"
	"sync"
	common "blockchain/common"
	keygen "blockchain/keygen"
	"fmt"
	"os"
)

var lock sync.Mutex

// TheChain BlockChain struct.
type TheChain struct {
	Chain []*Block `json:"chain"`
}

var singleChain *BlockChain

/*
message Block{
    int64 Index = 1;
    string PVHash = 2;
    int64 Timestamp = 3;
    string MerkleRoot = 4;
    int64 Nonce = 5;
    repeated string Hash = 6;
}
*/

func CreateChain() {
	singleChain = newChain()
	fmt.Println("1")
	alice_pubkeys := make([]string, 10)
	outpay := make([]int32, 10)
	for i := 0; i < 10; i++ {
		outpay[i] = 100
	}
	cs := string("")
	err := error(nil)
	for i := 0; i < 10; i++ {
		alice_pubkeys[i], cs, err = keygen.Signature(fmt.Sprintf("alice%d", i), []byte(""));
	}
	if (len(cs) < 0) {
		fmt.Println("failed to generate key pair: %v", err)
		return
	}
	if (err != nil) {
		fmt.Println("failed to generate key pair: %v", err)
		return
	}
	noka_pubkey ,cs , err := keygen.Signature("nokamoto", []byte(""));
	initialTx := &Trans{
		NumInputs:  1,
		NumOutputs: 10,
		Tx_Ins:   []string{},
		Signature: "",
		PayOut: outpay,
		Pubkey: noka_pubkey,
		Tx_Outs:  alice_pubkeys,
		Locktime: 0,
	}
	hashedMessage := initialTx.Hash()
	pubkey, signature, err := keygen.Signature("nokamoto",hashedMessage)
	fmt.Println("5")
	if err != nil {
		fmt.Println("failed to sign transaction: %v", err)
		return
	}
	fmt.Println("6")
	if (len(pubkey) < 0) {
		fmt.Println("failed to sign transaction: %v", err)
		return
	}
	// Sign the transaction with Noka's private key
	initialTx = &Trans{
		NumInputs:  1,
		NumOutputs: 10,
		Tx_Ins:   []string{},
		Signature: signature,
		Pubkey: pubkey,
		PayOut: outpay,
		Tx_Outs:  alice_pubkeys,
		Locktime: 0,
	}
	// Write to utxoset.json
	filepath := "utxoset.json"
	btx,err := initialTx.ToJson()
	fmt.Println("7")
	if err != nil {
		fmt.Println("failed to sign transaction: %v", err)
		return
	}
	stx := fmt.Sprintf("%s", btx)
	fmt.Println(stx)
	err = os.WriteFile(filepath, []byte(stx), 0644)
	if err != nil {
		fmt.Println("failed to write to utxoset.json: %v", err)
		return
	}
	fmt.Println("8")
	Genesis := &Block{
		Index:     0,
		PVHash:    "",
		Timestamp: 0,
		MerkleRoot: "",
		Nonce:     0,
		Hash: []string{stx},
	}
	fmt.Println("9")
	singleChain.Chain = append(singleChain.Chain, Genesis)
	fmt.Println("10")

}

func newChain() *BlockChain {
	theChain := make([]*Block, 0)
	return &BlockChain{Chain: theChain}
}

// FormatChain format received []byte to a blockchain object.
func FormatChain(b []byte) (*BlockChain, error) {
	c := &BlockChain{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, err
}

// AppendChain append a valid block to the chain's tail.
func AppendChain(b *Block) error {
	lock.Lock()
	defer lock.Unlock()
	if !b.IsValid(GetChainTail()) {
		return common.Error(common.ErrInvalidBlock)
	}
	singleChain.Chain = append(singleChain.Chain, b)
	return nil
}

// FetchChain fetch the whole chain.
func FetchChain() *BlockChain {
	return singleChain
}

// GetChainTail get the tail block of the chain.
func GetChainTail() *Block {
	return singleChain.Chain[GetChainLen()-1]
}

// GetChainLen get the chain's length.
func GetChainLen() int64 {
	return int64(len(singleChain.Chain))
}

// ReplaceChain replace the chain by a longer valid chain.
func ReplaceChain(c2 *BlockChain) error {
	lock.Lock()
	defer lock.Unlock()
	if int64(len(c2.Chain)) <= GetChainLen() {
		return common.Error(common.ErrInvalidBlock)
	}
	// TODO use a faster algorithm to check the whole chain.
	for i, b := range c2.Chain {
		if i == 0 {
			if CompareBlock(b, singleChain.Chain[0]) == false {
				return common.Error(common.ErrInvalidGenesisBlock)
			}
			continue
		}
		if !b.IsValid(c2.Chain[i-1]) {
			return common.Error(common.ErrInvalidBlock)
		}
	}
	singleChain.Chain = c2.Chain
	return nil
}

