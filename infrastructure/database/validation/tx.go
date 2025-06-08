package validation

import (
	"gorm.io/gorm"
	"kpl-base/infrastructure/database/transaction"
)

func ValidateTransaction(tx interface{}) (*transaction.Repository, error) {
	db := &transaction.Repository{}
	if tx == nil {
		return db, nil
	}

	db, ok := tx.(*transaction.Repository)
	if !ok {
		return nil, gorm.ErrInvalidTransaction
	}

	return db, nil
}
