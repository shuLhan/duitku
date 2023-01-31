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
	Email     string
	ApiKey    string

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

	// The hostname extracted from ServerUrl.
	host string

	UserID int64
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

	if opts.UserID <= 0 {
		return fmt.Errorf(`invalid or empty UserID: %d`, opts.UserID)
	}
	if len(opts.Email) == 0 {
		return fmt.Errorf(`invalid or empty Email: %s`, opts.Email)
	}
	if len(opts.ApiKey) == 0 {
		return fmt.Errorf(`invalid or empty ApiKey: %s`, opts.ApiKey)
	}

	return nil
}
