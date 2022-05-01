package redis

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	Init()
	res, err := CheckShare("6ca77f74-a0cd-5ba5-290b-3a2c4ebf04b0", "K.hYe")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", res)
}

func TestStoreShare(t *testing.T) {
	Init()
	err := StoreShare("6ca77f74-a0cd-5ba5-290b-3a2c4ebf04b0", "", 123, 100)
	if err != nil {
		t.Fatal(err)
	}
}
