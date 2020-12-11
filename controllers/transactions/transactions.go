package transactions

import (
	"challenge/models"
	"challenge/controllers/database"
	"github.com/stretchr/stew/slice"
	"errors"
)

var validTypes = []string{"credit", "debit"}

func validTransaction(transaction *models.Transaction) (bool, string) {
	if !slice.Contains(validTypes, transaction.Type) {
		return false, "Transaction type not valid"
	}
	if transaction.Ammount <= 0 {
		return false, "Transaction ammount must be greater than zero"
	}
	return true, ""
}

func List() ([]models.Transaction, error) {
	transactions := database.FindAllTransactions()
	return transactions, nil
}

func Create(transaction *models.Transaction) (*models.Transaction, error) {
	if valid, err := validTransaction(transaction) ; !valid {
		return nil, errors.New(err)
	}
	transactionEntity, err := database.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	return transactionEntity, nil
}

func GetById(transactionId uint) (*models.Transaction, error) {
	transactionEntity, err := database.GetTransactionById(transactionId)
	if err != nil {
		return nil, err
	}
	return transactionEntity, nil
}