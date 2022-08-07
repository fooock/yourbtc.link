package main

import (
	"fmt"
	"github.com/btcsuite/btcd/rpcclient"
)

// bitcoinClient is the client responsible to talk with our
// Bitcoind RPC daemon
type bitcoinClient struct {
	config *rpcclient.ConnConfig
	client *rpcclient.Client
}

// newBitcoinClient creates a new bitcoind client based in the data
// available in the `bitcoind.conf` file.
func newBitcoinClient(config *bitcoinConfig) (*bitcoinClient, error) {
	// Connection configuration for our new client
	connConfig := &rpcclient.ConnConfig{
		Host:         fmt.Sprintf("%s:%d", config.RpcBind, config.RpcPort),
		Endpoint:     "ws",
		User:         config.RpcUser,
		Pass:         config.RpcPassword,
		Params:       config.Chain,
		DisableTLS:   true,
		HTTPPostMode: true,
	}
	client, err := rpcclient.New(connConfig, nil)
	if err != nil {
		return nil, err
	}
	return &bitcoinClient{connConfig, client}, nil
}

func (c *bitcoinClient) close() {
	c.client.Shutdown()
}
