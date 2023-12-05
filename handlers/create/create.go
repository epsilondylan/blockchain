package create

import (
	"github.com/epsilondylan/blockchain/common"
	idl "github.com/epsilondylan/blockchain/idls/create"
	pto "github.com/epsilondylan/blockchain/protocal"
)

// GenerateBlock create a new block and spread it.
func GenerateBlock(req *idl.CRequest) *idl.CResponse {
	resp := idl.NewCResponseIDL()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	pto.DataQueue <- req
	return resp
}
