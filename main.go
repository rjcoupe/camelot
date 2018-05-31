package main

import (
	"github.com/rjcoupe/camelot/client"
)

func main() {
	btcd := client.BitcoindClient{RPCServerHost: "localhost", RPCServerPort: 8333}
	btcd.Connect()
	btcd.GetDifficulty()
}
