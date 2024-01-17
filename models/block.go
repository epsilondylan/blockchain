package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	dhash "blockchain/hash"

	"blockchain/common"
)


// FormatBlock Format received []byte to a block object.
func FormatBlock(b []byte) (*Block, error) {
	block := &Block{}// 区块
	err := json.Unmarshal(b, block)// 解析
	if err != nil {// 解析失败
		return nil, err
	}
	return block, nil
}

// GenerateBlock Generate a new block, it takes sometime and can be stopped by using the following function.
// hash = PVHash+Timestamp+Data+n+Nonce.
func GenerateBlock(pvHash, data string, index int64) *Block {
	var metaData string // 元数据（用于处理）
	time := time.Now().UnixNano()// 时间戳
	tStr := strconv.FormatInt(time, 10)// 时间戳字符串
	nStr := strconv.FormatInt(index, 10)// index字符串
	metaData = pvHash + tStr + data + nStr// 元数据的字符串
	hash, nonce := dhash.HashwithDifficulty([]byte(metaData), common.HashDifficulty)// 计算hash
	return &Block{
		PVHash:    pvHash,
		Timestamp: time,
		Data:      data,
		Index:     index,
		Nonce:     nonce,
		Hash:      fmt.Sprintf("%x", hash),
	}
}

// Interupt stop calculating hash for the block.
func (b *Block) Interupt() bool {// 中断
	return dhash.StopHash()// 停止计算
}

// IsValid return if the block is legal.
func (b *Block) IsValid(pvb *Block) bool {// 是否合法
	var metaData string// 元数据
	if b.PVHash != pvb.Hash || (pvb.Index+1) != b.Index {// 判断前后区块是否合法
		return false
	}

	//check the validity of the trans data
	// t, err := FormatTrans([]byte(b.Data))// 解析交易
	// if err != nil {		
	// 	return false// 解析失败
	// }else {
	// 	fmt.Println("2")
	// // for _,transtemp := range t.TransPool{
	// // 	err = transtemp.IsVaild()// 验证交易
	// // 	if err != nil {
	// // 		return false// 验证失败
	// // 	}
	// // }
	// fmt.Println("3")
	// }

	tStr := strconv.FormatInt(b.Timestamp, 10)// 时间戳字符串
	nStr := strconv.FormatInt(b.Index, 10)// index字符串
	noStr := strconv.FormatInt(b.Nonce, 10)
	metaData = b.PVHash + tStr + b.Data + nStr
	return dhash.Verification(append([]byte(metaData), []byte(noStr)...), b.Hash)// 验证hash
}

// IsTempValid return if the block is temporary legal.
func (b *Block) IsTempValid() bool {// 是否临时合法
	var metaData string// 元数据
	//check the validity of the trans data
	// t, err := FormatTrans([]byte(b.Data))// 解析交易
	// if err != nil {
	// 	return false
	// }
	// for _,transtemp := range t.TransPool{
	// 	err = transtemp.IsVaild()// 验证交易
	// 	if err != nil {
	// 		return false// 验证失败
	// 	}
	// }

	tStr := strconv.FormatInt(b.Timestamp, 10)
	nStr := strconv.FormatInt(b.Index, 10)
	noStr := strconv.FormatInt(b.Nonce, 10)
	metaData = b.PVHash + tStr + b.Data + nStr
	return dhash.Verification(append([]byte(metaData), []byte(noStr)...), b.Hash)
}

func CompareBlock(b1, b2 *Block) bool {// 比较区块的全部内容
	if b1.PVHash != b2.PVHash || b1.Timestamp != b2.Timestamp || b1.Data != b2.Data || b1.Index != b2.Index || b1.Nonce != b2.Nonce || b1.Hash != b2.Hash {
		return false
	}
	return true	
}