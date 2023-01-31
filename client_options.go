// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"fmt"
	"net/url"
	"os"

	"github.com/shuLhan/share/lib/ini"
)

// ClientOptions configuration for HTTP client.
type ClientOptions struct {
	ServerUrl string `ini:"duitku::server_url"`

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
	MerchantCode string `ini:"duitku::merchant_code"`

	// MerchantApiKey The API key for signing merchant related request.
	MerchantApiKey string `ini:"duitku::merchant_api_key"`

	// MerchantCallbackUrl The URL that will be used by Duitku to
	// confirm payments made by your customers.
	MerchantCallbackUrl string `ini:"duitku::merchant_callback_url"`

	// MerchantReturnUrl The URL that Duitku will direct the customer
	// after the transaction is successful or canceled.
	MerchantReturnUrl string `ini:"duitku::merchant_return_url"`

	// Merchant code and API key for payment through Indomaret.
	IndomaretMerchantCode string `ini:"duitku::indomaret_merchant_code"`
	IndomaretApiKey       string `ini:"duitku::indomaret_api_key"`

	// DisburseApiKey API key for signing disbursement request.
	DisburseApiKey string `ini:"duitku::disburse_api_key"`

	// DisburseEmail The email registered for disbursement in Duitku.
	DisburseEmail string `ini:"duitku::disburse_email"`

	// DisburseUserID user ID for disbursement request.
	DisburseUserID int64 `ini:"duitku::disburse_user_id"`
}

// LoadClientOptions load ClientOptions from file.
// See the client.conf.example for an example for client configuration.
func LoadClientOptions(file string) (opts *ClientOptions, err error) {
	var (
		logp    = `LoadClientOptions`
		content []byte
	)
	content, err = os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	opts = &ClientOptions{}
	err = ini.Unmarshal(content, opts)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return opts, nil
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
