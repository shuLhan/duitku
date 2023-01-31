// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// InquiryStatus request for transfer status.
type InquiryStatus struct {
	Request

	DisburseID int64 `json:"disburseId"`
}

// Sign the request, fill the UserID, Email, Timestamp, and generate the
// Signature.
func (inq *InquiryStatus) Sign(opts ClientOptions) {
	inq.UserID = opts.DisburseUserID
	inq.Email = opts.DisburseEmail
	inq.Timestamp = time.Now().UnixMilli()

	var plain string = fmt.Sprintf(`%s%d%d%s`, inq.Email, inq.Timestamp, inq.DisburseID, opts.DisburseApiKey)
	var plainHash [sha256.Size]byte = sha256.Sum256([]byte(plain))

	inq.Signature = hex.EncodeToString(plainHash[:])
}
