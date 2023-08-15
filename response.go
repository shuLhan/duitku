// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// List of known [status codes].
//
// [status codes]: https://docs.duitku.com/disbursement/en/#status-code
const (
	StatusCodeSuccess             = `00`   // Approved or completed successfully.
	StatusCodeError               = `EE`   // General Error.
	StatusCodeErrTimeout          = `TO`   // Response time out from ATM Bersama Network (Do not retry).
	StatusCodeErrLink             = `LD`   // Link problem between Duitku and ATM Bersama Network.
	StatusCodeErrNF               = `NF`   // Transaction has not recorded on Remittance gateway.
	StatusCodeErrAccountInvalid   = `76`   // Invalid destination account.
	StatusCodeErrCallbackWait     = `80`   // Waiting for callback.
	StatusCodeErrOther            = `-100` // Other error (do not retry).
	StatusCodeErrUserID           = `-120` // User not found.
	StatusCodeErrUserBlocked      = `-123` // User has been blocked.
	StatusCodeErrAmount           = `-141` // Amount transfer invalid.
	StatusCodeErrTxFinish         = `-142` // Transaction already Finished.
	StatusCodeErrBankH2H          = `-148` // Bank not support H2H.
	StatusCodeErrBankNotFound     = `-149` // Bank not found.
	StatusCodeErrCallbackNotFound = `-161` // Callback URL not found.
	StatusCodeErrSignature        = `-191` // Wrong signature.
	StatusCodeErrAccountBlocked   = `-192` // Account number is blacklisted.
	StatusCodeErrEmail            = `-213` // Email is not valid.
	StatusCodeErrTransferNotFound = `-420` // Transfer not Found.
	StatusCodeErrFundInsufficient = `-510` // Insufficient Fund.
	StatusCodeErrFundLimit        = `-920` // Limit Exceeded.
	StatusCodeErrIP               = `-930` // IP not whitelisted.
	StatusCodeErrVendorTimeout    = `-951` // Time Out Vendor.
	StatusCodeErrParam            = `-952` // Invalid Parameter.
	StatusCodeErrTimestampExpired = `-960` // Timestamp is expired (5 minutes).
)

// Response contains commons fields for each HTTP response.
type Response struct {
	Code string `json:"responseCode"`
	Desc string `json:"responseDesc"`
}

// IsSuccess return true if the response code equal to 00.
func (res *Response) IsSuccess() bool {
	return res.Code == StatusCodeSuccess
}
