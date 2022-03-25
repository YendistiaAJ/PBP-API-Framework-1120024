package models

import d "test_revel/app/db"

const (
	InsertTransactionSQL            = `INSERT INTO transactions (userid, productid, quantity) VALUES (?, ?, ?)`
	GetTransactionSQL               = `SELECT id, userid, productid, quantity FROM transactions`
	GetTransactionByIdSQL           = GetTransactionSQL + ` WHERE id=?`
	UpdateTransactionSQL            = `UPDATE transactions SET userid=?, productid=?, quantity=? WHERE id=?`
	DeleteTransactionSQL            = `DELETE FROM transactions WHERE id=?`
	DeleteTransactionByUserIdSQL    = `DELETE FROM transactions WHERE userid=?`
	DeleteTransactionByProductIdSQL = `DELETE FROM transactions WHERE productid=?`
)

type Transaction struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func GetTransactionById(id string) (Transaction, int) {
	db := d.Connect()
	defer db.Close()

	var transaction Transaction
	rows, err := db.Query(GetUserByIdSQL, id)
	if err != nil {
		return transaction, 0
	}

	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
			return transaction, 0
		}
	}

	return transaction, 1
}

func DeleteTransactionByUserId(id string) int {
	db := d.Connect()
	defer db.Close()

	_, err := db.Exec(DeleteTransactionByUserIdSQL, id)
	if err != nil {
		return 0
	}

	return 1
}
