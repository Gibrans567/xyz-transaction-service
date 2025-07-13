package usecase

import (
	"context"
	"errors"
	"xyz-multifinance/transaction-service/internal/domain"
)

type TransactionUsecase struct {
	repo TransactionRepository
}

func NewTransactionUsecase(repo TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{repo: repo}
}

func (uc *TransactionUsecase) CreateTransaction(ctx context.Context, tx *domain.Transaction) error {
	if tx.CustomerID <= 0 { return errors.New("customer ID tidak valid") }
	if tx.OTR <= 0 { return errors.New("OTR harus lebih besar dari nol") }
	if tx.Tenor <= 0 { return errors.New("tenor tidak valid") }
	if tx.ContractNumber == "" { return errors.New("nomor kontrak tidak boleh kosong") }
	return uc.repo.CreateWithLimitUpdate(ctx, tx)
}