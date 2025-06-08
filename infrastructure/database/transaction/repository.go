package transaction

import (
	"context"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) DB() *gorm.DB {
	return r.db
}

func (r Repository) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r Repository) CommitOrRollback(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		log.Println("Error occurred, rolling back transaction:", err)
		tx.WithContext(ctx).Debug().Rollback()
		return
	}

	err = tx.WithContext(ctx).Commit().Error
	if err != nil {
		log.Println("Error committing transaction:", err)
		return
	}

	log.Println("Transaction committed successfully")
}
