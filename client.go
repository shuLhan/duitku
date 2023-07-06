// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	libhttp "github.com/shuLhan/share/lib/http"
)

// List of known and implemented HTTP API paths.
const (
	PathListBank     = `/webapi/api/disbursement/listBank`
	PathCheckBalance = `/webapi/api/disbursement/checkBalance`

	// Paths for transfer online.
	PathInquiry        = `/webapi/api/disbursement/inquiry`
	PathInquirySandbox = `/webapi/api/disbursement/inquirysandbox` // Used for testing.

	// Endpoints to check transfer status.
	PathInquiryStatus        = `/webapi/api/disbursement/inquirystatus`
	PathInquiryStatusSandbox = `/webapi/api/disbursement/inquirystatus` // Used for testing.

	PathTransfer        = `/webapi/api/disbursement/transfer`
	PathTransferSandbox = `/webapi/api/disbursement/transfersandbox` // Used for testing.

	// Paths for Clearing.
	PathInquiryClearing        = `/webapi/api/disbursement/inquiryclearing`
	PathInquiryClearingSandbox = `/webapi/api/disbursement/inquiryclearingsandbox` // Used for testing.

	PathTransferClearing        = `/webapi/api/disbursement/transferclearing`
	PathTransferClearingSandbox = `/webapi/api/disbursement/transferclearingsandbox` // Used for testing.

	PathMerchantInquiry           = `/webapi/api/merchant/v2/inquiry`
	PathMerchantPaymentMethod     = `/webapi/api/merchant/paymentmethod/getpaymentmethod`
	PathMerchantTransactionStatus = `/webapi/api/merchant/transactionStatus`
)

// Client HTTP client for duitku disbursement and payment APIs.
type Client struct {
	*libhttp.Client

	opts ClientOptions
}

// NewClient create and initialize new Client.
func NewClient(opts ClientOptions) (cl *Client, err error) {
	var (
		logp      = `NewClient`
		httpcOpts = libhttp.ClientOptions{
			ServerUrl: opts.ServerUrl,
		}
	)

	err = opts.initAndValidate()
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	cl = &Client{
		Client: libhttp.NewClient(&httpcOpts),
		opts:   opts,
	}

	return cl, nil
}

// CheckBalance get the current balances.
func (cl *Client) CheckBalance() (bal *Balance, err error) {
	var (
		logp = `CheckBalance`
		req  = CreateDisburseRequest(cl.opts)

		httpRes *http.Response
		resBody []byte
	)

	httpRes, resBody, err = cl.PostJSON(PathCheckBalance, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, httpRes.Status)
	}

	err = json.Unmarshal(resBody, &bal)
	if err != nil {
		return nil, fmt.Errorf(`%s: %s`, logp, err)
	}

	return bal, nil
}

// ClearingInquiry initiate the transfer for Clearing using LLG, RTGS, H2H, or
// BI-FAST.
func (cl *Client) ClearingInquiry(req *ClearingInquiry) (res *ClearingInquiryResponse, err error) {
	var (
		logp = `ClearingInquiry`
		path = PathInquiryClearing

		httpRes *http.Response
		resBody []byte
	)

	req.Sign(cl.opts)

	// Since the path is different in test environment, we check the host
	// here to set it.
	if cl.opts.host != hostLive {
		path = PathInquiryClearingSandbox
	}

	httpRes, resBody, err = cl.PostJSON(path, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, httpRes.Status)
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return res, nil
}

// ClearingTransfer do the clearing transfer to the bank account.
//
// Return without an error does not mean the transfer success, you need to
// check the response Code.
func (cl *Client) ClearingTransfer(req *ClearingTransfer) (res *ClearingTransferResponse, err error) {
	var (
		logp = `ClearingTransfer`
		path = PathTransferClearing

		httpRes *http.Response
		resBody []byte
	)

	req.Sign(cl.opts)

	// Since the path is different in test environment, we check the host
	// here to set it.
	if cl.opts.host != hostLive {
		path = PathTransferClearingSandbox
	}

	httpRes, resBody, err = cl.PostJSON(path, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, httpRes.Status)
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return res, nil
}

// InquiryStatus get the transfer status of ClearingTransfer or RtolTransfer.
func (cl *Client) InquiryStatus(disburseID int64) (res *InquiryStatusResponse, err error) {
	var (
		logp = `InquiryStatus`
		path = PathInquiryStatus
		req  = InquiryStatus{
			DisburseID: disburseID,
		}

		resHttp *http.Response
		resBody []byte
	)

	// Since the path is different in test environment, we check the host
	// here to set it.
	if cl.opts.host != hostLive {
		path = PathInquiryStatusSandbox
	}

	req.Sign(cl.opts)

	resHttp, resBody, err = cl.PostJSON(path, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if resHttp.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, resHttp.Status)
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return res, nil
}

// ListBank fetch list of banks for disbursement.
// The returned list bank is sorted by code and name in ascending order.
func (cl *Client) ListBank() (banks []Bank, err error) {
	var (
		logp = `ListBank`
		req  = CreateDisburseRequest(cl.opts)
		res  = struct {
			Data interface{} `json:"Banks"`
			Code string      `json:"responseCode"`
			Desc string      `json:"responseDesc"`
		}{}

		httpRes *http.Response
		resBody []byte
	)

	httpRes, resBody, err = cl.PostJSON(PathListBank, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, httpRes.Status)
	}

	res.Data = &banks

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, fmt.Errorf(`%s: %s`, logp, err)
	}
	if res.Code != ResCodeSuccess {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, res.Code, res.Desc)
	}

	// Sort the list by Code and then by Name in case code is equal.
	sort.Slice(banks, func(x, y int) bool {
		var cmp = strings.Compare(banks[x].Code, banks[y].Code)
		if cmp == -1 {
			return true
		}
		if cmp == 1 {
			return false
		}
		cmp = strings.Compare(banks[x].Name, banks[y].Name)
		return cmp == -1
	})

	return banks, nil
}

// MerchantInquiry request payment to the Duitku system (via virtual account
// numbers, QRIS, e-wallet, and so on).
//
// Ref: https://docs.duitku.com/api/en/#request-transaction
func (cl *Client) MerchantInquiry(req *MerchantInquiry) (resp *MerchantInquiryResponse, err error) {
	var (
		logp = `MerchantInquiry`

		httpRes *http.Response
		resBody []byte
	)

	req.sign(cl.opts)

	httpRes, resBody, err = cl.PostJSON(PathMerchantInquiry, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 400 {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, httpRes.Status, resBody)
	}

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return resp, nil
}

// MerchantPaymentMethod get active payment methods from the merchant (your)
// project.
//
// Ref: https://docs.duitku.com/api/en/#get-payment-method
func (cl *Client) MerchantPaymentMethod(req *PaymentMethod) (resp *PaymentMethodResponse, err error) {
	var (
		logp = `MerchantPaymentMethod`
		path = PathMerchantPaymentMethod

		httpRes *http.Response
		resBody []byte
	)

	req.Sign(cl.opts)

	httpRes, resBody, err = cl.PostJSON(path, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 400 {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, httpRes.Status, resBody)
	}

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return resp, nil
}

// MerchantPaymentStatus get the [status of payment] from customer.
//
// [status of payment]: https://docs.duitku.com/api/en/#check-transaction
func (cl *Client) MerchantPaymentStatus(req *PaymentStatus) (resp *PaymentStatusResponse, err error) {
	var (
		logp = `MerchantPaymentStatus`

		params  url.Values
		httpRes *http.Response
		resBody []byte
	)

	req.sign(cl.opts)

	params, err = libhttp.MarshalForm(*req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	httpRes, resBody, err = cl.PostForm(PathMerchantTransactionStatus, nil, params)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if httpRes.StatusCode >= 400 {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, httpRes.Status, resBody)
	}

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return resp, nil
}

// Options return the current client configuration.
func (cl *Client) Options() (opts ClientOptions) {
	return cl.opts
}

// RtolInquiry get the information of the name of the account owner of the
// transfer destination.
//
// After getting this information, customers can determine whether the purpose
// of such a transfer is in accordance with the intended or not.
// If appropriate, the customer can proceed to the transfer process.
//
// Ref: https://docs.duitku.com/disbursement/en/#transfer-online
func (cl *Client) RtolInquiry(req *RtolInquiry) (res *RtolInquiryResponse, err error) {
	var (
		logp = `RtolInquiry`
		path = PathInquiry

		resHttp *http.Response
		resBody []byte
	)

	// Since the path is different in test environment, we check the host
	// here to set it.
	if cl.opts.host != hostLive {
		path = PathInquirySandbox
	}

	req.Sign(cl.opts)

	resHttp, resBody, err = cl.PostJSON(path, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if resHttp.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, resHttp.Status)
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return res, nil
}

// RtolTransfer do the actual transfer to customer's bank account.
//
// Transfer will be limited from 25 to 50 Million per transaction depending on
// the beneficiary bank account.
//
// Return without an error does not mean the transfer success, you need to
// check the response Code.
//
// Ref: https://docs.duitku.com/disbursement/en/#online-transfer-transfer-request
func (cl *Client) RtolTransfer(req *RtolTransfer) (res *RtolTransferResponse, err error) {
	var (
		logp = `RtolTransfer`
		path = PathTransfer

		resHttp *http.Response
		resBody []byte
	)

	// Since the path is different in test environment, we check the host
	// here to set it.
	if cl.opts.host != hostLive {
		path = PathTransferSandbox
	}

	req.Sign(cl.opts)

	resHttp, resBody, err = cl.PostJSON(path, nil, req)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}
	if resHttp.StatusCode >= 500 {
		return nil, fmt.Errorf(`%s: %s`, logp, resHttp.Status)
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	// The actual transfer does not return the disburseID back, so we set
	// it here.
	res.DisburseID = req.DisburseID

	return res, nil
}
