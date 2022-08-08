package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// fetchBitcoinInfo is a server handler to return the bitcoin data
// or an error if some unexpected error happens
func fetchBitcoinInfo(btcClient *bitcoinClient) func(context *gin.Context) {
	return func(context *gin.Context) {
		result, err := btcClient.info()
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, result)
	}
}
