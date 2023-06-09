package consistenthansh

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

type student struct {
}

func TestMap_Add(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	hash.Add("6", "4", "2")

	testCase := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"37": "2",
	}
	for k, v := range testCase {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s,should have yielded %s", k, v)
		}
	}
	hash.Add("8")

	testCase["27"] = "8"

	for k, v := range testCase {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s,should have yielded %s", k, v)
		}
	}
}

func TestMap_Get(t *testing.T) {
	fmt.Println(10 >> 1)
	search := sort.Search(10, func(i int) bool {
		return i > 2
	})
	fmt.Println(search)
}
