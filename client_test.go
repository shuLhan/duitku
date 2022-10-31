// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"encoding/json"
	"testing"

	"github.com/shuLhan/share/lib/test"
)

func TestClient_DisbursementCheckBalance(t *testing.T) {
	var (
		tdata   *test.Data
		balance *Balance
		err     error
		exp     []byte
		got     []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement_checkbalance_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	balance, err = testClient.DisbursementCheckBalance()
	if err != nil {
		t.Fatal(err)
	}

	// Set the value to zero, since their value may be different on each
	// run.
	balance.Current.Scan(0)
	balance.Effective.Scan(0)

	got, err = json.MarshalIndent(balance, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	exp = tdata.Output[`response.json`]

	test.Assert(t, `DisbursementCheckBalance`, string(exp), string(got))
}

func TestClient_DisbursementListBank(t *testing.T) {
	var (
		tdata *test.Data
		err   error
		banks []Bank
		exp   []byte
		got   []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement_listbank_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	banks, err = testClient.DisbursementListBank()
	if err != nil {
		t.Fatal(err)
	}

	got, err = json.MarshalIndent(banks, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(`%s`, got)

	exp = tdata.Output[`response.json`]

	test.Assert(t, `DisbursementListBank`, string(exp), string(got))
}
