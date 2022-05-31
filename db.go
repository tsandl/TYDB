package taodb

<<<<<<< HEAD
import "github.com/tsandl/TYDB/leveldb"
=======
import (
	"github.com/tsandl/TYDB/dbserver"
)
>>>>>>> version_1.0.0

type DB interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	State(string) (string, error)
	Iterator(prefix string) (map[string]string, error)
	IteratorOnlyKey(prefix string) ([]string, error)
	CloseDB() error //用于关闭数据库，因为leveldb没有表的概念，我把库上升到表的概念
<<<<<<< HEAD
	OpenDB(dbPath string) (*leveldb.LevelDB, error)
=======
	OpenDB(dbPath string) (*dbserver.LevelDB, error)
>>>>>>> version_1.0.0
}
