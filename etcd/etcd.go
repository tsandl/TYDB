package etcd

type Etcd interface {
	put(key, value string) error
	get(key string) []byte
	delete(key string) error
	deleteWithPrefix(prefix string) error
	getWithPrefix(prefix string) []byte
	watch(key string)
	watchWithPrefix(prefix string)
	lease(key, value string, ttl int64)
	keepAlive(key, value string, ttl int64)
}
