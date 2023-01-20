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
		tdata      *test.Data
		gotBalance *Balance
		expBalance *Balance
		err        error
		exp        []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement/checkbalance_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	gotBalance, err = testClient.CheckBalance()
	if err != nil {
		t.Fatal(err)
	}

	exp = tdata.Output[`response.json`]
	err = json.Unmarshal(exp, &expBalance)
	if err != nil {
		t.Fatal(err)
	}

	expBalance.Current = gotBalance.Current
	expBalance.Effective = gotBalance.Effective

	test.Assert(t, `CheckBalance`, gotBalance, expBalance)
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

		transferReq *ClearingTransfer
		transferRes *ClearingTransferResponse

		tdata *test.Data
		err   error
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

	transferReq = NewClearingTransfer(&inquiryReq, &inquiryRes)

	transferRes, err = testClient.ClearingTransfer(transferReq)
	if err != nil {
		t.Fatal(err)
	}

	// We cannot compare the response, because for each call to server
	// it will return different DisburseID.

	t.Logf(`transferRes: %+v`, transferRes)

	test.Assert(t, `AccountName`, `Test Account`, transferRes.AccountName)
}

func TestClient_InquiryStatus_sandbox(t *testing.T) {
	var (
		tdata *test.Data

		reqInquiry *RtolInquiry
		resInquiry *RtolInquiryResponse
		expInquiry *RtolInquiryResponse

		reqTransfer *RtolTransfer
		resTransfer *RtolTransferResponse
		expTransfer *RtolTransferResponse

		resInqueryStatus *InquiryStatusResponse
		expInquiryStatus *InquiryStatusResponse

		err error
		exp []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement/inquirystatus_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`inquiry_request.json`], &reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	// Do inquiry bank account ...

	resInquiry, err = testClient.RtolInquiry(reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	exp = tdata.Output[`inquiry_response.json`]
	err = json.Unmarshal(exp, &expInquiry)
	if err != nil {
		t.Fatal(err)
	}
	expInquiry.CustRefNumber = resInquiry.CustRefNumber
	expInquiry.DisburseID = resInquiry.DisburseID
	test.Assert(t, `RtolInquiry`, expInquiry, resInquiry)

	// Do the transfer ...

	reqTransfer = NewRtolTransfer(reqInquiry, resInquiry)

	resTransfer, err = testClient.RtolTransfer(reqTransfer)
	if err != nil {
		t.Fatal(err)
	}

	exp = tdata.Output[`transfer_response.json`]
	err = json.Unmarshal(exp, &expTransfer)
	if err != nil {
		t.Fatal(err)
	}
	expTransfer.CustRefNumber = resTransfer.CustRefNumber
	expTransfer.DisburseID = resTransfer.DisburseID
	test.Assert(t, `RtolTransfer`, expTransfer, resTransfer)

	// Inquiry transfer status ...

	resInqueryStatus, err = testClient.InquiryStatus(resTransfer.DisburseID)
	if err != nil {
		t.Fatal(err)
	}

	exp = tdata.Output[`inquirystatus_response.json`]
	err = json.Unmarshal(exp, &expInquiryStatus)
	if err != nil {
		t.Fatal(err)
	}
	expInquiryStatus.CustRefNumber = resInqueryStatus.CustRefNumber
	test.Assert(t, `InquiryStatus`, expInquiryStatus, resInqueryStatus)
}

func TestClient_RtolInquiry_sandbox(t *testing.T) {
	var (
		err        error
		tdata      *test.Data
		reqInquiry *RtolInquiry
		gotInquiry *RtolInquiryResponse
		expInquiry *RtolInquiryResponse
		rawb       []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement/rtol_inquiry_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	rawb = tdata.Input[`request.json`]
	err = json.Unmarshal(rawb, &reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	gotInquiry, err = testClient.RtolInquiry(reqInquiry)
	if err != nil {
		t.Fatal(err)
	}
	rawb = tdata.Output[`response.json`]
	err = json.Unmarshal(rawb, &expInquiry)
	if err != nil {
		t.Fatal(err)
	}
	// Set dynamic field values in expected response.
	expInquiry.CustRefNumber = gotInquiry.CustRefNumber
	expInquiry.DisburseID = gotInquiry.DisburseID
	test.Assert(t, `RtolInquiry`, expInquiry, gotInquiry)
}

func TestClient_RtolTransfer_sandbox(t *testing.T) {
	var (
		err   error
		tdata *test.Data

		reqInquiry *RtolInquiry
		gotInquiry *RtolInquiryResponse
		expInquiry *RtolInquiryResponse

		reqTransfer *RtolTransfer
		gotTransfer *RtolTransferResponse
		expTransfer *RtolTransferResponse

		rawb []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement/rtol_transfer_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	// Do inquiry ...

	rawb = tdata.Input[`inquiry_request.json`]
	err = json.Unmarshal(rawb, &reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	gotInquiry, err = testClient.RtolInquiry(reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	rawb = tdata.Output[`inquiry_response.json`]
	err = json.Unmarshal(rawb, &expInquiry)
	if err != nil {
		t.Fatal(err)
	}
	expInquiry.CustRefNumber = gotInquiry.CustRefNumber
	expInquiry.DisburseID = gotInquiry.DisburseID
	test.Assert(t, `RtolInquiry`, expInquiry, gotInquiry)

	// Do transfer ...

	reqTransfer = NewRtolTransfer(reqInquiry, gotInquiry)

	gotTransfer, err = testClient.RtolTransfer(reqTransfer)
	if err != nil {
		t.Fatal(err)
	}

	rawb = tdata.Output[`transfer_response.json`]
	err = json.Unmarshal(rawb, &expTransfer)
	if err != nil {
		t.Fatal(err)
	}
	expTransfer.CustRefNumber = gotTransfer.CustRefNumber
	expTransfer.DisburseID = gotTransfer.DisburseID
	test.Assert(t, `RtolTransfer`, expTransfer, gotTransfer)
}

func TestClient_ListBank(t *testing.T) {
	var (
		tdata *test.Data
		err   error
		banks []Bank
		exp   []byte
		got   []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement/listbank_test.txt`)
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

	exp = tdata.Output[`response.json`]
	test.Assert(t, `ListBank`, string(exp), string(got))
}
