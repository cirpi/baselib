package test

import (
	"testing"

	"github.com/cirpi/baselib/db/sql"
)

func TestMysqlConn(t *testing.T) {
	cfg := sql.Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Pass:     "root",
		Database: "test_db",
	}
	_, er := sql.NewMysqlConn(cfg)
	if er != nil {
		t.Error(er)
	}
}
