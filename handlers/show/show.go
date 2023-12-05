package show

import (
	"github.com/epsilondylan/blockchain/common"
	idl "github.com/epsilondylan/blockchain/idls/create"
	pto "github.com/epsilondylan/blockchain/protocal
)

type result struct {
}

// Show join to the blockchain system by connect to a peer
func Show(req *idl.SRequest) *idl.SResponse {
	resp := idl.NewJResponse()
	single := pto.GetProtocal()
	if req.Chain {
		resp.Chain = models.FetchChain()
	}
	if req.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp
}
