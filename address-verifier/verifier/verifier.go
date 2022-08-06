package verifier

import (
	"bytes"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func VerifyAddressPrivateKey(addr []byte, sig []byte) bool {
	// Define msg and get its hash value
	msg := []byte("Random Message")
	msgHash := crypto.Keccak256Hash(msg)

	// Get public key from signatrue
	pubKeyECDSA, err := crypto.SigToPub(msgHash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	// Get the address from public key
	addressFromSig := crypto.PubkeyToAddress(*pubKeyECDSA).Bytes()

	// Match the generate adress and input adress
	matches := bytes.Equal(addr, addressFromSig)
	return matches
}
