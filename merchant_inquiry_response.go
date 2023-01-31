// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// MerchantInquiryResponse contains response from MerchantInquiry.
type MerchantInquiryResponse struct {
	Response

	// Indicates which project used in this transaction.
	MerchantCode string `json:"merchantCode"`

	// Reference number from Duitku (need to be save on your system).
	Reference string `json:"reference"`

	// Payment link for direction to Duitku payment page.
	PaymentUrl string `json:"paymentUrl"`

	// Payment number or virtual account.
	VANumber string `json:"vaNumber"`

	// QR string is used if you use QRIS payment (you need to generate QR
	// code from this string).
	QRString string `json:"qrString"`

	// Payment amount.
	Amount string `json:"amount"`
}
