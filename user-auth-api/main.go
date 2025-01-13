package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string) // username -> hashed password

var jwtSecret = []byte("jwt_secret_key")

// generateToken generates a JWT token for the given username.
// The token contains the username as a claim and an expiration time set to 1 hour from the current time.
// It returns the signed token string or an error if signing fails.
//
// Parameters:
//   - username: The username to include in the token claims.
//
// Returns:
//   - string: The signed JWT token.
//   - error: An error if the token signing fails.
func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// validateToken parses and validates a JWT token string.
// It returns the parsed token if the token is valid, or an error if the token is invalid or parsing fails.
//
// Parameters:
//   - tokenString: The JWT token string to be validated.
//
// Returns:
//   - *jwt.Token: The parsed JWT token if validation is successful.
//   - error: An error if the token is invalid or parsing fails.
func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

// registerHandler handles user registration requests. It expects a JSON payload
// containing a username and password. If the username already exists, it returns
// a 409 Conflict status. If the input is invalid, it returns a 400 Bad Request status.
// On successful registration, it stores the user and returns a 201 Created status.
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if username already exists
	if _, exists := users[user.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Hash the password before storing (use bcrypt in a real app)
	users[user.Username] = user.Password

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// loginHandler handles user login requests.
// It expects a JSON payload with username and password, decodes it into a User struct,
// and verifies the credentials against a stored user map.
// If the credentials are valid, it generates a JWT token and returns it in the response.
// If the credentials are invalid or there is an error during processing, it returns an appropriate HTTP error.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Verify username and password
	if storedPassword, exists := users[user.Username]; !exists || storedPassword != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := generateToken(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get token from header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	token, err := validateToken(authHeader)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	w.Write([]byte(fmt.Sprintf("Welcome, %s!", username)))
}

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
