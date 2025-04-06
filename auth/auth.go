package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey") // Replace with env var in prod

// Mock user store (in-memory)
var users = map[string]string{} // username:password

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid Requeted body", http.StatusNotFound)
		return
	}
	if _, exists := users[creds.Username]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	users[creds.Username] = creds.Password
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		http.Error(w, "Invalid Requeted body", http.StatusNotFound)
		return
	}
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
