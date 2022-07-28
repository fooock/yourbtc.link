package main

import (
	"github.com/pebbe/zmq4"
	"time"
)

// zmqConnection represents a connection to a read only socket where
// we have a subscription based on a known filter.
type zmqConnection struct {
	socket *zmq4.Socket
}

// zmqConfiguration is the configuration needed to be able to establish
// the socket connection and maintain it. It also contains the
// name of the subscription filter.
type zmqConfiguration struct {
	endpoint     string
	subscription string
	timeout      time.Duration
}

// newZmqConnection tries to create the socket connection or return
// an error. If this method returns the connection, then we can assume
// that the connection is configured and ready to start receiving
// messages when needed.
func newZmqConnection(config zmqConfiguration) (*zmqConnection, error) {
	socket, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		return nil, err
	}
	err = socket.SetSubscribe(config.subscription)
	if err != nil {
		return nil, err
	}
	err = socket.SetConnectTimeout(config.timeout)
	if err != nil {
		return nil, err
	}
	err = socket.Connect(config.endpoint)
	if err != nil {
		return nil, err
	}
	return &zmqConnection{socket: socket}, nil
}

// close the socket connection if possible.
func (s *zmqConnection) close() {
	s.socket.Close()
}
