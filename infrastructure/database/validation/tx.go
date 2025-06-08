package validation

import (
	"gorm.io/gorm"
	"kpl-base/infrastructure/database/transaction"
)

func ValidateTransaction(tx interface{}) (*transaction.Repository, error) {
	db, ok := tx.(*transaction.Repository)
	if !ok {
		return nil, gorm.ErrInvalidTransaction
	}

	return db, nil
}
