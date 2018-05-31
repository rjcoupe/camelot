package client

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

var rpcClient *rpc.Client
var jsonClient *rpc.Client

// BitcoindClient represents a client interface to a bitcoind node
type BitcoindClient struct {
	RPCServerHost string
	RPCServerPort int
	jsonClient    *rpc.Client
}

// Connect initiates a connection to the bitcoind node
func (client *BitcoindClient) Connect() {
	var err error
	address := client.RPCServerHost + ":" + strconv.Itoa(client.RPCServerPort)
	rpcClient, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	client.jsonClient = jsonrpc.NewClient(rpcClient)
}

// Send is a utility function to send a method and, optionally, arguments to the bitcoind node's JSON-RPC interface
func (client BitcoindClient) Send(method string, args interface{}) string {
	var reply string
	err := client.jsonClient.Call(method, args, &reply)
	spew.Dump(reply)
	if err != nil {
		log.Fatal("JSON call error:", err)
	}
	return reply
}

// GetDifficulty reports the current difficulty level
func (client BitcoindClient) GetDifficulty() int64 {
	response := client.Send("getdifficulty", nil)
	spew.Dump(response)
	return 1
}
