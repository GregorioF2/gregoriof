package transactions

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gregoriof_challenge/controllers/transactions"
	"github.com/gregoriof_challenge/models"
	"strconv"
	"strings"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if r.Method == "POST" {
		CreateTransaction(w, r)
	} else {
		if strings.Contains(r.URL.Path, "/transactions/") {
			GetTransaction(w, r)
		} else {
			ListTransactions(w, r)
		}
	}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
  var transaction models.Transaction
	err := decoder.Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid json structure. Must be json with 'type' and 'ammount' ", 400)
	}
	createdTransaction, customError := transactions.Create(&transaction)

	if customError != nil {
		http.Error(w, customError.Error.Error(), customError.Status)
		return
	}

	buffer, err := json.Marshal(createdTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buffer)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path: ", r.URL.Path)
	transactionId, err := strconv.ParseUint(r.URL.Path[len("/transactions/"):], 10, 32)
	if err != nil {
		http.Error(w, "Transaction ID must be an unsigned integer", 400)
	}
	transaction, err := transactions.GetById(uint(transactionId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if transaction == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	buffer, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buffer)
}

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	// All this APi to be requests from other domains.

	list, err := transactions.List()
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buffer, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buffer)
}

