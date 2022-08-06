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
	RpcUser     string `toml:"rpcuser"`
	RpcPassword string `toml:"rpcpassword"`
}

// readBitcoinConfig reads the content of the `bitcoin.conf` file
// and returns its content or an error
func readBitcoinConfig(file string) (*bitcoinConfig, error) {
	content, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer content.Close()
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
	return &config, nil
}
