package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn(config Config, opts ...Option) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User,
		config.Pass, config.Host, config.Port, config.Database)
	db, er := sql.Open("mysql", dsn)
	if er != nil {
		return db, er
	}
	er = db.Ping()
	if er != nil {
		return nil, er
	}
	for _, opt := range opts {
		opt(db)
	}
	return db, nil
}
