package utilities

import (
	"fmt"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	email := "hey@hey.ai"
	// tureEmail := "hey@hey.ai"
	// falseEmail := "thasdhasdhasdhasdhsd@@@asdasfsdfFewf"

	validateEmail := ValidateEmail()(email)
	if validateEmail != true {
		fmt.Printf("Invalide email format: %v\n", validateEmail)
	}
	fmt.Println("Valide Email")
}
