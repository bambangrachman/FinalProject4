package repository

import (
	"finalproject4/model"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	GetAllTransactionHistory(transaction *[]model.TransactionHistory) error
	GetTransactionHistoryByUserId(transaction *[]model.TransactionHistory) error
	GetTransactionHistoryByID(id int) (model.TransactionHistory, error)
	CreateTransactionHistory(transactionHistory model.TransactionHistory) (model.TransactionHistory, error)
	DeleteTransactionHistory(transaction model.TransactionHistory) error
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *transactionHistoryRepository {
	return &transactionHistoryRepository{db}
}

func (r *transactionHistoryRepository) GetAllTransactionHistory(transactionHistory *[]model.TransactionHistory) error {
	err := r.db.Preload("Product").Preload("User").Find(&transactionHistory).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionHistoryRepository) GetTransactionHistoryByUserId(transaction *[]model.TransactionHistory) error {
	err := r.db.Preload("Product").Find(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionHistoryRepository) GetTransactionHistoryByID(id int) (model.TransactionHistory, error) {
	var transactionHistory model.TransactionHistory
	err := r.db.Preload("Product").Preload("User").Where("id = ?", id).Find(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}
	return transactionHistory, nil
}

func (r *transactionHistoryRepository) CreateTransactionHistory(transactionHistory model.TransactionHistory) (model.TransactionHistory, error) {
	err := r.db.Create(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}
	return transactionHistory, nil
}

func (r *transactionHistoryRepository) DeleteTransactionHistory(transaction model.TransactionHistory) error {
	err := r.db.Debug().Delete(&transaction).Error
	if err == nil {
		return err
	}
	return err
}
