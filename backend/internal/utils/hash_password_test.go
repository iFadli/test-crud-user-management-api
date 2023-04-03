package utils

import (
	"testing"
)

func TestHashAndComparePassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Errorf("Error hashing password: %s", err)
	}

	if len(hashedPassword) == 0 {
		t.Errorf("Hashed password is empty")
	}

	if !ComparePasswordAndHash(password, hashedPassword) {
		t.Errorf("Password and hash comparison failed")
	}

	if ComparePasswordAndHash("wrongpassword", hashedPassword) {
		t.Errorf("Incorrect password was matched")
	}
}
