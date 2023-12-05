package show

import (
	handler "github.com/epsilondylan/blockchain/handlers/show"
	idl "github.com/epsilondylan/blockchain/idls/show"
)

// JController ...
type SController struct {
}

// GenIdl ...
func (c *SController) GenIdl() interface{} {
	return idl.NewJRequest()
}

// Do ...
func (c *SController) Do(req interface{}) interface{} {
	r := req.(*idl.SRequest)
	return handler.Show(r)
}
