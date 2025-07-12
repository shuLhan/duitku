// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"bytes"
	"encoding/json"
	"testing"

	"git.sr.ht/~shulhan/pakakeh.go/lib/test"
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
	var (
		tdata      *test.Data
		reqInquiry *ClearingInquiry
		gotInquiry *ClearingInquiryResponse
		expInquiry *ClearingInquiryResponse
		err        error
		rawb       []byte
	)

	tdata, err = test.LoadData(`testdata/disbursement/clearing_inquiry_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`request.json`], &reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	gotInquiry, err = testClient.ClearingInquiry(reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	rawb = tdata.Output[`response.json`]
	err = json.Unmarshal(rawb, &expInquiry)
	if err != nil {
		t.Fatal(err)
	}
	expInquiry.CustRefNumber = gotInquiry.CustRefNumber
	expInquiry.DisburseID = gotInquiry.DisburseID
	test.Assert(t, `ClearingInquiry`, expInquiry, gotInquiry)
}

func TestClient_ClearingTransfer_sandbox(t *testing.T) {
	var (
		reqInquiry *ClearingInquiry
		resInquiry *ClearingInquiryResponse

		reqTransfer *ClearingTransfer
		gotTransfer *ClearingTransferResponse
		expTransfer *ClearingTransferResponse

		tdata *test.Data
		err   error
	)

	tdata, err = test.LoadData(`testdata/disbursement/clearing_transfer_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Input[`inquiry_request.json`], &reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	resInquiry, err = testClient.ClearingInquiry(reqInquiry)
	if err != nil {
		t.Fatal(err)
	}

	reqTransfer = NewClearingTransfer(reqInquiry, resInquiry)

	gotTransfer, err = testClient.ClearingTransfer(reqTransfer)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(tdata.Output[`transfer_response.json`], &expTransfer)
	if err != nil {
		t.Fatal(err)
	}
	expTransfer.CustRefNumber = gotTransfer.CustRefNumber
	expTransfer.DisburseID = gotTransfer.DisburseID
	test.Assert(t, `ClearingTransfer`, expTransfer, gotTransfer)
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
	expInquiryStatus.Response.Desc = `SUCCESS`
	test.Assert(t, `InquiryStatus`, expInquiryStatus, resInqueryStatus)
}

func TestClient_MerchantInquiry(t *testing.T) {
	var (
		tdata *test.Data
		err   error
	)

	err = initClientMerchant()
	if err != nil {
		t.Skip(err)
	}

	tdata, err = test.LoadData(`testdata/merchant/inquiry_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	var (
		req  *MerchantInquiry
		resp *MerchantInquiryResponse
		tag  string
	)

	tag = `request.json`
	err = json.Unmarshal(tdata.Input[tag], &req)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = testClientMerchant.MerchantInquiry(req)
	if err != nil {
		t.Fatal(err)
	}

	var (
		exp []byte
		got []byte
	)

	resp.MerchantCode = `[redacted]`

	got, err = json.MarshalIndent(resp, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	tag = `response.json`
	exp = tdata.Output[tag]
	exp = bytes.ReplaceAll(exp, []byte(`$ref`), []byte(resp.Reference))
	exp = bytes.ReplaceAll(exp, []byte(`$payment_url`), []byte(resp.PaymentUrl))
	exp = bytes.ReplaceAll(exp, []byte(`$va`), []byte(resp.VANumber))

	t.Logf(`MerchantInquiry: response: %s`, got)

	test.Assert(t, `MerchantInquiry`, string(exp), string(got))

	// Test checking the transaction status.

	var (
		paymentReq = &PaymentStatus{
			MerchantCode: req.PaymentMethod,
			OrderID:      req.MerchantOrderId,
		}
		paymentResp *PaymentStatusResponse
	)

	paymentResp, err = testClientMerchant.MerchantPaymentStatus(paymentReq)
	if err != nil {
		t.Fatal(err)
	}

	got, err = json.MarshalIndent(paymentResp, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	tag = `tx_status_response.json`
	exp = tdata.Output[tag]
	exp = bytes.ReplaceAll(exp, []byte(`$ref`), []byte(resp.Reference))
	t.Logf(`MerchantPaymentStatus: response: %s`, got)
	test.Assert(t, `MerchantPaymentStatus`, string(exp), string(got))
}

func TestClient_MerchantPaymentMethod(t *testing.T) {
	var (
		tdata *test.Data
		err   error
	)

	err = initClientMerchant()
	if err != nil {
		t.Skip(err)
	}

	tdata, err = test.LoadData(`testdata/merchant/payment_method_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	var (
		req  *PaymentMethod
		resp *PaymentMethodResponse
		tag  string
	)

	tag = `payment_method_request.json`
	err = json.Unmarshal(tdata.Input[tag], &req)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = testClientMerchant.MerchantPaymentMethod(req)
	if err != nil {
		t.Fatal(err)
	}

	var (
		exp []byte
		got []byte
	)

	got, err = json.MarshalIndent(resp, ``, `  `)
	if err != nil {
		t.Fatal(err)
	}

	tag = `payment_method_response.json`
	exp = tdata.Output[tag]
	test.Assert(t, `PaymentMethod`, string(exp), string(got))
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
