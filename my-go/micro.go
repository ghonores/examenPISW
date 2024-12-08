package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Middleware para habilitar CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Permitir solo desde el frontend
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Responder a solicitudes preflight (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Gonzalo"},
		{ID: 2, Name: "Sebastian"},
		{ID: 3, Name: "UCN"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", getUsers)

	// Aplicar el middleware
	http.ListenAndServe(":8080", enableCORS(mux))
}
