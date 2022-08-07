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
chain='regtest'
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
	if config.Chain != "regtest" {
		t.Fatalf("Expected chain 'regtest' but found %s", config.Chain)
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
	// default value for rpcBind
	if config.RpcBind != "localhost" {
		t.Fatalf("Expected rpcBind 'localhost' but was: %s", config.RpcBind)
	}
	// if rpcPort is not set then its value will be 0
	if config.RpcPort != 8332 {
		t.Fatalf("rpcPort expected value is 8332 but was %d", config.RpcPort)
	}
	// chain default value is 'main'
	if config.Chain != "main" {
		t.Fatalf("Expected chain 'main' but was: %s", config.Chain)
	}
}
