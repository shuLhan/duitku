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
		clOpts = ClientOptions{
			ServerUrl:      ServerUrlSandbox,
			DisburseUserID: 3551,
			DisburseEmail:  `test@chakratechnology.com`,
			DisburseApiKey: `de56f832487bc1ce1de5ff2cfacf8d9486c61da69df6fd61d5537b6b7d6d354d`,
		}

		err error
	)

	testClient, err = NewClient(clOpts)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
