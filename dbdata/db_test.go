package dbdata

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yyxhgy/anylink-server/base"
)

func preIpData() {
	tmpDb := path.Join(os.TempDir(), "anylink_test.db")
	base.Cfg.DbType = "sqlite3"
	base.Cfg.DbSource = tmpDb
	initDb()
}

func closeIpdata() {
	xdb.Close()
	tmpDb := path.Join(os.TempDir(), "anylink_test.db")
	os.Remove(tmpDb)
}

func TestDb(t *testing.T) {
	ast := assert.New(t)
	preIpData()
	defer closeIpdata()

	u := User{Username: "a"}
	err := Add(&u)
	ast.Nil(err)

	ast.Equal(u.Id, 1)
}
