package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Shashwat5522/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T,account Account) Entry {
	arg := CreateEntryParams{
		AccountID:account.ID,
		Amount: utils.RandomMoney(),
		
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {
	account:=createRandomAccount(t)
	CreateRandomEntry(t,account)
}

func TestGetEntry(t *testing.T) {
	account:=createRandomAccount(t)
	entry1:=CreateRandomEntry(t,account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)

}
func TestUpdateEntry(t *testing.T) {
	account:=createRandomAccount(t)
	entry1:=CreateRandomEntry(t,account)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: entry1.Amount,
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)

}

func TestDeleteEntry(t *testing.T) {
	account:=createRandomAccount(t)
	entry1:=CreateRandomEntry(t,account)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)

}

func TestListEntries(t *testing.T) {

	for i := 0; i < 10; i++ {
		account:=createRandomAccount(t)
		CreateRandomEntry(t,account)
	}
	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
