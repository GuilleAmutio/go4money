package db

import (
	"context"
	"testing"
	"time"

	"github.com/guilleamutio/go4money/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListAllTransfers(t *testing.T) {
	transfer := createRandomTransfer(t)

	transfers, err := testQueries.ListAllTransfers(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Contains(t, transfers, transfer)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestListTransfersFromAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		arg1 := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg1)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
		require.Equal(t, arg1.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg1.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg1.Amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

	}

	transfers, err := testQueries.ListTransfersFromAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, transfer.FromAccountID, account1.ID)
	}
}

func TestListTransfersToAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		arg1 := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg1)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
		require.Equal(t, arg1.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg1.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg1.Amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
	}

	transfers, err := testQueries.ListTransfersToAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, transfer.ToAccountID, account2.ID)
	}
}
