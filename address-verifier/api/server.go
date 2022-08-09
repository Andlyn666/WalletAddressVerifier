package main

import (
	"address-verifier/message"
	"address-verifier/verifier"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Message string `json:"message"`
}

type VerifyResult struct {
	Verified bool `json:"verified"`
}

func main() {
	// Define the map to store the random message of each address
	var adressMsgMap map[string][]byte
	adressMsgMap = make(map[string][]byte)

	e := echo.New()
	// /get_message API
	e.POST("/get_message", func(c echo.Context) error {
		address := c.FormValue("address")
		var msg = message.GetRandomMessage()
		if msg == nil {
			return c.JSON(http.StatusOK, "null")
		}
		res := &Message{
			Message: hex.EncodeToString(msg),
		}
		adressMsgMap[address] = msg

		return c.JSON(http.StatusOK, res)
	})

	// /verify API
	e.POST("/verify", func(c echo.Context) error {
		address := c.FormValue("address")
		signedMessage := c.FormValue("signedMessage")
		addressBytes, _ := hex.DecodeString(address)
		signedMessageBytes, _ := hex.DecodeString(signedMessage)
		message := adressMsgMap[address]
		var result bool
		if message == nil {
			result = false
		} else {
			result = verifier.VerifyAddressPrivateKey(addressBytes, signedMessageBytes, message)
		}
		res := &VerifyResult{
			Verified: result,
		}
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
