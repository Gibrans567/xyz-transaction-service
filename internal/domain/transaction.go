package domain

import "time"

type Transaction struct {
	ID                int64     `json:"id"`
	ContractNumber    string    `json:"contract_number"`
	CustomerID        int64     `json:"customer_id"`
	OTR               float64   `json:"otr"`
	AdminFee          float64   `json:"admin_fee"`
	InstallmentAmount float64   `json:"installment_amount"`
	InterestAmount    float64   `json:"interest_amount"`
	AssetName         string    `json:"asset_name"`
	Tenor             int       `json:"tenor"`
	TransactionDate   time.Time `json:"transaction_date"`
}
