package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	// Flag to read the bitcoin.conf file. This is used to create the
	// RPC client with the available `user` and `password`.
	conf := flag.String("conf", "", "Path to bitcoin.conf file")
	if *conf == "" {
		panic("No bitcoin.conf set")
	}
	// Server route configuration
	router := gin.Default()
	api := router.Group("/api/v1")
	api.GET("info", nil)

	if err := router.Run(); err != nil {
		panic(err)
	}

}
