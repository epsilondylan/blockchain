package models

import (
	"fmt"
	"time"
	dhash "blockchain/hash"
	"blockchain/common"
	"crypto/sha256"
	"bytes"
)


type utxo struct {
	pubkey string
	hash   string
}

type utxoSet struct {
	utxos []utxo
}

// GenerateBlock Generate a new block, it takes sometime and can be stopped by using the following function.
// hash = Index+PVHash+Timestamp+MerkleRoot
//type Block struct {
//	Index      uint32 
//	PVHash     [32]byte
//	Timestamp  int32 
//  MerkleRoot []byte 
//	Nonce      int32
//	TxHashes   [][]byte
//}
/*
func GenerateBlock(index int32, pvhash []byte, hashes [][]byte) (*Block){
	var metaData []byte// 元数据
	indexBytes := make([]byte, 4)
	timestampBytes := make([]byte, 4)
	tree := NewMerkleTree(len(hashes))
	tree.PopulateTree(nil, hashes)
	MerkleRoot := tree.Root()
	binary.BigEndian.PutUint32(indexBytes, uint32(index))
	time := int32(time.Now().UnixNano())
	binary.BigEndian.PutUint32(timestampBytes, uint32(time))
	// Append the byte slices to metaData
	metaData = append(metaData, indexBytes...)
	metaData = append(metaData, pvhash...)
	metaData = append(metaData, timestampBytes...)
	metaData = append(metaData, MerkleRoot...)
	hash, nonce := dhash.HashwithDifficulty(metaData, common.HashDifficulty)// 计算hash
	if (len(hash) == 0) {
		return nil
	}

	return &Block{
		Index:     uint32(index),
		PVHash:    pvhash,
		Timestamp: time,
		MerkleRoot: MerkleRoot,
		Nonce:     int32(nonce),
		
	}
}
*/
func GenerateBlock(index int64, pvhash string, jsonTxs []string) (*Block){
	var metaData string// 元数据
	// Populate the tree with the hashes
	rootHashes := make([][]byte, len(jsonTxs))
	for i, jTx := range jsonTxs {
		hashtemp := sha256.Sum256([]byte(jTx))
		rootHashes[i] =  hashtemp[:]
	}
	MerkleRoot := []byte("")
	if(len(rootHashes) > 0) {
	MerkleRoot = NewMerkleTree(rootHashes).hash
	}

	time := time.Now().UnixNano()
	// Append the byte slices to metaData
	metaData = metaData + fmt.Sprintf("%d", index)
	metaData = metaData + pvhash
	metaData = metaData + fmt.Sprintf("%d", time)
	metaData = metaData + string(MerkleRoot)
	hash, nonce := dhash.HashwithDifficulty([]byte(metaData), common.HashDifficulty)// 计算hash
	return &Block{
		Index:     int64(index),
		PVHash:    pvhash,
		Timestamp: time,
		MerkleRoot: string(MerkleRoot),
		Nonce:      int64(nonce),
		Hash:  jsonTxs,
		Selfhash: fmt.Sprintf("%x", hash),
	}

}


// Interupt stop calculating hash for the block.
func (b *Block) Interupt() bool {// 中断
	return dhash.StopHash()// 停止计算
}

// IsValid return if the block is legal.
func (b *Block) IsValid(pvb *Block) bool {// 是否合法

	//check the validity of the trans data
	bytehashes := make([][]byte, len(b.Hash))
	for i, tx := range b.Hash {
		temphash := sha256.Sum256([]byte(tx))
		bytehashes[i] = temphash[:]
	}
	if(len(bytehashes) > 0) {
	root := NewMerkleTree(bytehashes)
	MRTree := bytes.Equal(root.hash, []byte(b.MerkleRoot))
	if MRTree == false {
		return false// 不合法
	}
	}
	return true// 合法
}


// MerkleNode represents a node in the Merkle tree
type MerkleNode struct {
	left  *MerkleNode
	right *MerkleNode
	hash  []byte
}

// NewMerkleNode creates a new MerkleNode with the given hash
func NewMerkleNode(hash []byte) *MerkleNode {
	return &MerkleNode{nil, nil, hash}
}

// NewMerkleTree creates a Merkle tree from a list of data blocks
func NewMerkleTree(data [][]byte) *MerkleNode {
	var nodes []MerkleNode

	// Create leaf nodes
	for _, datum := range data {
		temp := sha256.Sum256(datum)
		node := NewMerkleNode(temp[:])
		nodes = append(nodes, *node)
	}

	// Build the tree
	for len(nodes) > 1 {
		var newLevel []MerkleNode

		// Combine pairs of nodes to create parent nodes
		for i := 0; i < len(nodes); i += 2 {
			left := &nodes[i]
			var right *MerkleNode

			if i+1 < len(nodes) {
				right = &nodes[i+1]
			}

			parentHash := sha256.Sum256(append(left.hash, right.hash...))
			parent := NewMerkleNode(parentHash[:])
			parent.left = left
			parent.right = right

			newLevel = append(newLevel, *parent)
		}

		nodes = newLevel
	}

	return &nodes[0]
}

func CompareBlock(b1 *Block, b2 *Block) bool {
	if b1.Index != b2.Index {
		return false
	}
	if b1.Timestamp != b2.Timestamp {
		return false
	}
	if b1.Nonce != b2.Nonce {
		return false
	}
	if b1.PVHash != b2.PVHash {
		return false
	}
	if b1.MerkleRoot != b2.MerkleRoot {
		return false
	}
	return true
}