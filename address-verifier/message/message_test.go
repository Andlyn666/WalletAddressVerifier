package message

import (
	"fmt"
	"testing"
)

func TestRandomMessageGenerate(t *testing.T) {
	fmt.Printf("TestRandomMessageGenerate Start \n")
	msg := GetRandomMessage()
	if msg == nil {
		t.Errorf("GetRandomMessage() return nil")
	}
	fmt.Printf("TestRandomMessageGenerate Finished\nMessage : %v \n", msg)
}
