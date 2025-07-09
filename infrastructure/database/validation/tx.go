package validation

import (
	"gin-clean-architecture/infrastructure/database/transaction"
	"gorm.io/gorm"
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
