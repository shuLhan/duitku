// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"fmt"
	"net/url"
)

// ClientOptions configuration for HTTP client.
type ClientOptions struct {
	ServerUrl string

	// The hostname extracted from ServerUrl.
	host string

	// The merchant code is the project code obtained from the Duitku
	// merchant page.
	// This code is useful as an identifier of your project in each
	// request using the /merchant/* APIs.
	// You can get this code on every project you register on the
	// [merchant portal].
	//
	// [merchant portal]: https://passport.duitku.com/merchant/Project
	MerchantCode string

	// MerchantApiKey The API key for signing merchant related request.
	MerchantApiKey string

	// Merchant code and API key for payment through Indomaret.
	IndomaretMerchantCode string
	IndomaretApiKey       string

	// DisburseApiKey API key for signing disbursement request.
	DisburseApiKey string

	// DisburseEmail The email registered for disbursement in Duitku.
	DisburseEmail string

	// DisburseUserID user ID for disbursement request.
	DisburseUserID int64
}

// validate each field values.
func (opts *ClientOptions) validate() (err error) {
	var (
		urlServer *url.URL
	)

	urlServer, err = url.Parse(opts.ServerUrl)
	if err != nil {
		return fmt.Errorf(`invalid or empty ServerUrl: %s`, opts.ServerUrl)
	}
	opts.host = urlServer.Host

	if opts.DisburseUserID <= 0 {
		return fmt.Errorf(`invalid or empty DisburseUserID: %d`, opts.DisburseUserID)
	}
	if len(opts.DisburseEmail) == 0 {
		return fmt.Errorf(`invalid or empty DisburseEmail: %s`, opts.DisburseEmail)
	}
	if len(opts.DisburseApiKey) == 0 {
		return fmt.Errorf(`invalid or empty DisburseApiKey: %s`, opts.DisburseApiKey)
	}

	return nil
}
