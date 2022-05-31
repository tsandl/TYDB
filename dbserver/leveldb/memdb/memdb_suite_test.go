package memdb

import (
	"testing"

	"github.com/tsandl/TYDB/dbserver/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
