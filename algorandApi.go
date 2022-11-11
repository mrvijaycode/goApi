package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Create a new algo account.
// @Description get the status of server.
// @Tags root
// @Accept json
// @Produce json
// @Success 200
// @Router /createalgoaccount [get]
func CreateAlgorandAccount(c *gin.Context) {

	account := crypto.GenerateAccount()
	passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)
	myAddress := account.Address.String()

	var returnAccount interface{}

	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
	} else {

		type Account struct {
			Address    string
			PublicKey  string
			PrivateKey string
			PassPhrase string
		}

		//marshalAccount, _ := json.Marshal(account)
		//fmt.Println("###", string(marshalAccount))

		fmt.Printf("My address: %s\n", myAddress)
		fmt.Printf("My passphrase: %s\n", passphrase)

		marshalPrivateKey, _ := json.Marshal(account.PrivateKey)
		fmt.Printf("Private key: %s\n", string(marshalPrivateKey))

		marshalPubliKey, _ := json.Marshal(account.PublicKey)
		fmt.Printf("Public key: %s\n", string(marshalPubliKey))

		marshal_account, _ := json.Marshal(Account{Address: myAddress, PublicKey: string(marshalPubliKey), PrivateKey: string(marshalPrivateKey), PassPhrase: passphrase})
		//fmt.Println("Account: \n", marshal_account)
		//returnAccount = fmt.Sprintf("Account: %s\n", marshal_account)
		var algoAccount interface{}
		json.Unmarshal(marshal_account, &algoAccount)
		//json.Unmarshal(account, &algoAccount)

		returnAccount = algoAccount

		//fmt.Printf("My marshal: %T\n", marshaled_slice)
		//fmt.Printf("My address: %T\n", account)
		//fmt.Printf("My address: %s\n", myAddress)
		//fmt.Printf("My passphrase: %s\n", passphrase)
		//fmt.Printf("My PrivateKey: %x\n", account.PrivateKey)
		//fmt.Printf("My PublicKey: %x\n", account.PublicKey)
		//fmt.Println("--> Copy down your address and passphrase for future use.")
		//fmt.Println("--> Once secured, press ENTER key to continue...")
		//c.JSON(http.StatusOK, marshal_account)

	}

	c.JSON(http.StatusOK, returnAccount)
}
