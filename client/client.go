package main

import (
	"errors"
	"fmt"
	"os"
	"net"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"

	"github.com/julienschmidt/httprouter"
	logger "github.com/shengkehua/xlog4go"
	"github.com/epsilondylan/blockchain/common"
	
	pto "github.com/epsilondylan/blockchain/protocal"
	"flag"
)

func main() {
	var (
		ServerPort string
		P2PPort	   string
	)

	flag.StringVar(&ServerPort, "server", "", "-server=:10024")
	flag.StringVar(&P2PPort, "p2p", "", "-p2p=:12345")
	flag.Parse()
	if ServerPort == "" || P2PPort == "" {
		useage()
		printAndDie(errors.New("Unable to get a avilable port for p2p node"))

	}

	ip := getIP()
	if ip == "" {
		printAndDie(errors.New("Unable to get a avilable ip"))
	}
	// ip = "127.0.0.1"

	// init protocal
	pto.InitPto(ip+P2PPort, common.P2PTimeOut)

	// call this func will block current goroutine
	if err := server.Serve(ip+ServerPort); err != nil {
		printAndDie(err)
		return
	}
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func printAndDie(err error) {
	fmt.Fprintf(os.Stderr, "init failed, err:%s", err)
	os.Exit(-1)
}

func useage() {
	fmt.Fprintf(os.Stdout, "please run \"%s --help\" and get help info\n", os.Args[0])
	os.Exit(-1)
}

func Serve(addr string) error {
	s := httpsvr.New(addr,
		httpsvr.SetReadTimeout(time.Millisecond*200),
		httpsvr.SetWriteTimeout(time.Millisecond*200),
		httpsvr.SetMaxAccess(100),
	)
	go GracefulExit(s)
	s.AddRoute("POST", "/blockchain/create", &create.CController{})
	s.AddRoute("POST", "/blockchain/join", &join.JController{})
	s.AddRoute("POST", "/blockchain/show", &show.SController{})
	return s.Serve()
}

// GracefulExit 优雅退出
func GracefulExit(svr *httpsvr.Server) {
	sigc := make(chan os.Signal, 0)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	<-sigc
	println("closing agent...")
	svr.GracefulExit()
	println("agent closed.")
	os.Exit(0)
}

var defaultResponse = []byte(`{"errno":0,"errmsg":"ok"}`)

// Server ...
type Server struct {
	addr   string
	router *httprouter.Router
	opt    *option
	oriSvr *http.Server
	// 接入控制：1.为了限制最大并发数(chan)，以及关闭入口；2.为了优雅退出(sync.WaitGroup)
	ac *Access
}

// New ...
func New(addr string, opts ...ServerOption) *Server {
	opt := &option{}
	for _, o := range opts {
		o(opt)
	}
	if addr == "" {
		addr = "127.0.0.1:10024"
	}
	s := &Server{
		addr:   addr,
		router: httprouter.New(),
		opt:    opt,
	}
	if s.opt.maxAccess == 0 {
		s.opt.maxAccess = 1024
	}
	s.ac = NewAccessor(s.opt.maxAccess)
	s.oriSvr = &http.Server{Addr: addr, Handler: s}
	return s
}

// ServeHTTP implement net/http.router
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

// AddRoute ...
func (s *Server) AddRoute(method, path string, ctrl IController) {
	var handle httprouter.Handle = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// http框架虽然会panic-recover但是自己也必须recover，因为接入记录panic后不会正常消去
		defer func() {
			if err := recover(); err != nil {
				w.Write([]byte("Server is busy."))
				stack := make([]byte, 2048)
				stack = stack[:runtime.Stack(stack, false)]
				f := "PANIC: %s\n%s"
				logger.Error(f, err, stack)
			}
		}()
		err := s.ac.InControl()
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(getErrMsg(err))
			return
		}
		defer s.ac.OutControl()
		nt := time.Now()
		// 打印输入请求
		if s.opt.dumpAccess {
			body, _ := ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			logger.Info("request_uri=%s||client_ip=%s||request_body=%s",
				r.URL,
				r.Host,
				string(body))
		}
		// 解析输入参数
		idl := ctrl.GenIdl()
		body, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, idl)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write(getErrMsg(err))
			return
		}

		do := func(r *http.Request, w http.ResponseWriter) {
			var data []byte
			resp := ctrl.Do(idl)
			if resp == nil {
				data = defaultResponse
			}
			data, _ = json.Marshal(resp)
			et := time.Now().Sub(nt)
			logger.Info("request_uri=%s||response=%s||proc_time=%s",
				r.URL, string(data), et.String())
			w.WriteHeader(200)
			w.Write(data)
		}

		do(r, w)
	}
	s.router.Handle(method, path, handle)
}

// Serve ...
func (s *Server) Serve() error {
	fmt.Printf("Serving %s", s.addr)
	return s.oriSvr.ListenAndServe()
}

// GracefulExit ....
func (s *Server) GracefulExit() {
	s.ac.Stop()
}