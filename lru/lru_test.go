package lru

import (
	"fmt"
	"testing"
)

type testStruct struct {
}

func (testStruct *testStruct) len() int {
	return 1
}

func TestLRU_Add(t *testing.T) {
	cache := New(int64(0), nil)
	fmt.Print(cache)
}
