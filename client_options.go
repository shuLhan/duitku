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
