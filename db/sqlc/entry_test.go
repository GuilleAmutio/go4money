package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/guilleamutio/go4money/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	arg1 := ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg1)
	require.NoError(t, err)

	// Select random ID from
	k := rand.Int63n(10 - 1)
	accountID := accounts[k].ID

	arg2 := CreateEntryParams{
		AccountID: accountID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg2)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg2.AccountID, entry.AccountID)
	require.Equal(t, arg2.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, account := range entries {
		require.NotEmpty(t, account)
	}
}

func TestListEntriesBetweenDates(t *testing.T) {
	time1 := time.Now()

	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	time2 := time.Now()

	arg := ListEntriesBetweenDatesParams{
		CreatedAt:   time1,
		CreatedAt_2: time2,
	}

	entries, err := testQueries.ListEntriesBetweenDates(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 10)

	for _, entry := range entries {
		require.GreaterOrEqual(t, entry.CreatedAt, time1)
		require.LessOrEqual(t, entry.CreatedAt, time2)
	}
}
