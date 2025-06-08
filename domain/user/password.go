package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const BcryptCost = 10

type Password struct {
	Password string
}

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password{}, fmt.Errorf("password must be at least 8 characters")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return Password{}, fmt.Errorf("failed to hash password: %w", err)
	}

	return Password{
		Password: hashedPassword,
	}, nil
}

func NewPasswordFromSchema(hashedPassword string) Password {
	return Password{
		Password: hashedPassword,
	}
}

func (p Password) IsPasswordMatch(plainPassword []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), plainPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(bytes), err
}
