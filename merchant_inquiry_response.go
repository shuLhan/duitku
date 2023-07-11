// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// MerchantInquiryResponse contains response from MerchantInquiry.
type MerchantInquiryResponse struct {
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

	// Status code transaction.
	//   - 00 - Success
	//   - 01 - Process
	//   - 02 - Failed/Expired
	Code string `json:"statusCode"`

	// Description that explain the status Code.
	Message string `json:"statusMessage"`

	// Response body when http request failed
	ErrorMessage string `json:"Message,omitempty"`
}
