package repositories

import (
	"go-api/models"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(models.TransactionRequest) (models.Transaction, error)
}

func RepoTransaction(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) CreateTransaction(transaction models.TransactionRequest) (models.Transaction, error) {
	var trxdata = models.Transaction{
		AdultQty: transaction.AdultQty,
		ChildQty: transaction.ChildQty,
		TicketID: transaction.TicketID,
		BillerID: transaction.BillerID,
	}
	var trx models.Transaction

	e := r.db.Debug().Create(trxdata).First(&trx).Error

	return trx, e

}

func (r *repo) GetTransactions(transaction models.Transaction) ([]models.Transaction, error) {
	var transactions []models.Transaction
	e := r.db.Debug().Find(&transactions).Error
	return transactions, e
}

func (r *repo) GetTransaction(transaction models.Transaction) (models.Transaction, error) {
	e := r.db.Debug().Where("id = ?", transaction.ID).First(&transaction).Error
	return transaction, e
}
