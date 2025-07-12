// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"git.sr.ht/~shulhan/pakakeh.go/lib/math/big"
)

// Bank contains bank information from response of ListBank.
type Bank struct {
	MaxTransfer *big.Rat `json:"maxAmountTransfer"`

	Code string `json:"bankCode"`
	Name string `json:"bankName"`
}
