package gosayhello_test

import (
	"fmt"
	"github.com/yaohaihan/gosayhello"
	"testing"
)

func TestHello(t *testing.T) {
	data := "jack"
	expect := fmt.Sprintf("Hello %s", data)

	result := gosayhello.SayHello(data)

	if result != expect {
		t.Fatalf("expected result %s, but got %s", expect, result)

	}
}
