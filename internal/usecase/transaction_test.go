package usecase

import (
	"context"
	"errors"
	"testing"
	"xyz-multifinance/transaction-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) CreateWithLimitUpdate(ctx context.Context, tx *domain.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func TestCreateTransaction_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	uc := NewTransactionUsecase(mockRepo)
	testTx := &domain.Transaction{CustomerID: 1, OTR: 50000, Tenor: 3, ContractNumber: "XYZ123"}
	mockRepo.On("CreateWithLimitUpdate", mock.Anything, testTx).Return(nil)
	err := uc.CreateTransaction(context.Background(), testTx)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateTransaction_RepositoryError(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	uc := NewTransactionUsecase(mockRepo)
	testTx := &domain.Transaction{CustomerID: 1, OTR: 1000000, Tenor: 6, ContractNumber: "XYZ999"}
	repoError := errors.New("limit tidak mencukupi")
	mockRepo.On("CreateWithLimitUpdate", mock.Anything, testTx).Return(repoError)
	err := uc.CreateTransaction(context.Background(), testTx)
	assert.Error(t, err)
	assert.Equal(t, repoError, err)
	mockRepo.AssertExpectations(t)
}