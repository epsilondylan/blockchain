package protocal

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	dhash "github.com/epsilondylan/blockchain/hash"
	p2p "github.com/epsilondylan/pheromones"

	"github.com/epsilondylan/blockchain/common"
	"github.com/epsilondylan/blockchain/models"
)

const (
	// RequireBlock 请求最新block
	RequireBlock = "require a block"

	// Publish 发布最新block
	DeliveryBlock = "delivery a block"

	// DeliveryChain 发送整条链
	DeliveryChain = "delivery the block"

	// RequireChain 请求整条链
	RequireChain = "require the block"

	// 未知命令
	UnknownCmd = "unknown cmd"

	defultByte = 10240
)

type Protocal struct {
	HostName string
	Router   p2p.Router // 路由
	to       time.Duration
}

func NewProtocal(name string, r p2p.Router, to time.Duration) *Protocal {
	return &Protocal{name, r, to}
}

func (p *Protocal) GetConnType() p2p.ConnType {
	return p.Router.GetConnType()
}

func (p *Protocal) Handle(c net.Conn, msg []byte) ([]byte, error) { // 处理请求
	if msg == nil {
		return nil, nil
	}
	cType := p.Router.GetConnType()
	req := &p2p.MsgPto{}            // 请求
	resp := &p2p.MsgPto{}           // 响应
	err := json.Unmarshal(msg, req) // 解析请求
	if err != nil {
		return nil, p2p.Error(p2p.ErrMismatchProtocalReq)
	}
	resp.Name = p.HostName
	switch req.Operation { // 根据请求的操作类型进行处理
	case RequireBlock: // 收到请求最新的block
		if cType == p2p.ShortConnection { // 如果是短连接，需要将连接加入路由
			err = p.Router.AddRoute(req.Name, req.Name) // 将请求的name加入路由
			if err != nil {
				fmt.Println(err)
			}
		} else {
			if p.Router.AddRoute(req.Name, c) == nil {
				go p.IOLoop(c)
			}
		}
		c, _ := json.Marshal(models.GetChainTail())
		resp.Operation = DeliveryBlock
		resp.Data = c
	case DeliveryBlock: // 收到最新的block
		dhash.StopHash()        // 停止计算
		defer dhash.StartHash() // defer开启计算
		block, err := models.FormatBlock(req.Data)
		if err != nil {
			return nil, p2p.Error(p2p.ErrMismatchProtocalResp)
		}
		// if the block's index is shorter or invalidate
		tailBlock := models.GetChainTail() // 获取最新的block
		if *block == *tailBlock {
			return nil, nil
		}
		if !block.IsTempValid() || block.Index <= tailBlock.Index { // 如果block不合法或者index小于最新的block
			return nil, common.Error(common.ErrInvalidBlock)
		}
		// if the block can append to the tail
		if block.IsValid(tailBlock) {
			models.AppendChain(block) // 将block添加到链上
			// 并需要向外广播
			go p.spreads(block)
			return nil, nil
		}
		// if the block's index is longer
		resp.Operation = RequireChain
	case RequireChain:
		c, _ := json.Marshal(models.FetchChain())
		resp.Operation = DeliveryChain
		resp.Data = c
	case DeliveryChain:
		dhash.StopHash()
		defer dhash.StartHash()
		chain, err := models.FormatChain(req.Data)
		if err != nil {
			return nil, p2p.Error(p2p.ErrMismatchProtocalResp)
		}
		err = models.ReplaceChain(chain)
		if err != nil {
			return nil, common.Error(common.ErrInvalidChain)
		}
		// 向外广播 models.GetChainTail()
		go p.spreads(models.GetChainTail())
		return nil, nil
	default:
		fmt.Printf("@%s@report: %s operation from @%s@ finished\n", p.HostName, req.Operation, req.Name)
		return nil, nil
	}
	ret, err := json.Marshal(resp)
	fmt.Printf("@%s@report: %s operation from @%s@ succeed\n", p.HostName, req.Operation, req.Name)
	return ret, nil
}

// 长连接的话，需要在加入路由的时刻起携程 循环监控
func (p *Protocal) IOLoop(c net.Conn) {
	fmt.Printf("@%s@report,开启长连接监听: localhost=%s||remotehost=%s\n", p.HostName, c.LocalAddr(), c.RemoteAddr())
	for {
		msg, err := p.read(c)
		if err != nil {
			c.Close()
			return
		}
		fmt.Printf("长连接收到信息, localhost=%s||remotehost=%s||msg=%s\n", c.LocalAddr(), c.RemoteAddr(), string(msg))
		resp, err := p.Handle(c, msg)
		if err != nil || resp == nil {
			fmt.Printf("结束此次会话, localconn=%s||remoteconn=%s||resp=%s||err=%s\n", c.LocalAddr(), c.RemoteAddr(), resp, err)
			continue
		}
		c.SetWriteDeadline(time.Now().Add(p.to))
		_, err = c.Write(resp)
		if err != nil {
			return
		}
		fmt.Printf("长连接发送信息, localconn=%s||remoteconn=%s||msg=%s\n", c.LocalAddr(), c.RemoteAddr(), resp)
	}
}

func (p *Protocal) read(r io.Reader) ([]byte, error) {
	buf := make([]byte, defultByte)
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	// read读出来的是[]byte("abcdefg"+0x00)，带一个结束符，需要去掉
	return buf[:n], nil
}

func (p *Protocal) Add(name string, addr string) error { // 添加路由
	if p.Router.GetConnType() == p2p.ShortConnection {
		return p.Router.AddRoute(name, addr) // 短连接直接加入路由
	}
	c, err := net.Dial("tcp", addr) // 建立长连接
	if err != nil {
		return err
	}
	if p.Router.AddRoute(name, c) == nil {
		go p.IOLoop(c)
	}
	return err
}

func (p *Protocal) GetRouter() p2p.Router {
	return p.Router
}

func (p *Protocal) DispatchAll(msg []byte) map[string][]byte {
	return p.Router.DispatchAll(msg)
}

func (p *Protocal) Dispatch(name string, msg []byte) ([]byte, error) {
	return p.Router.Dispatch(name, msg)
}

func (p *Protocal) Delete(name string) error {
	return p.Router.Delete(name)
}

// spread the latest block to all peers
func (p *Protocal) spreads(block *models.Block) {
	blockStr, err := json.Marshal(block)
	if err != nil {
		return
	}
	req := &p2p.MsgPto{
		Name:      hostAddr,
		Operation: DeliveryBlock,
		Data:      blockStr,
	}
	reqStr, err := json.Marshal(req)
	if err != nil || reqStr == nil {
		return
	}
	peerList := p.GetRouter().FetchPeers()
	if p.GetRouter().GetConnType() == p2p.ShortConnection {
		p.spreadShort(reqStr, peerList)
	}
}

// 同步等待和所有peer通信完毕
func (p *Protocal) spreadShort(reqStr []byte, peerList map[string]interface{}) { // 短连接
	for _, v := range peerList { // 遍历所有peerList
		wg.Add(1)              // 等待组+1
		go func(addr string) { // 开启协程
			for reqStr != nil { // 如果请求不为空
				b, err := p.Dispatch(addr, reqStr) // 向addr发送请求
				if err != nil {
					println("操作失败", err.Error())
					return
				} // 如果没有错误
				reqStr = nil                     // 请求置空
				reqStr, err = p.Handle(nil, b)   // 处理请求
				fmt.Println(string(reqStr), err) //	打印请求和错误
			}
			wg.Done()
		}(v.(p2p.EndPointS).Addr) // 将v转换为EndPointS类型，取出addr
	}
	wg.Wait()
}


