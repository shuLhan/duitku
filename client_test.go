// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"encoding/json"
	"testing"

	"github.com/shuLhan/share/lib/test"
)

func TestClient_CheckBalance(t *testing.T) {
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

	balance, err = testClient.CheckBalance()
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

	test.Assert(t, `CheckBalance`, string(exp), string(got))
}

func TestClient_ClearingInquiry_sandbox(t *testing.T) {
	t.Skip(`This test require external call to server`)

	var (
		inquiryReq ClearingInquiry
		err        error
		tdata      *test.Data
		inquiryRes *ClearingInquiryResponse
	)

	tdata, err = test.LoadData(`testdata/disbursement/clearing_inquiry_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`request.json`], &inquiryReq)
	if err != nil {
		t.Fatal(err)
	}

	inquiryRes, err = testClient.ClearingInquiry(&inquiryReq)
	if err != nil {
		t.Fatal(err)
	}

	// We cannot compare the response, because for each call to server
	// it will return different DisburseID.

	t.Logf(`inquiryRes: %+v`, inquiryRes)

	test.Assert(t, `AccountName`, `Test Account`, inquiryRes.AccountName)
}

func TestClient_ClearingTransfer_sandbox(t *testing.T) {
	t.Skip(`This test require external call to server`)

	var (
		inquiryReq ClearingInquiry
		inquiryRes ClearingInquiryResponse

		err         error
		tdata       *test.Data
		transferRes *ClearingTransferResponse
	)

	tdata, err = test.LoadData(`testdata/disbursement/clearing_transfer_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`inquiry_request.json`], &inquiryReq)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(tdata.Input[`inquiry_response.json`], &inquiryRes)
	if err != nil {
		t.Fatal(err)
	}

	transferRes, err = testClient.ClearingTransfer(&inquiryReq, &inquiryRes)
	if err != nil {
		t.Fatal(err)
	}

	// We cannot compare the response, because for each call to server
	// it will return different DisburseID.

	t.Logf(`transferRes: %+v`, transferRes)

	test.Assert(t, `AccountName`, `Test Account`, transferRes.AccountName)
}

func TestClient_RtolInquiry_sandbox(t *testing.T) {
	t.Skip(`This test require external call to server`)

	var (
		inquiryReq RtolInquiry
		err        error
		tdata      *test.Data
		inquiryRes *RtolInquiryResponse
	)

	tdata, err = test.LoadData(`testdata/disbursement_rtol_inquiry_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`request.json`], &inquiryReq)
	if err != nil {
		t.Fatal(err)
	}

	inquiryRes, err = testClient.RtolInquiry(&inquiryReq)
	if err != nil {
		t.Fatal(err)
	}

	// We cannot compare the response, because for each call to server
	// it will return different CustRefNumber and DisburseID.

	t.Logf(`inquiryRes: %+v`, inquiryRes)

	test.Assert(t, `AccountName`, `Test Account`, inquiryRes.AccountName)
}

func TestClient_RtolTransfer_sandbox(t *testing.T) {
	t.Skip(`This test require external call to server`)

	var (
		err         error
		tdata       *test.Data
		inquiryReq  *RtolInquiry
		inquiryRes  *RtolInquiryResponse
		transferRes *RtolTransferResponse
	)

	tdata, err = test.LoadData(`testdata/disbursement_rtol_transfer_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`request_inquiry.json`], &inquiryReq)
	if err != nil {
		t.Fatal(err)
	}

	inquiryRes, err = testClient.RtolInquiry(inquiryReq)
	if err != nil {
		t.Fatal(err)
	}

	transferRes, err = testClient.RtolTransfer(inquiryReq, inquiryRes)
	if err != nil {
		t.Fatal(err)
	}

	// We cannot compare the response, because for each call to server
	// it will return different CustRefNumber and DisburseID.

	t.Logf(`RtolTransfer response: %+v`, transferRes)

	test.Assert(t, `AccountName`, `Test Account`, inquiryRes.AccountName)
}

func TestClient_ListBank(t *testing.T) {
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

	banks, err = testClient.ListBank()
	if err != nil {
		t.Fatal(err)
	}

	got, err = json.MarshalIndent(banks, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(`%s`, got)

	exp = tdata.Output[`response.json`]

	test.Assert(t, `ListBank`, string(exp), string(got))
}
