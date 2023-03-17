// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type PaymentStatus struct {
	MerchantCode string `form:"merchantCode"`
	OrderID      string `form:"merchantOrderId"`

	// Transaction identification code.
	// Formula: md5(merchantCode + merchantOrderId + apiKey).
	Signature string `form:"signature"`
}

func (req *PaymentStatus) sign(opts ClientOptions) {
	var merchant = opts.Merchant(req.MerchantCode)

	req.MerchantCode = merchant.Code

	var (
		plain    = fmt.Sprintf(`%s%s%s`, req.MerchantCode, req.OrderID, merchant.ApiKey)
		md5Plain = md5.Sum([]byte(plain))
	)

	req.Signature = hex.EncodeToString(md5Plain[:])
}
