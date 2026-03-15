package sql

import "database/sql"

type Option func(*sql.DB)

func WithMaxIdle(n int) Option {
	return func(d *sql.DB) {
		d.SetMaxIdleConns(n)
	}
}

func WithMaxConn(n int) Option {
	return func(d *sql.DB) {
		d.SetMaxOpenConns(n)
	}
}
