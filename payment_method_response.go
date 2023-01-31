// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// PaymentMethodResponse contains list of payments enabled by merchant.
type PaymentMethodResponse struct {
	Response

	PaymentFee []PaymentFee `json:"paymentFee"`
}
