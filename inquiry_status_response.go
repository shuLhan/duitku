// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import "github.com/shuLhan/share/lib/math/big"

type InquiryStatusResponse struct {
	Response

	// Destination Bank Code.
	BankCode string `json:"bankCode"`

	// Destination account number.
	BankAccount string `json:"bankAccount"`

	// Disbursement transfer amount.
	Amount *big.Rat `json:"amountTransfer"`

	// Bank Account owner.
	AccountName string `json:"accountName"`

	// 9 Digit Customer reference number that will be printed when the
	// transaction is successful.
	CustRefNumber string `json:"custRefNumber"`
}
