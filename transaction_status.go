// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type transactionStatus struct {
	MerchantCode string `form:"merchantCode"`
	OrderID      string `form:"merchantOrderId"`

	// Transaction identification code.
	// Formula: md5(merchantCode + merchantOrderId + apiKey).
	Signature string `form:"signature"`
}

func (tx *transactionStatus) sign(opts ClientOptions, paymentMethod string) {
	var merchant = opts.Merchant(paymentMethod)

	tx.MerchantCode = merchant.Code

	var (
		plain    = fmt.Sprintf(`%s%s%s`, tx.MerchantCode, tx.OrderID, merchant.ApiKey)
		md5Plain = md5.Sum([]byte(plain))
	)

	tx.Signature = hex.EncodeToString(md5Plain[:])
}
