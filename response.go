// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// List of known response code.
const (
	ResCodeSuccess             = `00`   // Approved or completed successfully.
	resCodeError               = `EE`   // General Error.
	resCodeErrTimeout          = `TO`   // Response time out from ATM Bersama Network (Do not retry).
	resCodeErrLink             = `LD`   // Link problem between Duitku and ATM Bersama Network.
	resCodeErrNF               = `NF`   // Transaction has not recorded on Remittance gateway.
	resCodeErrAccountInvalid   = `76`   // Invalid destination account.
	resCodeErrCallbackWait     = `80`   // Waiting for callback.
	resCodeErrOther            = `-100` // Other error (do not retry).
	resCodeErrUserID           = `-120` // User not found.
	resCodeErrUserBlocked      = `-123` // User has been blocked.
	resCodeErrAmount           = `-141` // Amount transfer invalid.
	resCodeErrTxFinish         = `-142` // Transaction already Finished.
	resCodeErrBankH2H          = `-148` // Bank not support H2H.
	resCodeErrBankNotFound     = `-149` // Bank not found.
	resCodeErrCallbackNotFound = `-161` // Callback URL not found.
	resCodeErrSignature        = `-191` // Wrong signature.
	resCodeErrAccountBlocked   = `-192` // Account number is blacklisted.
	resCodeErrEmail            = `-213` // Email is not valid.
	resCodeErrTransferNotFound = `-420` // Transfer not Found.
	resCodeErrFundInsufficient = `-510` // Insufficient Fund.
	resCodeErrFundLimit        = `-920` // Limit Exceeded.
	resCodeErrIP               = `-930` // IP not whitelisted.
	resCodeErrVendorTimeout    = `-951` // Time Out Vendor.
	resCodeErrParam            = `-952` // Invalid Parameter.
	resCodeErrTimestampExpired = `-960` // Timestamp is expired (5 minutes).
)

// Response contains commons fields for each HTTP response.
type Response struct {
	Code string `json:"responseCode"`
	Desc string `json:"responseDesc"`
}

// IsSuccess return true if the response code equal to 00.
func (res *Response) IsSuccess() bool {
	return res.Code == ResCodeSuccess
}
