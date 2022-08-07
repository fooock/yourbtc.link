package main

import (
	"github.com/pelletier/go-toml/v2"
	"io"
	"io/ioutil"
	"os"
)

// bitcoinConfig is the content of the bitcoin.conf in a
// structured way
type bitcoinConfig struct {
	RpcBind     string `toml:"rpcbind"`
	RpcPort     int    `toml:"rpcport"`
	RpcUser     string `toml:"rpcuser"`
	RpcPassword string `toml:"rpcpassword"`
	Chain       string
}

func (c *bitcoinConfig) defaultPort() int {
	switch c.Chain {
	case "testnet":
		return 18332
	case "signet":
		return 38332
	case "regtest":
		return 18443
	// In any other case we assume that the connection
	// is done by using the main network
	default:
		return 8332
	}
}

// readBitcoinConfig reads the content of the `bitcoin.conf` file
// and returns its content or an error
func readBitcoinConfig(file string) (*bitcoinConfig, error) {
	content, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(content *os.File) {
		// close and ignore error
		_ = content.Close()
	}(content)

	// actual configuration parsing
	return readContent(content)
}

func readContent(reader io.Reader) (*bitcoinConfig, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var config bitcoinConfig
	err = toml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	// If rpcBind is not set, then by default we use a sane
	// default like localhost
	if config.RpcBind == "" {
		config.RpcBind = "localhost"
	}
	// If chain is not set we use 'main' as default
	if config.Chain == "" {
		config.Chain = "main"
	}
	// If port is not set then we use the default one
	// based on the selected chain
	if config.RpcPort == 0 {
		config.RpcPort = config.defaultPort()
	}
	return &config, nil
}
