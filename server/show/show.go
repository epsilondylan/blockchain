package show

import (
	models "github.com/epsilondylan/blockchain/models"
	pto "github.com/epsilondylan/blockchain/protocal"
)

// JController ...
type SController struct {
}

// GenIdl ...
func (c *SController) GenIdl() interface{} {
	return NewJRequest()
}

// Do ...
func (c *SController) Do(req interface{}) interface{} {
	r := req.(*SRequest)
	return Show(r)
}

type result struct {
}

// idls
// JRequest request struct
type SRequest struct {
	Chain bool `json:"chain"`
	Peer  bool `json:"peer"`
}

// NewJRequest ...
func NewJRequest() *SRequest {
	return &SRequest{}
}

// JResponse response struct
type SResponse struct {
	Chain interface{} `json:"chain"`
	Peer  interface{} `json:"peer"`
}

// NewJResponse ...
func NewJResponse() *SResponse {
	return &SResponse{}
}

// Show join to the blockchain system by connect to a peer
func Show(req *SRequest) *SResponse {
	resp := NewJResponse()
	single := pto.GetProtocal()
	if req.Chain {
		resp.Chain = models.FetchChain()
	}
	if req.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp
}
