package test

import (
	"testing"

	"github.com/cirpi/baselib/db/clickhouse"
)

func TestClickhouseConn(t *testing.T) {
	cfg := clickhouse.Config{
		Host:     "localhost",
		Port:     "9001",
		User:     "root",
		Pass:     "root",
		Database: "test-db",
	}
	_, er := clickhouse.NewClickhouseConn(cfg)
	if er != nil {
		t.Error(er)
	}
}
