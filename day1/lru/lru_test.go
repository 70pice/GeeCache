package lru

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"
)

type String string

type Student struct {
	x int
}

func (s Student) hello() {
	fmt.Println("hello", s.x)
	s.x++
}
func (s *Student) hello2() {
	fmt.Println("hello", s.x)
	s.x++
}
func (d String) Len() int {
	return len(d)
}

func TestAdd(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("qwe"))
	lru.Add("key2", String("asd"))
	lru.Add("key3", String("zxc"))
	for item := lru.ll.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("qwe"))
	lru.Add("key2", String("asd"))
	lru.Add("key3", String("zxc"))
	lru.RemoveOldest()
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("qwe"))
	lru.Add("key2", String("asd"))
	lru.Add("key3", String("zxc"))

	lru.Get("key1")
}

func TestCache_Add(t *testing.T) {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))
}

func add1(r rune) rune { return r + 1 }

func TestCache_Len(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
