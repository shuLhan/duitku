// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// List of callback status code.
const (
	CallbackCodeSuccess = `00`
	CallbackCodeFail    = `01`
)

// ClearingCallbackResponse contains fields that must be set in order to
// response from ClearingTransfer callback.
//
// The signature created using the following formula:
//
//	SHA256(email + bankCode + bankAccount + accountName +
//		custRefNumber + amountTransfer + disburseId + secretKey)
type ClearingCallbackResponse struct {
	AccountName   string `json:"accountName"`
	BankAccount   string `json:"bankAccount"`
	BankCode      string `json:"bankCode"`
	CustRefNumber string `json:"custRefNumber"`
	Email         string `json:"email"`

	ErrorMsg   string `json:"errorMessage"`
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`

	Signature string `json:"signature"`

	Amount     int64 `json:"amountTransfer"`
	DisburseID int64 `json:"disburseId"`
	UserID     int64 `json:"userId"`
}

// NewClearingCallbackResponse create ClearingCallbackResponse from
// Clearing Transfer response.
//
// The StatusCode is set to success initially.
func NewClearingCallbackResponse(transferRes ClearingTransferResponse) (cbres *ClearingCallbackResponse) {
	cbres = &ClearingCallbackResponse{
		AccountName:   transferRes.AccountName,
		BankAccount:   transferRes.BankAccount,
		BankCode:      transferRes.BankCode,
		CustRefNumber: transferRes.CustRefNumber,

		StatusCode: CallbackCodeSuccess,

		Amount:     transferRes.Amount.Int64(),
		DisburseID: transferRes.DisburseID,
	}
	return cbres
}

// Sign set the Signature SHA256 and convert to hex.
func (cbres *ClearingCallbackResponse) Sign(opts ClientOptions) {
	var (
		bb        bytes.Buffer
		plainHash [sha256.Size]byte
	)

	bb.WriteString(cbres.Email)
	bb.WriteString(cbres.BankCode)
	bb.WriteString(cbres.BankAccount)
	bb.WriteString(cbres.AccountName)
	bb.WriteString(cbres.CustRefNumber)
	fmt.Fprintf(&bb, `%d%d`, cbres.Amount, cbres.DisburseID)
	bb.WriteString(opts.ApiKey)

	plainHash = sha256.Sum256(bb.Bytes())

	cbres.Signature = hex.EncodeToString(plainHash[:])
}
