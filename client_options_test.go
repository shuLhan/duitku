// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"encoding/json"
	"testing"

	"git.sr.ht/~shulhan/pakakeh.go/lib/ini"
	"git.sr.ht/~shulhan/pakakeh.go/lib/test"
)

func TestLoadClientOptions(t *testing.T) {
	var (
		tdata *test.Data
		err   error
	)

	tdata, err = test.LoadData(`testdata/client_options_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	var (
		opts ClientOptions
		tag  string
		data []byte
	)

	data = tdata.Input[`all.conf`]

	err = ini.Unmarshal(data, &opts)
	if err != nil {
		t.Fatal(err)
	}

	err = opts.initAndValidate()
	if err != nil {
		t.Fatal(err)
	}

	var (
		exp []byte
		got []byte
	)

	got, err = json.MarshalIndent(&opts, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	tag = `all.json`
	exp = tdata.Output[tag]
	test.Assert(t, tag, string(exp), string(got))
}
