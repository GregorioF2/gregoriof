package transactions

import (
	"github.com/gregoriof/models"
	"github.com/gregoriof/controllers/database"
	"github.com/gregoriof/utils"
	"github.com/stretchr/stew/slice"
	"net/http"
	"errors"
)

// Constant intended array
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

func Create(transaction *models.Transaction) (*models.Transaction, *utils.CustomError) {
	if valid, err := validTransaction(transaction) ; !valid {
		return nil, &utils.CustomError {
			Error: errors.New(err),
			Status: http.StatusBadRequest,
		}
	}
	transactionEntity, err := database.CreateTransaction(transaction)
	if err != nil {
		return nil, &utils.CustomError {
			Error: errors.New("Unexpected error creating transaction"),
			Status: http.StatusInternalServerError,
		}
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