package main

import (
	"log"
	"net/http"

	"github.com/gabrielmenezes/financial-control/handlers"
	"github.com/gorilla/mux"
)

func main(){
  r := mux.NewRouter()

  r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
  r.HandleFunc("/transactions", GetTransactions).Methods("GET")

  log.Println("Server is running on port 8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}

func GetTransactions(w http.ResponseWriter, r *http.Request){
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("List os Transactions"))
}
