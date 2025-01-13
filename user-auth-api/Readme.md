# Project: User Authentication API

## What We Learned

### 1. JSON Web Tokens (JWT)

#### **JWT Components:**

Header: Specifies the algorithm used (e.g., HS256).
Payload: Contains user-specific data (e.g., username) and claims like expiration time (exp).

Signature: Ensures the token is tamper-proof using a secret key.

#### **JWT Workflow:**

Server generates a JWT with user information and signs it with a secret key.
Client stores the token (e.g., in local storage) and sends it in the Authorization header.

Server validates the token for protected endpoints.

### 2. Token Generation

Purpose: Create a JWT after a user logs in successfully.

Code Highlights:
Define claims (e.g., username, exp).
Sign the token with a secret key using the jwt library:

```go
func generateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 1).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
```

Expiration: Ensures the token becomes invalid after a set duration (1 hour in this case).

### 3. Token Validation

Purpose: Verify the integrity of the token and extract user information.
Code Highlights:
Parse and validate the token using the secret key:

```go
func validateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
}
```

Validation Process:
Verify the signature matches the secret key.
Check the exp (expiration) claim.

### 4. Protected Endpoints

Purpose: Restrict access to certain routes using token validation.
Code Highlights:
Extract the token from the Authorization header.
Validate the token and retrieve claims (e.g., username):

```go
func protectedHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    token, err := validateToken(authHeader)
    claims := token.Claims.(jwt.MapClaims)
    username := claims["username"].(string)
    w.Write([]byte(fmt.Sprintf("Welcome, %s!", username)))
}
```

Security:
Return 401 Unauthorized if the token is invalid or missing.

### 5. Secure User Management

User Registration:
Accept username and password from the client.
Hash the password before storing it (use libraries like bcrypt in production).
User Login:
Verify stored credentials with client-provided data.
Generate a JWT for successful logins.

## New Concepts Implemented

### **JWT Authentication:**

Stateless authentication for secure communication.
Token expiration to minimize misuse.

### **Protected Routes:**

Restricted access using token validation.

### **Error Handling:**

Meaningful HTTP error codes (400, 401, 409) for various scenarios.

### **In-Memory User Store:**

Simulated a database using a Go map for simplicity.
