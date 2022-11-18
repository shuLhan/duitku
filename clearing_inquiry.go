// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// ClearingInquiry contains request to initiate transfer from merchant to
// customer's bank account using [Clearing type].
//
// For Signature it use the following formula:
//
//	SHA256(email + timestamp + bankCode + type + bankAccount +
//		amountTransfer + purpose + apiKey)
//
// [Clearing type]: https://docs.duitku.com/disbursement/en/#clearing-inquiry-request
type ClearingInquiry struct {
	// 9 digits customer reference number.
	CustRefNumber string `json:"custRefNumber"`

	// Type of clearing: LLG, RTGS, H2H, or BIFAST.
	Type string `json:"type"`

	RtolInquiry
}

// Sign the request, fill the UserID, Email, Timestamp, and generate the
// Signature.
func (inq *ClearingInquiry) Sign(opts ClientOptions) {
	inq.UserID = opts.UserID
	inq.Email = opts.Email
	inq.Timestamp = time.Now().UnixMilli()

	var (
		plain = fmt.Sprintf(`%s%d%s%s%s%d%s%s`, inq.Email,
			inq.Timestamp, inq.BankCode, inq.Type,
			inq.BankAccount, inq.Amount, inq.Purpose, opts.ApiKey)
		plainHash [sha256.Size]byte = sha256.Sum256([]byte(plain))
	)
	inq.Signature = hex.EncodeToString(plainHash[:])
}
