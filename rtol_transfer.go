// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// RtolTransfer containts request to transfer amount from merchant to
// customer's bank account, using the previous data obtained from the inquiry
// process.
//
// The formula to generate Signature is SHA256(email + timestamp + bankCode +
// bankAccount + accountName + custRefNumber + amountTransfer + purpose +
// disburseId + secretKey).
//
// Ref: https://docs.duitku.com/disbursement/en/#online-transfer-transfer-request
type RtolTransfer struct {
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

// NewRtolTransfer create new RtolTransfer request from request and response
// of RtolInquiry.
func NewRtolTransfer(inquiryReq *RtolInquiry, inquiryRes *RtolInquiryResponse) (req *RtolTransfer) {
	req = &RtolTransfer{
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

// Sign the request, fill the UserID, Email, Timestamp, and generate the
// Signature.
func (req *RtolTransfer) Sign(opts ClientOptions) {
	var (
		now = time.Now()

		bb        bytes.Buffer
		plainHash [sha256.Size]byte
	)

	req.UserID = opts.DisburseUserID
	req.Email = opts.DisburseEmail
	req.Timestamp = now.UnixMilli()

	fmt.Fprintf(&bb, `%s%d%s%s%s%s%d%s%d%s`, req.Email, req.Timestamp,
		req.BankCode, req.BankAccount, req.AccountName,
		req.CustRefNumber, req.Amount, req.Purpose,
		req.DisburseID, opts.DisburseApiKey)

	plainHash = sha256.Sum256(bb.Bytes())

	req.Signature = hex.EncodeToString(plainHash[:])
}
