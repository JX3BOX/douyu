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

func TestBatchGetRoomInfo(t *testing.T) {
	dy, err := New(TestAID, TestKey)
	if err != nil {
		t.Fatal(err)
		return
	}
	list, err := dy.BatchGetRoomInfo(BatchGetRoomInfoParams{RIds: []int{1889960}})
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	log.Println(list)
}
