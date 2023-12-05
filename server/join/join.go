package join

import (
	handler "github.com/epsliondylan/blockchain/handlers/join"
	idl "github.com/epsliondylan/blockchain/idls/join"
)

// JController ...
type JController struct {
}

// GenIdl ...
func (c *JController) GenIdl() interface{} {
	return idl.NewJRequest()
}

// Do ...
func (c *JController) Do(req interface{}) interface{} {
	r := req.(*idl.JRequest)
	return handler.AddPeer(r)
}
