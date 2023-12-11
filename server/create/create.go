package create

import (
	"github.com/epsilondylan/blockchain/common"
	pto "github.com/epsilondylan/blockchain/protocal"
)

// CRequest request struct
type CRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// NewCRequestIDL ...
func NewCRequestIDL() *CRequest {
	return &CRequest{}
}

// CResponse response struct
type CResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

// NewCResponseIDL ...
func NewCResponseIDL() *CResponse {
	return &CResponse{}
}

func GenerateBlock(req *CRequest) *CResponse {
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
	r := req.(*CRequest)
	return handler.GenerateBlock(r)
}
