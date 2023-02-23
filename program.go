//Write API Endpoints

// Path: api.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "
			
			",
		}
		json.NewEncoder(w).Encode(user)
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
