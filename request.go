// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// request define common HTTP request fields.
type request struct {
	// Merchant email, filled from ClientOptions.Email.
	Email string `json:"email"`

	// Hash of some fields in the request along with its ApiKey.
	Signature string `json:"signature"`

	// Merchant ID, filled from ClientOptions.UserID.
	UserID int64 `json:"userId"`

	// Unix Timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
}

func createRequest(opts ClientOptions) (req request) {
	req.UserID, _ = strconv.ParseInt(opts.UserID, 10, 64)
	req.Email = opts.Email
	req.Timestamp = time.Now().UnixMilli()

	var (
		plain   = fmt.Sprintf(`%s%d%s`, req.Email, req.Timestamp, opts.ApiKey)
		hashRaw = sha256.Sum256([]byte(plain))
	)

	req.Signature = hex.EncodeToString(hashRaw[:])

	return req
}
