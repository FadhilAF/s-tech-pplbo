package password

import (
	"golang.org/x/crypto/bcrypt"
)

func (p passwordHashBcrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p passwordHashBcrypt) ComparePassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
	}
	return nil
}
