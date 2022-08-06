package main

import (
	"address-verifier/message"
	"address-verifier/verifier"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/get_message", func(c echo.Context) error {
		return c.String(http.StatusOK, "message : "+string(message.GetRandomMessage()))
	})

	e.POST("/verify", func(c echo.Context) error {
		address := c.FormValue("address")
		signedMessage := c.FormValue("signedMessage")
		addressBytes, _ := hex.DecodeString(address)
		signedMessageBytes, _ := hex.DecodeString(signedMessage)

		result := verifier.VerifyAddressPrivateKey(addressBytes, signedMessageBytes)
		var resultStr string
		if result {
			resultStr = "true"
		} else {
			resultStr = "false"
		}

		return c.String(http.StatusOK, "verified : "+resultStr)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
