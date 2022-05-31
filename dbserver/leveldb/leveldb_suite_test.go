package leveldb

import (
	"testing"

	"github.com/tsandl/TYDB/dbserver/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
