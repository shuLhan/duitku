// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	testClient         *Client
	testClientMerchant *Client
)

func TestMain(m *testing.M) {
	var (
		opts *ClientOptions
		err  error
	)

	opts, err = LoadClientOptions(`client.conf.example`)
	if err != nil {
		log.Fatal(err)
	}

	testClient, err = NewClient(*opts)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func initClientMerchant() (err error) {
	if testClientMerchant != nil {
		return nil
	}

	var (
		logp = `initClientMerchant`

		opts *ClientOptions
	)

	opts, err = LoadClientOptions(`client.conf`)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	testClientMerchant, err = NewClient(*opts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	return nil
}
