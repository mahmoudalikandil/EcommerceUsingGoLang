package utilities

import (
	"testing"
)

func TestCheckPassword(t *testing.T) {
	password := "blabla55$"
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("err while hashing password: %v", err)
	}

	err = CheckPassword(password, hash)
	if err != nil {
		t.Fatalf("err while checking password: %v", err)
	}
}
