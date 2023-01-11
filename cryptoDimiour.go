package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Get Dimiour accounts.
// @Description display dimiour accounts.
// @Tags dimiour
// @Accept json
// @Produce json
// @Success 200
// @Router /currencies [get]
func DimiourgetBalances(c *gin.Context) {

	var returnData interface{}

	url := "https://api.hitbtc.com/api/3/public/currency"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var data map[string]interface{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Printf("Error %s", err)
		return
	}
	returnData = data
	c.JSON(http.StatusOK, returnData)
}

// HealthCheck godoc
// @Summary coin value
// @Description Get coin
// @Tags dimiour
// @Accept json
// @Produce json
// @Param coin query string true "Enter coin type:"
// @Success 200
// @Router /currency/{coin} [get]
func DimiourgetCoinBlance(c *gin.Context) {

	var returnData interface{}

	coin := c.Param("coin")

	url := "https://api.hitbtc.com/api/3/public/currency?currencies=" + coin

	fmt.Printf("url is => %s", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var data map[string]interface{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Printf("Error %s", err)
		return
	}
	returnData = data
	c.JSON(http.StatusOK, returnData)
}
