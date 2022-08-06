package message

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRandomMessageGenerate(t *testing.T) {
	fmt.Printf("TestRandomMessageGenerate Start \n")
	msg := GetRandomMessage()
	if !bytes.Equal([]byte("Random Message"), msg) {
		t.Errorf("GetRandomMessage() expect :%v got:%v", "Random Message", msg)
	}
	fmt.Printf("TestRandomMessageGenerate Finished \n")
}
