package db

import (
	"context"
	"fmt"

	"github.com/fitv/min/ent"
	_ "github.com/go-sql-driver/mysql"
)

// DB is a wrapper around *ent.Client.
type DB struct {
	client *ent.Client
}

// TxFunc is the function type used by WithTx.
type TxFunc func(tx *ent.Tx) error

// Option represents the options for the DB.
type Option struct {
	Driver string
	Dns    string
	Debug  bool
}

// New returns a new DB instance.
func New(opt *Option) (*DB, error) {
	var entOptions []ent.Option
	if opt.Debug {
		entOptions = append(entOptions, ent.Debug())
	}
	db := &DB{}
	client, err := ent.Open(opt.Driver, opt.Dns, entOptions...)
	if err != nil {
		return nil, err
	}
	db.client = client
	return db, nil
}

// Client returns the underlying *ent.Client.
func (db *DB) Client() *ent.Client {
	return db.client
}

// Close closes the database client.
func (db *DB) Close() error {
	return db.client.Close()
}

// WithTx runs the given function in a transaction.
func (db *DB) WithTx(ctx context.Context, fn TxFunc) error {
	tx, err := db.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
