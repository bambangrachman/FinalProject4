package service

import (
	"errors"
	"finalproject4/model"
	"finalproject4/repository"
)

type TransactionHistoryService interface {
	GetAllTransactionHistory(role string, UserID int) ([]model.TransactionHistory, error)
	GetTransactionHistoryByUserId(UserID int) ([]model.TransactionHistory, error)
	CreateTransactionHistory(transactionHistory model.TransactionHistoryInput, UserID int) (model.TransactionHistory, error)
	DeleteTransactionHistory(id_task int, UserID int) error
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository}
}

func (s *transactionHistoryService) GetAllTransactionHistory(role string, UserID int) ([]model.TransactionHistory, error) {
	var transactionHistory []model.TransactionHistory
	if role == "admin" {
		err := s.transactionHistoryRepository.GetAllTransactionHistory(&transactionHistory)
		if len(transactionHistory) == 0 {
			return transactionHistory, errors.New("no transaction history")
		}
		if err != nil {
			return []model.TransactionHistory{}, err
		}
		return transactionHistory, nil
	}
	if role == "customer" {
		err := s.transactionHistoryRepository.GetAllTransactionHistory(&transactionHistory)
		if len(transactionHistory) == 0 {
			return transactionHistory, errors.New("no transaction history")
		}
		if err != nil {
			return []model.TransactionHistory{}, err
		}
		var filteredTransactionHistory []model.TransactionHistory
		for _, transaction := range transactionHistory {
			if transaction.UserID == UserID {
				filteredTransactionHistory = append(filteredTransactionHistory, transaction)
			}
		}
		return filteredTransactionHistory, nil
	}
	return []model.TransactionHistory{}, errors.New("Login first only user or admin role access this")
}

func (s *transactionHistoryService) GetTransactionHistoryByUserId(UserID int) ([]model.TransactionHistory, error) {
	var transaction []model.TransactionHistory
	err := s.transactionHistoryRepository.GetTransactionHistoryByUserId(&transaction)
	if err != nil {
		return transaction, err
	}
	if len(transaction) == 0 {
		return transaction, errors.New("no transaction history")
	}
	return transaction, nil
}

func (s *transactionHistoryService) CreateTransactionHistory(transaction model.TransactionHistoryInput, UserID int) (model.TransactionHistory, error) {

	transactionHistory := model.TransactionHistory{
		UserID:     UserID,
		ProductID:  transaction.ProductID,
		Quantity:   transaction.Quantity,
		TotalPrice: 0,
	}
	transactionResponse, err := s.transactionHistoryRepository.CreateTransactionHistory(transactionHistory)
	if err != nil {
		return transactionResponse, err
	}
	return transactionResponse, nil
}

func (s *transactionHistoryService) DeleteTransactionHistory(id_task int, UserID int) error {
	transaction, err := s.transactionHistoryRepository.GetTransactionHistoryByID(id_task)
	if err != nil {
		return err
	}
	if transaction.UserID != UserID {
		return errors.New("you are not the owner of this transaction history")
	}
	err = s.transactionHistoryRepository.DeleteTransactionHistory(transaction)

	if err != nil {
		return err
	}
	return nil
}
