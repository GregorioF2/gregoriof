package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gregoriof/routes"
	_ "net/http/pprof"
)


func handleRequest()  {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/transactions", transactions.TransactionsHandler)
	myRouter.HandleFunc("/transactions/{param1}", transactions.TransactionsHandler)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	handleRequest()
}