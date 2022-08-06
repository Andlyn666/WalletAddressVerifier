package verifier

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestVerifyAddressPrivateKeyTrue(t *testing.T) {
	fmt.Printf("TestVerifyAddressPrivateKeyOK Start \n")
	// Generate private key
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		t.Errorf("Private key generate failed")
	}

	// Generate public key
	pubKey := privateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		t.Errorf("Public Key is not of type *ecdsa.PublicKey")
	}

	// get the address from public key
	address := crypto.PubkeyToAddress(*pubKeyECDSA).Bytes()

	// Define msg and get its hash value
	msg := []byte("Random Message")
	msgHash := crypto.Keccak256Hash(msg)

	// sign the message by using the private key
	signature, err := crypto.Sign(msgHash.Bytes(), privateKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	result := VerifyAddressPrivateKey(address, signature)
	if !result {
		t.Errorf("Expect true but got %v", result)
	}
	fmt.Printf("TestVerifyAddressPrivateKeyOK Finished \n")
}
func TestVerifyAddressPrivateKeyFalse(t *testing.T) {
	fmt.Printf("TestVerifyAddressPrivateKeyFail Start \n")
	// Generate private key
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		t.Errorf("Private key generate failed")
	}

	// Generate public key
	pubKey := privateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		t.Errorf("Public Key is not of type *ecdsa.PublicKey")
	}

	// get the address from public key
	address := crypto.PubkeyToAddress(*pubKeyECDSA).Bytes()

	// Define msg and get its hash value
	msg := []byte("Random Message")
	msgHash := crypto.Keccak256Hash(msg)

	// Generate another private key
	privateKeyNew, err := crypto.HexToECDSA("aad9c8855b740a0b7ed4c223dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		t.Errorf("Private key generate new private key failed")
	}
	// sign the message by using the private key
	signature, err := crypto.Sign(msgHash.Bytes(), privateKeyNew)
	if err != nil {
		t.Errorf(err.Error())
	}
	result := VerifyAddressPrivateKey(address, signature)
	if result {
		t.Errorf("Expect false but got %v", result)
	}
	fmt.Printf("TestVerifyAddressPrivateKeyFail Finished \n")
}
