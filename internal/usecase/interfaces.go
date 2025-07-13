package usecase

import (
	"context"
	"xyz-multifinance/transaction-service/internal/domain"
)

type TransactionRepository interface {
	CreateWithLimitUpdate(ctx context.Context, tx *domain.Transaction) error
}