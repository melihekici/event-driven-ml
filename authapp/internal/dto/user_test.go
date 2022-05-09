package dto

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	u, err := NewUser("melih", "ekici", "melih.ekici4@gmail.com")
	if err != nil {
		t.Errorf("Creating user failed. %v", err)
	}
	if u.Username != "melih" {
		t.Errorf("Wrong username")
	}
	if u.Email != "melih.ekici4@gmail.com" {
		t.Errorf("Wrong email")
	}
	if u.Password == "ekici" {
		t.Errorf("Password is not hashed")
	}
}

func TestValidatePassword(t *testing.T) {
	u, _ := NewUser("melih", "ekici", "melih.ekici4@gmail.com")
	err := u.validatePassword("ekici")
	if err != nil {
		t.Errorf("Password validation failed %v", err)
	}
}

func TestValidateWrongPassword(t *testing.T) {
	u, _ := NewUser("melih", "ekici", "melih.ekici4@gmail.com")
	err := u.validatePassword("wrong_password")
	if err == nil {
		t.Errorf("Password validation should have failed.")
	}
}

func TestIsValid(t *testing.T) {
	correctUser := User{Username: "melih", Password: "ekici", Email: "melih.ekici4@gmail.com"}
	emptyUser := User{Username: "melih"}
	if !correctUser.isValid() {
		t.Errorf("Validation should have succeeded")
	}
	if emptyUser.isValid() {
		t.Error("Validation should have failed")
	}
}
