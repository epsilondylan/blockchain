
package models

import (
	"encoding/json"
	"sync"
	"blockchain/common"
	"fmt"
)

var lock sync.Mutex

var singleChain *BlockChain

func CreateChain() {
	var GenesisTransPool TransPool
	GenesisTrans := &Trans{
		Account: "A",
		Cipher:   "A",
		Transaction: "1000000",
	}
	GenesisTransPool.TransPool = append(GenesisTransPool.TransPool, GenesisTrans)
	tstr, err := json.Marshal(GenesisTransPool)
	if err != nil {
		fmt.Println(err)
	}
	singleChain = newChain()
	Genesis := &Block{
		PVHash: "0",
		Timestamp: 0,
		Data: string(tstr),
		Index: 0,
		Nonce: 0,
		Hash: "0"}
	singleChain.Chain = append(singleChain.Chain, Genesis)
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

