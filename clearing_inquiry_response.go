// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// ClearingInquiryResponse contains response from calling [Clearing Inquiry
// request].
//
// [Clearing Inquiry request]: https://docs.duitku.com/disbursement/en/#clearing-inquiry-request
type ClearingInquiryResponse struct {
	Type string `json:"type"`

	RtolInquiryResponse
}
