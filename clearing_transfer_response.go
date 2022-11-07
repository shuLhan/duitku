// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// ClearingTransferResponse contains response from Clearing Transfer.
type ClearingTransferResponse struct {
	Type string `json:"type"`

	RtolTransferResponse
}
