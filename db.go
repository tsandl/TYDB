package taodb

import "github.com/tsandl/TYDB/leveldb"

type DB interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	State(string) (string, error)
	Iterator(prefix string) (map[string]string, error)
	IteratorOnlyKey(prefix string) ([]string, error)
	CloseDB() error //用于关闭数据库，因为leveldb没有表的概念，我把库上升到表的概念
	OpenDB(dbPath string) (*leveldb.LevelDB, error)
}
