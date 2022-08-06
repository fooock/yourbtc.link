package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Flag to read the bitcoin.conf file. This is used to create the
	// RPC client with the available `user` and `password`.
	conf := flag.String("conf", "", "Path to bitcoin.conf file")
	flag.Parse()

	// If the configuration is not set then we break the program
	// execution since we don't have data to connect to
	if *conf == "" {
		panic("No bitcoin.conf set")
	}

	// Here we read the bitcoin.conf file and try to parse the
	// required fields in order to create a valid client. This
	// client will be used later to read data from the bitcoin node,
	// so if we have any problem, we panic since we can't do anything.
	bitcoinConfig, err := readBitcoinConfig(*conf)
	if err != nil {
		panic(err)
	}
	btcClient, err := newBitcoinClient(bitcoinConfig)
	if err != nil {
		panic(err)
	}
	defer btcClient.close()

	router := gin.Default()
	api := router.Group("/api/v1")

	// Request to retrieve the blockchain info. This method will return an
	// internal server error (500) if an unexpected error happens, like the
	// bitcoin node is down. In any other case, this will return the raw
	// response from the node in JSON
	api.GET("info", func(context *gin.Context) {
		result, err := btcClient.client.GetBlockChainInfo()
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, result)
	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
