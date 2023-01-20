// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// RtolTransferResponse contains response from online transfer.
//
// NOTE: the actual response from server does not return DisburseID.
type RtolTransferResponse struct {
	Purpose string `json:"purpose"`

	RtolInquiryResponse

	UserID int64 `json:"userId"`
}
