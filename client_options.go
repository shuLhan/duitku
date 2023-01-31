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

	// DefaultMerchant default merchant to be used for payment.
	DefaultMerchant Merchant `ini:"default-merchant"`

	// PaymentMerchant specific merchant to be used based on payment
	// method.
	PaymentMerchant map[string]Merchant `ini:"merchant"`

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

// Merchant return the PaymentMerchant based on paymentMethod.
// If no key found, it will return DefaultMerchant.
func (opts *ClientOptions) Merchant(paymentMethod string) (merchant Merchant) {
	var (
		found bool
	)

	merchant, found = opts.PaymentMerchant[paymentMethod]
	if !found {
		merchant = opts.DefaultMerchant
	}

	return merchant
}

// initAndValidate each field values.
func (opts *ClientOptions) initAndValidate() (err error) {
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

	var (
		method   string
		merchant Merchant
	)
	for method, merchant = range opts.PaymentMerchant {
		if len(merchant.CallbackUrl) == 0 {
			merchant.CallbackUrl = opts.DefaultMerchant.CallbackUrl
		}
		if len(merchant.ReturnUrl) == 0 {
			merchant.ReturnUrl = opts.DefaultMerchant.ReturnUrl
		}
		opts.PaymentMerchant[method] = merchant
	}

	return nil
}
