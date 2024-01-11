package models

import (
	"encoding/json"
	"blockchain/keygen"

)



// IsVaild return if a trans is legal.
func (t *Trans) IsVaild() error {
	return keygen.Verify(t.Account, t.Cipher, []byte(t.Transaction))
}

// FormatTrans format []byte to a trans object.
func FormatTrans(b []byte) (*TransPool,error) {
	transpool := &TransPool{}
	err := json.Unmarshal(b, transpool)
	if err != nil {
		return nil, err
	}
	return transpool, nil
}

// GenerateTransWithID generate a trans using user's ID.
func GenerateTransWithID(id, data string) (*Trans, error) {
	a, c, err := keygen.Signature(id, []byte(data))
	if err != nil {
		return nil, err
	}
	return &Trans{
		Account:     a,
		Cipher:      c,
		Transaction: data}, nil
}

// GenerateTransWithKey generate a trans using the key.
func GenerateTransWithKey(pb, pv, data string) (*Trans, error) {
	c, err := keygen.Signature2(pv, []byte(data))
	if err != nil {
		return nil, err
	}
	return &Trans{
		Account:     pb,
		Cipher:      c,
		Transaction: data}, nil
}
