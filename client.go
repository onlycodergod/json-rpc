package main

import (
	"bufio"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

// clientRequest - структура клиента rpc
type clientRequest []struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	TNVED   string `json:"tnved"`
}

type Reply struct {
	Data string
}

// clientReq - функция клиента
func clientReq() {
	client, err := jsonrpc.Dial("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply Reply
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v, Data: %v", reply, reply.Data)
	}
}
