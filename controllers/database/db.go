package database

import (
	"github.com/gregoriof/models"
	"sync"
	"time"
	"fmt"
)

var	lockTransactionTable sync.Mutex
var idSeq = 0


var transactions = []models.Transaction{}
var balance = float64(0)

func FindAllTransactions() []models.Transaction {
	return transactions
}

func CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	lockTransactionTable.Lock()
	idSeq += 1 
	transactionEntity := models.Transaction {
		Type: transaction.Type,
		Ammount: transaction.Ammount,
		Id: idSeq,
		EffectiveDate: time.Now().String(),
	}
	transactions = append(transactions, transactionEntity)
	lockTransactionTable.Unlock()
	return &transactionEntity, nil
}

func GetTransactionById(transactionId uint) (*models.Transaction, error) {
	var res *models.Transaction
	lockTransactionTable.Lock()
	fmt.Println("transactionId:", transactionId)
	for _, transaction := range(transactions) {
		fmt.Println("transactions.id:", transaction.Id)
		if transaction.Id == int(transactionId) {
			res = &transaction
			break
		}
	}
	lockTransactionTable.Unlock()
	return res, nil
}