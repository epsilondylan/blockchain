package main

import (
	"errors"
	"fmt"
	"net"
	"os"

	"flag"

<<<<<<< HEAD
	"github.com/Blockchain-CN/blockchain/common"
	pto "github.com/Blockchain-CN/blockchain/protocal"
	"github.com/Blockchain-CN/blockchain/server"
=======
	"github.com/epsilondylan/blockchain/common"
	pto "github.com/epsilondylan/blockchain/protocal"
>>>>>>> 6184c3173742a446ea79973f779ecb5fb37b4aa6
)

func main() {
	var (
		ServerPort string
		P2PPort    string
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
	if err := server.Serve(ip + ServerPort); err != nil {
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
