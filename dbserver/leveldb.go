package dbserver

import (
	"fmt"
	_leveldb "github.com/tsandl/TYDB/dbserver/leveldb"
	"github.com/tsandl/TYDB/dbserver/leveldb/iterator"
	opt2 "github.com/tsandl/TYDB/dbserver/leveldb/opt"
	"github.com/tsandl/TYDB/dbserver/leveldb/util"
	"github.com/tsandl/TYDB/s3"
)

const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
)

type LevelDB struct {
	DB *_leveldb.DB
}

func NewDB(dbPath string) *LevelDB {

	opt := new(opt2.Options)
	opt.CompactionTableSize = 4 * MiB
	opt.IteratorSamplingRate = 2 * MiB
	opt.WriteBuffer = 32 * MiB
	opts := s3.OpenOption{
		Bucket:        "tsandyzf",
		Path:          dbPath,
		Endpoint:      "http://127.0.0.1:9000",
		Ak:            "admin",
		Sk:            "admin123456",
		Region:        "us-east-1",
		LocalCacheDir: "/",
	}
	st, err := s3.NewS3Storage(opts)
	if err != nil {
		panic(err)
	}
	//dbInstance, err := _leveldb.OpenFile(dbPath, opt)
	dbInstance, err := _leveldb.Open(st, opt)

	if err != nil {
		panic(err)
	}
	return &LevelDB{dbInstance}
}

func (db *LevelDB) Set(key string, value []byte) error {
	return db.DB.Put([]byte(key), []byte(value), nil)
}

func (db *LevelDB) Get(key string) ([]byte, error) {
	data, err := db.DB.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *LevelDB) Del(key string) error {
	return db.DB.Delete([]byte(key), nil)
}

func (db *LevelDB) State(value string) (string, error) {
	if value == "" {
		value = "leveldb.stats"
	}
	if value == "type" {
		return "leveldb", nil
	}

	return db.DB.GetProperty(value)
}

func (db *LevelDB) Iterator(prefix string) (map[string]string, error) {
	data := make(map[string]string)
	var iter iterator.Iterator
	if prefix == "" {
		iter = db.DB.NewIterator(nil, nil)
		for ok := iter.Seek([]byte("")); ok; ok = iter.Next() {
			data[string(iter.Key())] = string(iter.Value()[:])
		}
	} else {
		iter = db.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
		for iter.Next() {
			data[string(iter.Key())] = string(iter.Value()[:])
		}
	}
	iter.Release()
	return data, iter.Error()
}

func (db *LevelDB) IteratorOnlyKey(prefix string) ([]string, error) {
	data := make([]string, 0)
	var iter iterator.Iterator
	if prefix == "" {
		iter = db.DB.NewIterator(nil, nil)
		for ok := iter.Seek([]byte("")); ok; ok = iter.Next() {
			data = append(data, string(iter.Key()))
		}
	} else {
		iter = db.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
		for iter.Next() {
			data = append(data, string(iter.Key()))
		}
	}
	iter.Release()
	return data, iter.Error()
}

func (db *LevelDB) CloseDB() error {
	err := db.DB.Close()
	if err != nil {
		fmt.Println("fail close db")
	} else {
		fmt.Println("success close db")
	}
	return err
}

func (db *LevelDB) OpenDB(dbPath string) (*LevelDB, error) {
	opt := new(opt2.Options)
	opt.CompactionTableSize = 4 * MiB
	opt.IteratorSamplingRate = 2 * MiB
	opt.WriteBuffer = 32 * MiB

	dbInstance, err := _leveldb.OpenFile(dbPath, opt)
	if err != nil {
		panic(err)
	}
	return &LevelDB{dbInstance}, err
}
