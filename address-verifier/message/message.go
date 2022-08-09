package message

import (
	"crypto/rand"
)

func GetRandomMessage() []byte {
	randomBytes := make([]byte, 20)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil
	}
	return randomBytes
}
