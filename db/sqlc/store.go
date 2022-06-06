package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
	// create a new database transaction
}

// create Store
func NewStore(db *sql.DB) *Store {
	 return &Store{
		 db: db,
		 Queries: New(db),
	 }
}

// excute a generic database transaction
func (store *Store) execTx(ctx context.Context, fn func (*Queries) error) error {
	// return transaction or error
	tx,err := store.db.BeginTx(ctx,nil)
	if err != nil {
		return err
	}

	//created transaction
	q := New(tx)

	//check error queries
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback();rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v",err,rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amout int64 `json:"amount"`
}

type TransferResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEnTry Entry `json:"to_entry"`
}

var txKey = struct{}{}
// bracket two is empty object we created

//TransferTx performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and update accounts balance within a single database transaction 
func(store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferResult,error) {
	var result TransferResult


	// create transaction record
	err := store.execTx(ctx,func(q *Queries) error {
		var err error
		// txName := ctx.Value(txKey)
		// fmt.Println(txName,"CREATE TRANSFER")
		result.Transfer,err = q.CreateTransfer(ctx,CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amout,
		})
		if err != nil {
			return err
		}
	
		// add 2 entries from - to and money 
		// fmt.Println(txName,"CREATE FROM ENTRY")
		result.FromEntry,err = q.CreateEntry(ctx,CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amout,
		})
		if err != nil {
			return err
		}
		// fmt.Println(txName,"CREATE TO ENTRY")
		result.ToEnTry, err = q.CreateEntry(ctx,CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amout,
		})
		if err != nil {
			return nil;
		}

		// GET - UPDATE ACOOUNT Balance -> Solution 1

		// fmt.Println(txName,"GET ACCOUNT 1")
		// account1,err := q.GetAccountForUpdate(ctx,arg.FromAccountID)
		// if(err != nil) {
		// 	return err
		// }
		// fmt.Println(txName,"UPDATE ACCOUNT 1")
		// result.FromAccount,err = q.UpdateAccount(ctx,UpdateAccountParams{
		// 	ID: arg.FromAccountID,
		// 	Balance: account1.Balance - arg.Amout,
		// })
		// if(err != nil) {
		// 	return err
		// }
		// fmt.Println(txName,"GET ACCOUNT 2")
		// account2,err := q.GetAccountForUpdate(ctx,arg.ToAccountID)
		// if(err != nil) {
		// 	return err
		// }
		// fmt.Println(txName,"UPDATE ACCOUNT 2")
		// result.ToAccount,err = q.UpdateAccount(ctx,UpdateAccountParams{
		// 	ID: arg.ToAccountID,
		// 	Balance: account2.Balance + arg.Amout,
		// })
		// if(err != nil) {
		// 	return err
		// }


		// 1 Query for solution 2:

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount,result.ToAccount,err = addMoney(ctx,q,arg.FromAccountID,-arg.Amout,arg.ToAccountID,arg.Amout)
			// result.FromAccount,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
			// 	ID: arg.FromAccountID,
			// 	Amount: -arg.Amout,
			// })
			// if err != nil {
			// 	return err
			// }
	
			// result.ToAccount,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
			// 	ID: arg.ToAccountID,
			// 	Amount: arg.Amout,
			// })
			// if err != nil {
			// 	return err
			// }
			
		} else {
			// result.ToAccount,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
			// 	ID: arg.ToAccountID,
			// 	Amount: arg.Amout,
			// })
			// if err != nil {
			// 	return err
			// }
			// result.FromAccount,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
			// 	ID: arg.FromAccountID,
			// 	Amount: -arg.Amout,
			// })
			// if err != nil {
			// 	return err
			// }
			result.ToAccount,result.FromAccount,err = addMoney(ctx,q,arg.ToAccountID,arg.Amout,arg.FromAccountID,-arg.Amout)
		}
		return nil
	})

	return result,err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account , account2 Account, err error) {
	account1,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
		ID: accountID1,
		Amount: amount1,
	})
	if(err != nil) {
		return 
	}

	account2,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
		ID: accountID2,
		Amount: amount2,
	})
	if(err != nil) {
		return 
	}
	return
}