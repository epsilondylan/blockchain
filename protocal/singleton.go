package protocal

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/epsilondylan/blockchain/models"
	p2p "github.com/epsilondylan/pheromones"
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

var (
	singleton *Protocal
	// DataQueue data channel
	DataQueue chan *CRequest
	wg        sync.WaitGroup
	hostAddr  string
)

func GetProtocal() *Protocal {
	return singleton
}

func DataQueueAppend(req *CRequest) {
	DataQueue <- req
}

// InitPto init the default protocal object
func InitPto(addr string, to time.Duration) {
	hostAddr = addr
	r1 := p2p.NewSRouter(to)
	p1 := NewProtocal(addr, r1, to)
	s1 := p2p.NewServer(p1, to)
	singleton = p1
	DataQueue = make(chan *CRequest, 100)
	println("P2P Servering on ", addr)
	go BlockPublisher()
	go s1.ListenAndServe(addr)
}

// AddPeer add a peer to the default protocal's router
func AddPeer(addr string) error {
	err := singleton.Add(addr, addr)
	if err != nil {
		return err
	}
	req := &p2p.MsgPto{
		Name:      hostAddr,
		Operation: RequireBlock,
	}
	reqStr, err := json.Marshal(req)
	if err != nil {
		return err
	}
	for reqStr != nil {
		b, err := singleton.Dispatch(addr, reqStr)
		if err != nil {
			println("操作失败", err.Error())
			singleton.Delete(addr)
			return err
		}
		reqStr = nil
		reqStr, err = singleton.Handle(nil, b)
		fmt.Println("处理的返回结果和错误码为：", string(reqStr), err)
	}
	return nil
}

// BlockPublisher loop to publish the block
func BlockPublisher() {
	for {
		select {
		case ud := <-DataQueue:
			// get user object
			user, err := models.Login(ud.Name)
			if err != nil {
				return
			}

			// get trans object
			trans, err := models.GenerateTransWithKey(user.Public, user.Private, ud.Data)
			if err != nil {
				return
			}
			transStr, err := json.Marshal(trans)
			if err != nil {
				return
			}

			// append a block to the chain until succeed
			for {
				// get block object
				block := models.GenerateBlock(models.GetChainTail().Hash, string(transStr), models.GetChainLen())

				// add to blockchain
				err = models.AppendChain(block)
				if err == nil {
					singleton.spreads(block)
					break
				}
			}
		}
	}
}
