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
	UserID    string
	Email     string
	ApiKey    string

	// The hostname extracted from ServerUrl.
	host string
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

	if len(opts.UserID) == 0 {
		return fmt.Errorf(`invalid or empty UserID: %s`, opts.UserID)
	}
	if len(opts.Email) == 0 {
		return fmt.Errorf(`invalid or empty Email: %s`, opts.Email)
	}
	if len(opts.ApiKey) == 0 {
		return fmt.Errorf(`invalid or empty ApiKey: %s`, opts.ApiKey)
	}
	return nil
}
