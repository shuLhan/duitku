// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// request define common HTTP request fields.
type request struct {
	UserID    string `json:"userID"`
	Email     string `json:"email"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

func createRequest(opts ClientOptions) (req request) {
	req.UserID = opts.UserID
	req.Email = opts.Email
	req.Timestamp = time.Now().UnixMilli()

	var (
		plain   = fmt.Sprintf(`%s%d%s`, req.Email, req.Timestamp, opts.ApiKey)
		hashRaw = sha256.Sum256([]byte(plain))
	)

	req.Signature = hex.EncodeToString(hashRaw[:])

	return req
}
