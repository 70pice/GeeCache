package geecache

// 用于根据传入的key选择相对应的节点
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 从对应的group查找缓存值
type PeerGetter interface {
	Get(Group string, key string) ([]byte, error)
}
