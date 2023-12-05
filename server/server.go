package server

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/epsilondylan/blockchain/server/create"
	"github.com/epsilondylan/blockchain/server/join"
	"github.com/epsilondylan/blockchain/server/show"
	"github.com/epsilondylan/httpsvr"
)

// Serve ...
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
