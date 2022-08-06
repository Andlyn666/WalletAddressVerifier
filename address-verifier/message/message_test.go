package message

import (
	"fmt"
	"testing"
)

func TestRandomMessageGenerate(t *testing.T) {
	msg := GetRandomMessage()
	if msg != "Random Message" {
		t.Errorf("GetRandomMessage() expect :%v got:%v", "Random Message", msg)
	}
}
