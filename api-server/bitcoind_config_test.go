package main

import (
	"bytes"
	"testing"
)

func TestBitcoinConfigNotExists(t *testing.T) {
	_, err := readBitcoinConfig("")
	if err == nil {
		t.Fatal()
	}
}

func TestValidBitcoinConfigFile(t *testing.T) {
	var content bytes.Buffer
	// This is our valid bitcoind.conf file content
	content.WriteString(`
rpcuser='yourbtc'
rpcpassword='yourbtc'
`)
	config, err := readContent(&content)
	if err != nil {
		t.Fatal(err)
	}
	if config.RpcUser != "yourbtc" && config.RpcPassword != "yourbtc" {
		t.Fatalf(
			"Expected user and password `yourbtc` but was %s/%s",
			config.RpcUser, config.RpcPassword,
		)
	}
}

func TestValidBitcoinConfigFileWithMissingField(t *testing.T) {
	var content bytes.Buffer
	content.WriteString(`rpcpassword='yourbtc'`)
	config, err := readContent(&content)
	if err != nil {
		t.Fatal(err)
	}
	if config.RpcPassword != "yourbtc" {
		t.Fatalf("Expected password `yourbtc` but was %s", config.RpcPassword)
	}
	if config.RpcUser != "" {
		t.Fatalf("Expected empty user but was %s", config.RpcUser)
	}
}
