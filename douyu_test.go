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
	list, err := dy.BatchGetRoomInfo(BatchGetRoomInfoParams{RIds: []int{8852876, 8889134}})
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	log.Println(list)
}
