// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"log"
	"os"
	"testing"
)

var (
	testClient *Client
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
