package main

import (
	"fmt"

	"github.com/apaarshrm39/learningGO/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Apaar",
		LastName:  "Sharma",
	}
	fmt.Println(u)
}
