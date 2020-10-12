package douyu

import (
	"log"
	"testing"
)

func TestGetToken(t *testing.T) {
	dy, err := New(TestAID, TestKey)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Println(dy.GetToken())
	log.Println(dy.GetToken())
}
