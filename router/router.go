package router

import (
	"github.com/gorilla/mux"
	"github.com/vaibhavkalal/go-postgres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET")
	router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stocks", middleware.CreateStock).Methods("POST")
	router.HandleFunc("/api/stocks/{id}", middleware.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stocks/{id}", middleware.DeleteStock).Methods("DELETE")

	return router
}
