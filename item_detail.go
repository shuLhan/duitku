// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// [ItemDetail] define the subset of product being payed during payment.
//
// [ItemDetail]: https://docs.duitku.com/api/en/#item-details
type ItemDetail struct {
	// [REQ] Name of the item.
	Name string `json:"name"`

	// [REQ] Quantity of the item bought.
	Quantity int64 `json:"quantity"`

	// [REQ] Price of the Item. Note: Don't add decimal
	Price int64 `json:"price"`
}
