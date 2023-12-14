package create

import (
	"github.com/epsilondylan/blockchain/common"
	pto "github.com/epsilondylan/blockchain/protocal"
)



// NewCRequestIDL ...
func NewCRequestIDL() *pto.CRequest {
	return &pto.CRequest{}
}

// CResponse response struct
type CResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

// NewCResponseIDL ...
func NewCResponseIDL() *pto.CResponse {
	return &pto.CResponse{}
}

func GenerateBlock(req *pto.CRequest) *pto.CResponse {
	resp := NewCResponseIDL()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	pto.DataQueueAppend(req)
	return resp
}

// CController ...
type CController struct {
}

// GenIdl ...
func (c *CController) GenIdl() interface{} {
	return NewCRequestIDL()
}

// Do ...
func (c *CController) Do(req interface{}) interface{} {
	r := req.(*pto.CRequest)
	return GenerateBlock(r)
}
