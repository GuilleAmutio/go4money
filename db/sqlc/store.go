package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Provides all functions to execute DB queries
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
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
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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
func (store *SQLStore) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
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

		// Avoid a deadlock scenario by ordering the queries
		if arg.FromAccountID < arg.ToAccountID {
			// Update FromAccount before ToAccount
			result.FromAccount, result.ToAccount, _ = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			// Update ToAccount before FromAccount
			result.ToAccount, result.FromAccount, _ = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return nil
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1 int64, balance1 int64, accountID2 int64, balance2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:      accountID1,
		Balance: balance1,
	})
	if err != nil {
		// No neeed to specify the returned values as named returns are used
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:      accountID2,
		Balance: balance2,
	})
	if err != nil {
		// No neeed to specify the returned values as named returns are used
		return
	}

	return
}
