package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"xyz-multifinance/transaction-service/internal/domain"
)

type mysqlTransactionRepository struct {
	db *sql.DB
}

func NewMysqlTransactionRepository(db *sql.DB) *mysqlTransactionRepository {
	return &mysqlTransactionRepository{db: db}
}

func (r *mysqlTransactionRepository) CreateWithLimitUpdate(ctx context.Context, t *domain.Transaction) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil { return fmt.Errorf("gagal memulai transaksi: %w", err) }
	defer tx.Rollback()

	var currentLimit float64
	err = tx.QueryRowContext(ctx, "SELECT amount FROM customer_limits WHERE customer_id = ? AND tenor_in_months = ? FOR UPDATE", t.CustomerID, t.Tenor).Scan(&currentLimit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { return errors.New("konsumen tidak memiliki limit untuk tenor yang dipilih") }
		return fmt.Errorf("gagal mengambil data limit: %w", err)
	}

	if currentLimit < t.OTR { return fmt.Errorf("limit tidak mencukupi (sisa: %v, butuh: %v)", currentLimit, t.OTR) }

	newLimit := currentLimit - t.OTR
	_, err = tx.ExecContext(ctx, "UPDATE customer_limits SET amount = ? WHERE customer_id = ? AND tenor_in_months = ?", newLimit, t.CustomerID, t.Tenor)
	if err != nil { return fmt.Errorf("gagal memperbarui limit: %w", err) }

	_, err = tx.ExecContext(ctx, `INSERT INTO transactions (contract_number, customer_id, otr, admin_fee, installment_amount, interest_amount, asset_name) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		t.ContractNumber, t.CustomerID, t.OTR, t.AdminFee, t.InstallmentAmount, t.InterestAmount, t.AssetName)
	if err != nil { return fmt.Errorf("gagal menyimpan transaksi: %w", err) }

	return tx.Commit()
}