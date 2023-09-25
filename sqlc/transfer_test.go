package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Shashwat5522/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: utils.RandomAccountID(),
		ToAccountID:   utils.RandomAccountID(),
		Amount:        utils.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: transfer1.Amount,
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T){
	for i:=0;i<10;i++{
		CreateRandomTransfer(t)
	}
	args:=ListTransfersParams{
		Limit:5,
		Offset:5,
	}
	transfers,err:=testQueries.ListTransfers(context.Background(),args)
	require.NoError(t,err)
	require.Len(t,transfers,5)

	for _,transfer:=range transfers{
		require.NotEmpty(t,transfer)
	}
}