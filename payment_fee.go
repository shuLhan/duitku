// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// PaymentFee contains [fee] for payment method.
//
// If the settings in the merchant portal fees are charged to the
// merchant, the TotalFee will appear 0.
// The TotalFee will appear if it is charged to the customer.
//
// [fee]: https://docs.duitku.com/api/en/#payment-fee
type PaymentFee struct {
	// Payment method code.
	PaymentMethod string `json:"paymentMethod"`

	// Payment method name.
	PaymentName string `json:"paymentName"`

	// Payment method image url.
	PaymentImage string `json:"paymentImage"`

	// Payment Fee.
	TotalFee string `json:"totalFee"`
}
