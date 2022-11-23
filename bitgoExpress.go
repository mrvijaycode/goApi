package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Get Bitgo wallet accounts.
// @Description Display the wallet accounts.
// @Tags root
// @Accept json
// @Produce json
// @Success 200
// @Router /bitgoAccounts [get]
func GetBitgoWalletAddress(c *gin.Context) {

	bitgoUrl := "https://app.bitgo-test.com/api/v2/talgo/wallet/6364b2ab8e0fdd0007f719603a696848/addresses"

	req, err := http.NewRequest(http.MethodGet, bitgoUrl, nil)
	req.Header.Set("Authorization", "Bearer v2xb129f1dbe8b1ff2f3c11c967f89012b22e1f6b68a80bdd8638019f9c50e54b0f")

	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		panic(err)
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var bitgoAddresses interface{}

	jsonErr := json.Unmarshal(body, &bitgoAddresses)

	if jsonErr != nil {
		fmt.Printf("Error %s", err)
		return
	}
	fmt.Println(bitgoAddresses)
	c.JSON(http.StatusOK, bitgoAddresses)
}
