package join

import (
	"github.com/epsilondylan/blockchain/common"
	idl "github.com/epsilondylan/blockchain/idls/create"
	pto "github.com/epsilondylan/blockchain/protocal"
)

// AddPeer join to the blockchain system by connect to a peer
func AddPeer(req *idl.JRequest) *idl.JResponse {
	resp := idl.NewJResponse()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	err := pto.AddPeer(req.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp
}
