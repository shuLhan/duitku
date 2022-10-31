// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"github.com/shuLhan/share/lib/math/big"
)

type Bank struct {
	MaxTransfer *big.Rat `json:"maxAmountTransfer"`

	Code string `json:"bankCode"`
	Name string `json:"bankName"`
}
