package table

import (
	"testing"

	"github.com/tsandl/TYDB/dbserver/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
