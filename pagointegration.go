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
	//fmt.Println("token Details :", string(tokentokeninterface))
	/*
		fmt.Println("access_token :", tokeninterface.access_token)
		fmt.Println("access_token :", tokeninterface.expires_in)
		fmt.Println("access_token :", tokeninterface.token_type)
	*/
	return tokeninterface, nil
	//return strBody, nil
}
func GetListFromPagoEntities(c *gin.Context) {
	//pago_trans_status_url := "https://ipos-gateway.test.pago.dev/payment-proxy/transaction-status/6dc85cc1-a9f9-4caa-bd1f-a9367d1ed813"

	//return "NO Implementation"
	requestURL := "https://ipos-gateway.test.pago.dev/payment-proxy/pos-entities"
	//requestURL := "http://api.open-notify.org/astros.json"
	//jsonBody := []byte(`{"client_message": "hello, server!"}`)
	//jsonBody := []byte(``)
	//bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		panic(err)
	}
	//req.Header.Set("Content-Type", "application/json")

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