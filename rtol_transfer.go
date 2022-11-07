package duitku

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// rtolTransfer containts request to transfer amount from merchant to
// customer's bank account, using the previous data obtained from the inquiry
// process.
//
// The formula to generate Signature is SHA256(email + timestamp + bankCode +
// bankAccount + accountName + custRefNumber + amountTransfer + purpose +
// disburseId + secretKey).
//
// Ref: https://docs.duitku.com/disbursement/en/#online-transfer-transfer-request
type rtolTransfer struct {
	// Bank Account owner, obtained after getting a response from the
	// inquiry process.
	AccountName string `json:"accountName"`

	// Destination Bank Code sent when inquiry process.
	BankCode string `json:"bankCode"`

	// Destination account number sent when inquiry procces.
	BankAccount string `json:"bankAccount"`

	// Customer reference number, obtained after getting a response from
	// the inquiry process.
	CustRefNumber string `json:"custRefNumber"`

	// Description of transfer purpose.
	Purpose string `json:"purpose"`

	Request

	// Disbursement transfer amount.
	Amount int64 `json:"amountTransfer"`

	// Disbursement Id provided by Duitku, obtained after getting a
	// response from the inquiry process.
	DisburseID int64 `json:"disburseId"`
}

// newRtolTransfer create new rtolTransfer request from request and response
// of RtolInquiry.
func newRtolTransfer(inquiryReq *RtolInquiry, inquiryRes *RtolInquiryResponse) (req *rtolTransfer) {
	req = &rtolTransfer{
		Amount:  inquiryReq.Amount,
		Purpose: inquiryReq.Purpose,

		AccountName:   inquiryRes.AccountName,
		BankCode:      inquiryRes.BankCode,
		BankAccount:   inquiryRes.BankAccount,
		CustRefNumber: inquiryRes.CustRefNumber,
		DisburseID:    inquiryRes.DisburseID,
	}
	return req
}

func (req *rtolTransfer) sign(opts ClientOptions) {
	var (
		now = time.Now()

		bb        bytes.Buffer
		plainHash [sha256.Size]byte
	)

	req.UserID = opts.UserID
	req.Email = opts.Email
	req.Timestamp = now.UnixMilli()

	fmt.Fprintf(&bb, `%s%d%s%s%s%s%d%s%d%s`, req.Email, req.Timestamp,
		req.BankCode, req.BankAccount, req.AccountName,
		req.CustRefNumber, req.Amount, req.Purpose,
		req.DisburseID, opts.ApiKey)

	plainHash = sha256.Sum256(bb.Bytes())

	req.Signature = hex.EncodeToString(plainHash[:])
}
