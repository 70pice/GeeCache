package geecache

type ByteView struct {
	// b存储最真实的缓存之
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return closeBytes(v.b)
}

func closeBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
func (v ByteView) String() string {
	return string(v.b)
}
