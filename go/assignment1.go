package main

import (
	//component
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//api
	"github.com/gin-gonic/gin"

	//horizon api
	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)
type coinmarketcap struct {
	ID              string
	Name            string
	Symbol          string
	PercentChange1h float64 `json:"percent_change_1h,string"`
}

func getDataAndCompareCoin(c *gin.Context) { //assignemnt 1
	ticker_symbol_1 := c.Param("ticker_symbol_1")
	ticker_symbol_2 := c.Param("ticker_symbol_2")
	url := "https://api.coinmarketcap.com/v1/ticker/"
	coinData := make([]coinmarketcap, 0)
	err := getCoinData(url, &coinData)
	if err != nil {
		log.Printf("%v", err)
	} else {
		//define variables
		check_ticker_symbol_1 := false
		check_ticker_symbol_2 := false
		var name_ticker_symbol_1 string
		var name_ticker_symbol_2 string
		var percentChange_symbol_1 float64
		var percentChange_symbol_2 float64

		for i := 0; i < len(coinData); i++ {
			symbol := coinData[i].Symbol
			if ticker_symbol_1 == symbol {
				name_ticker_symbol_1 = coinData[i].Name
				percentChange_symbol_1 = coinData[i].PercentChange1h
				check_ticker_symbol_1 = true
			}
			if ticker_symbol_2 == symbol {
				name_ticker_symbol_2 = coinData[i].Name
				percentChange_symbol_2 = coinData[i].PercentChange1h
				check_ticker_symbol_2 = true
			}
		}
		nameCoinMoreChange := "Nothing"
		checkEqual := false
		if check_ticker_symbol_1 == true && check_ticker_symbol_2 == true {
			for i := 0; i < 2; i++ {
				if percentChange_symbol_1 > percentChange_symbol_2 {
					nameCoinMoreChange = name_ticker_symbol_1
				}
				if percentChange_symbol_2 > percentChange_symbol_1 {
					nameCoinMoreChange = name_ticker_symbol_2
				}
				if percentChange_symbol_1 == percentChange_symbol_2 {
					nameCoinMoreChange = "Equal"
					checkEqual = true
				}
			}
			if checkEqual == false {
				c.JSON(200, gin.H{
					"name of coin": nameCoinMoreChange,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"error": "Nothing Data",
			})
		}
	}
}

func transferStellar(c *gin.Context) { //assignment 2
	source_account := c.Param("source_account")
	destination_account := c.Param("destination_account")
	amount := c.Param("amount")
	if _, err := horizon.DefaultTestNetClient.LoadAccount(destination_account); err != nil {
        panic(err)
    }

	//baseFee := uint64(100) //change to XLM (Lumen)
	tx, err := b.Transaction( //create a transaction
		b.SourceAccount{AddressOrSeed: source_account},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		//b.BaseFee{baseFee},
		b.Payment( // create a payment
			b.Destination{AddressOrSeed: destination_account},
			b.NativeAmount{Amount: amount},
		),
	)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Transaction failed, may occured from Source_Account (need to be Private Key) or Destination_Accounts (need to be Public Key) or Amount is not enough",
		})
		panic(err)
	}

	//transaction signer with XDR
	txe, err := tx.Sign(source_account)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Cannot sign, maybe you are using Public Key",
		})
		panic(err)
	}

	//getting Transaction signature 
	txeB64, err := txe.Base64()
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Cannot changes to Base64",
		})
		panic(err)
	}
	fmt.Printf("tx base64: %s", txeB64)

	//submit transaction
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Transaction failed",
		})
		panic(err)
	}
	c.JSON(200, gin.H{
		"Congratulations!": "Transaction success",
		"Hash": resp.Hash,
		"Ledger": resp.Ledger,
		"Result": resp.Result,
		"Env": resp.Env,
	})
	fmt.Println("transaction posted in ledger:", resp.Ledger)

}

func bonus(c *gin.Context) { // bonus
	source_account := c.Param("source_account")
	destination_account := c.Param("destination_account")
	amount := c.Param("amount")

	//check exist account
	if _, err := horizon.DefaultTestNetClient.LoadAccount(destination_account); err != nil {
        panic(err)
	}

	//baseFee := uint64(100) //change to XLM (Lumen)
	tx, err := b.Transaction( //create a transaction
		b.SourceAccount{AddressOrSeed: source_account},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		//b.BaseFee{baseFee},
		b.Payment( // create a payment
			b.Destination{AddressOrSeed: destination_account},
			b.CreditAmount{Code:"RYU",Issuer:source_account,Amount:amount},
		),
	)
	//transaction signer with XDR
	txe, err := tx.Sign(source_account)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Cannot sign, maybe you are using Public Key",
		})
		panic(err)
	}

	//getting Transaction signature 
	txeB64, err := txe.Base64()
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Cannot changes to Base64",
		})
		panic(err)
	}
	fmt.Printf("tx base64: %s", txeB64)

	//submit transaction
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "Transaction failed",
		})
		panic(err)
	}
	c.JSON(200, gin.H{
		"Congratulations!": "Transaction success",
		"Hash": resp.Hash,
		"Ledger": resp.Ledger,
		"Result": resp.Result,
		"Env": resp.Env,
	})
	fmt.Println("transaction posted in ledger:", resp.Ledger)
}

func getCoinData(url string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func splitSlash(data rune) bool {
	return data == '\n' || data == '\\'
}

func main() {
	api := gin.Default()
	api.GET("/api/comparecoin/:ticker_symbol_1/:ticker_symbol_2", getDataAndCompareCoin)
	api.GET("/api/transferstellar/:source_account/:destination_account/:amount", transferStellar)
	api.GET("/api/bonus/:source_account/:destination_account/:amount", bonus)

	api.Run(":5000")
}
