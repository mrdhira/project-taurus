package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

func (m *mysqlExt) Ping() error {
	return m.dbConn.Ping()
}

func (m *mysqlExt) closeConn(dbConn *sqlx.DB) error {
	if dbConn != nil {
		return dbConn.Close()
	}

	return nil
}

func (m *mysqlExt) Close() error {
	errG := new(errgroup.Group)

	errG.Go(func() error {
		return m.closeConn(m.dbConn)
	})

	errG.Go(func() error {
		return m.closeConn(m.dbReadOnlyConn)
	})

	return errG.Wait()
}

func (m *mysqlExt) BeginTx(ctx context.Context, opts *sql.TxOptions) (context.Context, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Check if transaction already exist
	// If exist, return the context, to prevent nested transaction
	if _, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx); ok {
		return ctx, nil
	}

	tx, err := m.dbConn.BeginTx(ctx, opts)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, CtxKeySqlTx, tx), nil
}

func (m *mysqlExt) BeginTxx(ctx context.Context, opts *sql.TxOptions) (context.Context, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Check if transaction already exist
	// If exist, return the context, to prevent nested transaction
	if _, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx); ok {
		return ctx, nil
	}

	tx, err := m.dbConn.BeginTxx(ctx, opts)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, CtxKeySqlxTx, tx), nil
}

func (m *mysqlExt) Commit(ctx context.Context) error {
	// Check if sql tx exist
	tx, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx)
	if ok {
		return tx.Commit()
	} else {
		// Check if sqlx tx exist
		tx, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx)
		if ok {
			return tx.Commit()
		}
	}

	return nil
}

func (m *mysqlExt) Rollback(ctx context.Context) error {
	// Check if sql tx exist
	tx, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx)
	if ok {
		return tx.Rollback()
	} else {
		// Check if sqlx tx exist
		tx, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx)
		if ok {
			return tx.Rollback()
		}
	}

	return nil
}

func (m *mysqlExt) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	// Check if sql tx exist
	tx, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx)
	if ok {
		return tx.GetContext(ctx, dest, query, args...)
	} else {
		// Check if sqlx tx exist
		tx, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx)
		if ok {
			return tx.GetContext(ctx, dest, query, args...)
		}
	}

	return m.dbReadOnlyConn.GetContext(ctx, dest, query, args...)
}

func (m *mysqlExt) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	// Check if sql tx exist
	tx, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx)
	if ok {
		return tx.SelectContext(ctx, dest, query, args...)
	} else {
		// Check if sqlx tx exist
		tx, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx)
		if ok {
			return tx.SelectContext(ctx, dest, query, args...)
		}
	}

	return m.dbReadOnlyConn.SelectContext(ctx, dest, query, args...)
}

func (m *mysqlExt) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Check if sql tx exist
	tx, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx)
	if ok {
		return tx.ExecContext(ctx, query, args...)
	} else {
		// Check if sqlx tx exist
		tx, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx)
		if ok {
			return tx.ExecContext(ctx, query, args...)
		}
	}

	return m.dbConn.ExecContext(ctx, query, args...)
}

func (m *mysqlExt) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	// Check if sql tx exist
	tx, ok := ctx.Value(CtxKeySqlTx).(*sqlx.Tx)
	if ok {
		return tx.NamedExecContext(ctx, query, arg)
	} else {
		// Check if sqlx tx exist
		tx, ok := ctx.Value(CtxKeySqlxTx).(*sqlx.Tx)
		if ok {
			return tx.NamedExecContext(ctx, query, arg)
		}
	}

	return m.dbConn.NamedExecContext(ctx, query, arg)
}
