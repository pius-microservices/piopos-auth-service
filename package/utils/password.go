package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func CheckPassword(bodyPass, dbPass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(bodyPass))
	
	return err == nil
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	// At least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	// At least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	// At least one number
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false
	}
	// At least one special character
	if !regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\=\{\}\[\]\|\\:;\"'<>,\./?]`).MatchString(password) {
		return false
	}
	return true
}
