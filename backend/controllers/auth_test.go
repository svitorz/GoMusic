package controllers_test

import (
	"testing"

	"github.com/svitorz/GoMusic/backend/controllers"
)

func TestAuth(t *testing.T) {
	t.Run("TestLogin", TestEncryptPassword)
	t.Run("TestLogin", TestVerifyPassword)
}

func TestEncryptPassword(t *testing.T) {
	password := "mysecretpassword"
	encryptedPassword, err := controllers.HashPassword(password)
	if err != nil {
		t.Errorf("Error encrypting password: %v", err)
	}
	if encryptedPassword == "" {
		t.Error("Encrypted password is empty")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "mysecretpassword"
	encryptedPassword, err := controllers.HashPassword(password)
	if err != nil {
		t.Errorf("Error encrypting password: %v", err)
	}
	isValid := controllers.VerifyPassword(encryptedPassword, password)
	if !isValid {
		t.Error("Password verification failed")
	}
	isInvalid := controllers.VerifyPassword(encryptedPassword, "wrongpassword")
	if isInvalid {
		t.Error("Password verification should have failed for wrong password")
	}
}
