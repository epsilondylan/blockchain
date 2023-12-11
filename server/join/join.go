package join

import (
	"github.com/epsilondylan/blockchain/common"
	pto "github.com/epsilondylan/blockchain/protocal"
)

func AddPeer(req *JRequest) *JResponse {
	resp := NewJResponse()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	err := pto.AddPeer(req.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp
}

type JRequest struct {
	PeerAddr string `json:"peer_addr"`
}

// NewJRequest ...
func NewJRequest() *JRequest {
	return &JRequest{}
}

// JResponse response struct
type JResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

// NewJResponse ...
func NewJResponse() *JResponse {
	return &JResponse{}
}

// JController ...
type JController struct {
}

// GenIdl ...
func (c *JController) GenIdl() interface{} {
	return NewJRequest()
}

// Do ...
func (c *JController) Do(req interface{}) interface{} {
	r := req.(*JRequest)
	return AddPeer(r)
}
