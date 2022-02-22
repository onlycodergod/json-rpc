package rpcexample

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// serverRequest - структура клиента rpc
type serverRequest struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	TNVED   string `json:"tnved"`
}

// serverReq - функция сервера
func serverReq() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	rpc.Register(inbound)
	for {
		conn, err := inbound.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
