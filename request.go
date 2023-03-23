// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Request define common HTTP request fields.
type Request struct {
	// Merchant email, filled from ClientOptions.Email.
	Email string `json:"email"`

	// Hash of some fields in the request along with its ApiKey.
	Signature string `json:"signature"`

	// Merchant ID, filled from ClientOptions.UserID.
	UserID int64 `json:"userId"`

	// Unix Timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
}

// CreateDisburseRequest create request for disbursement using ClientOptions
// opts.
func CreateDisburseRequest(opts ClientOptions) (req Request) {
	req.UserID = opts.DisburseUserID
	req.Email = opts.DisburseEmail
	req.Timestamp = time.Now().UnixMilli()

	var (
		plain   = fmt.Sprintf(`%s%d%s`, req.Email, req.Timestamp, opts.DisburseApiKey)
		hashRaw = sha256.Sum256([]byte(plain))
	)

	req.Signature = hex.EncodeToString(hashRaw[:])

	return req
}
