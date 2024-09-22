package sqlExt

import (
	"context"
	"database/sql"
)

type ISqlExt interface {
	Ping() error
	Close() error

	BeginTx(ctx context.Context, opts *sql.TxOptions) (context.Context, error)
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}
