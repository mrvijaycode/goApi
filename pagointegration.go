package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept json
// @Produce json
// @Success 200
// @Router /pagotoken [get]
func GetAuthToken(c *gin.Context) {
	fmt.Println(" inside GetAuthToken")
	//pagoAuthToken := "no_token_found"
	pagoAuthUrl := "https://auth-service.dev.pago.dev/token"
	pago_token, err := HttpGetAuthTokenFromPAGO(pagoAuthUrl)
	if err != nil {
		fmt.Println("error while accessing auth token from pago :", err)
		return
	}
	//fmt.Println("token : ",pago_token)
	c.JSON(http.StatusOK, pago_token)
	//return pagoAuthToken
}

func HttpGetAuthTokenFromPAGO(fullUrl string) (interface{}, error) {
	jsonbody := []byte(`{"permissions": "grant_type=client_credentials&scope=pos-gateway/processor"}`)
	reqbody := bytes.NewReader(jsonbody)
	req, err := http.NewRequest(http.MethodPost, fullUrl, reqbody)
	if err != nil {
		fmt.Printf("client: could not make new request: %s\n", err)
		return nil, err
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Basic NDBjZ2Jvb2k3OWt1MDB1NWdmNTU5OHB2ZmM6MWNiN2NiNWNrMjhhM203cGtkdnE4ZGI2YnFzcHNwdGlxNzg2ZDJvZHNoOTk2ZW9wZTZlcw==")
	/* req.Header = map[string][]string{
		"authorization": {"Basic NDBjZ2Jvb2k3OWt1MDB1NWdmNTU5OHB2ZmM6MWNiN2NiNWNrMjhhM203cGtkdnE4ZGI2YnFzcHNwdGlxNzg2ZDJvZHNoOTk2ZW9wZTZlcw=="},
	} */
	res, err1 := client.Do(req)
	if err1 != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		//panic(err)
		return nil, err1
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var tokeninterface interface{}
	fmt.Println("body from auth token")
	//strBody := string(body)
	//fmt.Println(strBody)

	jsonErr := json.Unmarshal(body, &tokeninterface)

	if jsonErr != nil {
		fmt.Printf("Error in JSON unmarshalling %s", err)
		return nil, jsonErr
	}
	return tokeninterface, nil
	//return strBody, nil
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept json
// @Produce json
// @Success 200
// @Router /pagoentities [get]
func GetListFromPagoEntities(c *gin.Context) {
	requestURL := "https://ipos-gateway.test.pago.dev/payment-proxy/pos-entities"
	//requestURL := "https://ipos-gateway.dev.pago.dev/payment-proxy/pos-entities"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
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

	var pago_entities interface{}

	jsonErr := json.Unmarshal(body, &pago_entities)

	if jsonErr != nil {
		fmt.Printf("Error %s", err)
		return
	}
	fmt.Println(pago_entities)
	c.JSON(http.StatusOK, pago_entities)
}

//var data1 interface{}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description Post transaction.
// @Tags root
// @Accept json
// @Produce json
// @Param transaction_details body interface{} true "Add transaction details"
// @Success 200
// @Router /postTransaction [post]
func PostTransaction(c *gin.Context) {
	postURL := "https://ipos-gateway.test.pago.dev/payment-proxy/transaction-id"

	// JSON body
	/*body := []byte(`{
	"posId": "pos_1",
	"merchantTransactionId":"65",
	"amount": 2.1,
	"receipt": "Receipt",
	"callbackUrl":"https://acme.com/transaction-status/123456789"
	}`)
	*/

	transaction_details := c.Request.Body

	jsonData, err := io.ReadAll(transaction_details)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("-- %s", jsonData)

	if err != nil {
		fmt.Println(err)
	}

	r, err := http.NewRequest(http.MethodPost, postURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	resbody, err := io.ReadAll(res.Body)

	var pago_res interface{}

	jsonErr := json.Unmarshal(resbody, &pago_res)
	if jsonErr != nil {
		fmt.Printf("Error %s", err)
	}

	c.JSON(http.StatusOK, pago_res)
}
