package models

import (
	"encoding/json"
	"github.com/epsilondylan/blockchain/keygen"
	pto "github.com/epsilondylan/blockchain/proto"
)



// IsVaild return if a trans is legal.
func (t *pto.Trans) IsVaild() error {
	return keygen.Verify(t.Account, t.Cipher, []byte(t.Transaction))
}

// FormatTrans format []byte to a trans object.
func FormatTrans(b []byte) (*pto.TransPool,err) {
	transpool := &pto.TransPool{}
	err := json.Unmarshal(b, transpool)
	if err != nil {
		return nil, err
	}
	return transpool, nil
}

// GenerateTransWithID generate a trans using user's ID.
func GenerateTransWithID(id, data string) (*pto.Trans, error) {
	a, c, err := keygen.Signature(id, []byte(data))
	if err != nil {
		return nil, err
	}
	return &pto.Trans{
		Account:     a,
		Cipher:      c,
		Transaction: data}, nil
}

// GenerateTransWithKey generate a trans using the key.
func GenerateTransWithKey(pb, pv, data string) (*pto.Trans, error) {
	c, err := keygen.Signature2(pv, []byte(data))
	if err != nil {
		return nil, err
	}
	return &pto.Trans{
		Account:     pb,
		Cipher:      c,
		Transaction: data}, nil
}
