// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// List of valid Code in TxStatusResponse.
const (
	MerchantTxStatusSuccess = `00`
	MerchantTxStatusProcess = `01`
	MerchantTxStatusFailed  = `02`
)

// TxStatusResponse contains response from checking merchant payment
// status.
type TxStatusResponse struct {
	OrderID   string `json:"merchantOrderId"`
	Reference string `json:"reference"`
	Amount    string `json:"amount"`

	// Status code transaction.
	//   - 00 - Success
	//   - 01 - Process
	//   - 02 - Failed/Expired
	Code string `json:"statusCode"`

	// Description that explain the status Code.
	Message string `json:"statusMessage"`
}
