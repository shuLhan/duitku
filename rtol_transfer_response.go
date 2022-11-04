package duitku

import "github.com/shuLhan/share/lib/math/big"

// RtolTransferResponse contains response from online transfer.
type RtolTransferResponse struct {
	Email         string   `json:"email"`
	BankCode      string   `json:"bankCode"`
	BankAccount   string   `json:"bankAccount"`
	AccountName   string   `json:"accountName"`
	CustRefNumber string   `json:"custRefNumber"`
	Amount        *big.Rat `json:"amountTransfer"`

	Response
}
