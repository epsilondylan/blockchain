package models

import (
	"errors"
	"crypto/rsa"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"crypto/sha256"
	"encoding/base64"
)


// TxOut represents a transaction output in Go
type TxOut struct {
	Amount       int
	pubkey       string
}

// NewTxOut creates a new TxOut instance in Go
func NewTxOut(amount int,pubkey string) *TxOut {
	return &TxOut{
		Amount:       amount,
		pubkey:       pubkey,
	}
}

// FormatTrans format []byte to a trans object.
func FormatTrans(s []string) (*TransPool,error) {
	transpool := &TransPool{}
	txgroup := []*Trans{}
	for _,v := range s {
		tx := FromJson([]byte(v))
		if tx == nil {
			return nil, errors.New("invalid transaction")
		}
		txgroup = append(txgroup, tx)
	}
	transpool.TransPool = txgroup
	return transpool, nil
}

//convert TxOut to json
func (txOut *TxOut) ToJson() ([]byte, error) {
	txOutJson, err := json.Marshal(*txOut)
	if err != nil {
		println("error in line 43")
		return nil, err
	}
	println("txOutJson:",string(txOutJson))
	return txOutJson, nil
}

func FromJson(jsonData []byte)(t *Trans) {
	t = &Trans{}
	err := json.Unmarshal(jsonData,t)
	if err != nil {
		return nil
	}
	println("t:",t)
	return t
}

func (tx *Trans)ToJson() ([]byte, error) {
	txJson, err := json.Marshal(*tx)
	if err != nil {
		return nil, err
	}
	return txJson, nil
}

func (tx *Trans)txvalidate(unhashedMessages [][]byte, payers_pubkeys[]string) error {
	// Parse the public key string into an *rsa.PublicKey
	for i, txIn := range tx.Tx_Ins {	
		//get txIn's pubkey...
		pb := payers_pubkeys[i]
		pbKey, err := base64.StdEncoding.DecodeString(pb)
		pubv, err := x509.ParsePKIXPublicKey(pbKey)
		ciphertext, err := base64.StdEncoding.DecodeString(txIn)
		if err != nil {
			return err
		}
		pub := pubv.(*rsa.PublicKey)
		err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, unhashedMessages[i], []byte(ciphertext))
		if err != nil {
			return errors.New("invalid signature")
		}
	}
	return nil
}


func (tx *Trans) Hash() ([]byte) {
	txJson, err := json.Marshal(*tx)
	if err != nil {
		return nil
	}
	hash := sha256.Sum256(txJson)
	return hash[:]
}

func Equal(a *Trans, b *Trans) bool {
	if a == nil || b == nil {
		return false
	}

	// 比较各个字段是否相等
	if a.NumInputs != b.NumInputs || a.NumOutputs != b.NumOutputs ||
		a.Signature != b.Signature || a.Locktime != b.Locktime {
		return false
	}

	// 比较切片字段
	if len(a.Tx_Ins) != len(b.Tx_Ins) || len(a.Tx_Outs) != len(b.Tx_Outs) {
		return false
	}

	for i := range a.Tx_Ins {
		if a.Tx_Ins[i] != b.Tx_Ins[i] {
			return false
		}
	}

	for i := range a.Tx_Outs {
		if a.Tx_Outs[i] != b.Tx_Outs[i] {
			return false
		}
	}

	return true
}
