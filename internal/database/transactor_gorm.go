package database

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type Transactor interface {
	Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error
}

type transactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) Transactor {
	return &transactor{db: db}
}

var txOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault, // Default db isolation level
	ReadOnly:  false,            // Allow writes in the transaction
}

func (t *transactor) Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		err := txFunc(injectTx(ctx, tx))
		if err != nil {
			return err
		}

		return nil
	}, txOptions)
	if err != nil {
		return err
	}

	return nil
}
