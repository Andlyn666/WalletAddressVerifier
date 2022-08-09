package verifier

import (
	"bytes"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func VerifyAddressPrivateKey(addr []byte, sig []byte, msg []byte) bool {
	// Define msg and get its hash value
	msgHash := crypto.Keccak256Hash(msg)

	// Get public key from signatrue
	pubKeyECDSA, err := crypto.SigToPub(msgHash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Get the address from public key
	addressFromSig := crypto.PubkeyToAddress(*pubKeyECDSA).Bytes()

	// Match the generate address and input address
	matches := bytes.Equal(addr, addressFromSig)
	return matches
}
