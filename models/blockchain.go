
package models

import (
	"encoding/json"
	"sync"

	pto "github.com/epsilondylan/blockchain/proto"

	"github.com/epsilondylan/blockchain/common"
)

var lock sync.Mutex

var singleChain *pto.BlockChain

func CreateChain() {
	singleChain = newChain()
	Genesis := &pto.Block{
		PVHash: "0",
		Timestamp: 0,
		Data: "This is Genesis Block",
		Index: 0,
		Nonce: 0,
		Hash: "0"}
	singleChain.Chain = append(singleChain.Chain, Genesis)
}

func newChain() *pto.BlockChain {
	theChain := make([]*pto.Block, 0)
	return &pto.BlockChain{theChain}
}

// FormatChain format received []byte to a blockchain object.
func FormatChain(b []byte) (*pto.BlockChain, error) {
	c := &pto.BlockChain{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, err
}

// AppendChain append a valid block to the chain's tail.
func AppendChain(b *pto.Block) error {
	lock.Lock()
	defer lock.Unlock()
	if !b.IsValid(GetChainTail()) {
		return common.Error(common.ErrInvalidBlock)
	}
	singleChain.Chain = append(singleChain.Chain, b)
	return nil
}

// FetchChain fetch the whole chain.
func FetchChain() *pto.BlockChain {
	return singleChain
}

// GetChainTail get the tail block of the chain.
func GetChainTail() *pto.Block {
	return singleChain.Chain[GetChainLen()-1]
}

// GetChainLen get the chain's length.
func GetChainLen() int64 {
	return int64(len(singleChain.Chain))
}

// ReplaceChain replace the chain by a longer valid chain.
func ReplaceChain(c2 *pto.BlockChain) error {
	lock.Lock()
	defer lock.Unlock()
	if int64(len(c2.Chain)) <= GetChainLen() {
		return common.Error(common.ErrInvalidBlock)
	}
	// TODO use a faster algorithm to check the whole chain.
	for i, b := range c2.Chain {
		if i == 0 {
			if *c2.Chain[i] != *singleChain.Chain[i] {
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

