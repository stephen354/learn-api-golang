package service

import (
	"belajar-api/model"
	"belajar-api/repository"
)

type TransactionService interface {
	Checkout(items []model.CheckoutItem) (*model.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
