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

func (r Repository) Begin(ctx context.Context) (*Repository, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &Repository{
		db: tx,
	}, nil
}

func (r Repository) CommitOrRollback(ctx context.Context, tx *Repository, err error) {
	if err != nil {
		log.Println("Error occurred, rolling back transaction:", err)
		tx.db.WithContext(ctx).Debug().Rollback()
		return
	}

	err = tx.db.WithContext(ctx).Commit().Error
	if err != nil {
		log.Println("Error committing transaction:", err)
		return
	}

	log.Println("Transaction committed successfully")
}
