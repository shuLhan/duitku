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

// ClearingTransfer contains request parameter for Clearing Transfer.
//
// Formula to generate signature:
//
//	SHA256(email + timestamp + bankCode + type + bankAccount +
//		accountName + custRefNumber + amountTransfer + purpose +
//		disburseId + apiKey)
//
// Ref: https://docs.duitku.com/disbursement/en/#clearing-transfer-request
type ClearingTransfer struct {
	Type string `json:"type"`

	RtolTransfer
}

// NewClearingTransfer create ClearingTransfer from Clearing Inquiry
// request and response.
//
// The following fields are set from response: AccountName, CustRefNumber,
// DisburseID, and Type.
func NewClearingTransfer(inqReq *ClearingInquiry, inqRes *ClearingInquiryResponse) (trf *ClearingTransfer) {
	trf = &ClearingTransfer{}

	trf.Amount = inqReq.Amount
	trf.BankAccount = inqReq.BankAccount
	trf.BankCode = inqReq.BankCode
	trf.Purpose = inqReq.Purpose

	trf.AccountName = inqRes.AccountName
	trf.CustRefNumber = inqRes.CustRefNumber
	trf.DisburseID = inqRes.DisburseID
	trf.Type = inqRes.Type

	return trf
}

// Sign the request, fill the UserID, Email, Timestamp, and generate the
// Signature.
func (trf *ClearingTransfer) Sign(opts ClientOptions) {
	var (
		now = time.Now()

		bb        bytes.Buffer
		plainHash [sha256.Size]byte
	)

	trf.UserID = opts.UserID
	trf.Email = opts.Email
	trf.Timestamp = now.UnixMilli()

	bb.WriteString(trf.Email)
	fmt.Fprintf(&bb, `%d`, trf.Timestamp)
	bb.WriteString(trf.BankCode)
	bb.WriteString(trf.Type)
	bb.WriteString(trf.BankAccount)
	bb.WriteString(trf.AccountName)
	bb.WriteString(trf.CustRefNumber)
	fmt.Fprintf(&bb, `%d`, trf.Amount)
	bb.WriteString(trf.Purpose)
	fmt.Fprintf(&bb, `%d`, trf.DisburseID)
	bb.WriteString(opts.ApiKey)

	plainHash = sha256.Sum256(bb.Bytes())

	trf.Signature = hex.EncodeToString(plainHash[:])
}
