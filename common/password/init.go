package password

import (

)

type PasswordHashManager interface {
  HashPassword(password string) (string, error)
  ComparePassword(password, hash string) error
}

var _ PasswordHashManager = &passwordHashBcrypt{}

func NewPasswordHashManagerBcrypt() PasswordHashManager {
	return &passwordHashBcrypt{}
}

type passwordHashBcrypt struct {
}
