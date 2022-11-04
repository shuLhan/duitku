package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// RtolInquiry contains request to initiate transfer from merchant to
// customer's bank account using [Online Transfer].
//
// The signature formula is SHA256(email + timestamp + bankCode +
// bankAccount + amountTransfer + purpose + apiKey).
//
// [Online Transfer]: https://docs.duitku.com/disbursement/en/#online-transfer-inquiry-request
type RtolInquiry struct {
	// Destination Bank Code.
	BankCode string `json:"bankCode"`

	// Destination account number.
	BankAccount string `json:"bankAccount"`

	// Description of transfer purpose.
	Purpose string `json:"purpose"`

	// Customer name provided by merchant.
	SenderName string `json:"senderName"`

	Request

	// Customer ID provided by merchant.
	SenderID int64 `json:"senderID"`

	// Disbursement transfer amount.
	Amount int64 `json:"amountTransfer"`
}

func (inq *RtolInquiry) sign(apiKey string) {
	var (
		plain = fmt.Sprintf(`%s%d%s%s%d%s%s`, inq.Email,
			inq.Timestamp, inq.BankCode, inq.BankAccount,
			inq.Amount, inq.Purpose, apiKey)
		plainHash [sha256.Size]byte = sha256.Sum256([]byte(plain))
	)
	inq.Signature = hex.EncodeToString(plainHash[:])
}
