package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/players", getPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", getPlayer).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("App: Manchester United Players API\nAuthor:Sagataly Aidyn"))
	}).Methods("GET")
	http.ListenAndServe(":8080", router)
}
