// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	// PaymentMethodDatetimeLayout define the date and time format for
	// PaymentMethod.DateTime.
	//
	// TODO: replace with time.DateTime once the go version is 1.20.
	PaymentMethodDatetimeLayout = `2006-01-02 15:04:05`
)

// PaymentMethod contains request for client MerchantPaymentMethod.
type PaymentMethod struct {
	// [REQ] Merchant code from Duitku.
	// Set by calling sign from ClientOptions.
	MerchantCode string `json:"merchantcode"`

	// [REQ] Format: yyyy-MM-dd HH:mm:ss
	// If empty, it will be set during sign using current date time.
	DateTime string `json:"datetime"`

	// [REQ] Formula: Sha256(merchantcode + paymentAmount + datetime +
	// apiKey).
	// The value of hash is stored as lowercase hexadecimal.
	Signature string `json:"signature"`

	// [REQ] Transaction amount. No decimal code (.) and no decimal digit.
	Amount int64 `json:"amount"`
}

// SetDateTime set the field DateTime using t.
func (req *PaymentMethod) SetDateTime(t time.Time) {
	req.DateTime = t.Format(PaymentMethodDatetimeLayout)
}

// Sign the request.
// Set the MerchantCode and DateTime if its empty, and then generate the
// Signature.
func (req *PaymentMethod) Sign(opts ClientOptions) {
	var (
		merchant = opts.DefaultMerchant

		plain string
		hash  [sha256.Size]byte
	)

	if len(req.MerchantCode) == 0 {
		req.MerchantCode = merchant.Code
	}

	if len(req.DateTime) == 0 {
		req.SetDateTime(time.Now())
	}

	plain = fmt.Sprintf(`%s%d%s%s`, req.MerchantCode, req.Amount, req.DateTime, merchant.ApiKey)
	hash = sha256.Sum256([]byte(plain))

	req.Signature = hex.EncodeToString(hash[:])
}
