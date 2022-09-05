package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Provides all functions to execute DB queries
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// Executes a function within a database transaction
// @Params
// @Store -> DB Connection and the queries allowed
// @Context
// @Function -> DB function to execute
// @Returns error. Can be nil or contains an error
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Performs a money transfer between two accounts.
// It will create a transfer record, add account entries and update accounts balance
// @Params
// @Store -> DB Connection and the queries allowed
// @Context
// @TransferTxParams -> Parameters of a transaction
// @Returns TransferTxResult and error
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Call the DB function to create transfer transfer.sql.go
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Call the DB function to create entry entry.sql.go
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount, // Negative amount as his amount decreases with transfers
		})
		if err != nil {
			return err
		}

		// Call the DB function to create entry entry.sql.go
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:      arg.FromAccountID,
			Balance: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:      arg.ToAccountID,
			Balance: arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
