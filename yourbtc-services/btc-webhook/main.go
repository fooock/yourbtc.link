package main

import (
	"flag"
	"log"
)

func main() {
	// Flag to connect to the actual ZMQ socket to receive
	// raw block updates.
	rawBlockFlag := flag.String(
		"rawBlock",
		"tcp://127.0.0.1:28332",
		"ZMQ TCP endpoint to read raw blocks",
	)
	// Creates the socket configuration to subscribe to BTC raw block updates.
	// You need to be sure to have the socket created by your Bitcoin node.
	rawBlockConfig := zmqConfiguration{
		endpoint:     *rawBlockFlag,
		subscription: "rawblock",
	}
	// Connect to the socket to be able to start receiving block updates
	blockConn, err := newZmqConnection(rawBlockConfig)
	if err != nil {
		panic(err)
	}
	defer blockConn.close()

	log.Printf("Connected to: socket=%s subscription=%s",
		rawBlockConfig.endpoint, rawBlockConfig.subscription)
}
