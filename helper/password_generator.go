package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePassword(databasePassword, payloadPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(payloadPassword))
	if err != nil {
		fmt.Println("helper error:", err)
		return err
	}

	return nil
}
