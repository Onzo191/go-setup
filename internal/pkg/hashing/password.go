package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, pepper string) (string, error) {
	passwordWithPepper := password + pepper
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithPepper), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(hashedPassword, password, pepper string) bool {
	passwordWithPepper := password + pepper
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithPepper))
	return err == nil
}
