// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	PathTransfer        = `/webapi/api/disbursement/transfer`
	PathTransferSandbox = `/webapi/api/disbursement/transfersandbox` // Used for testing.

	// Paths for Clearing.
	PathInquiryClearing        = `/webapi/api/disbursement/inquiryclearing`
	PathInquiryClearingSandbox = `/webapi/api/disbursement/inquiryclearingsandbox` // Used for testing.

	PathTransferClearing        = `/webapi/api/disbursement/transferclearing`
	PathTransferClearingSandbox = `/webapi/api/disbursement/transferclearingsandbox` // Used for testing.
)

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

	err = opts.validate()
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
		req  = CreateRequest(cl.opts)

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
	if bal.Code != ResCodeSuccess {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, bal.Code, bal.Desc)
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
	if res.Code != ResCodeSuccess {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, res.Code, res.Desc)
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

// tListBank fetch list of banks for disbursement.
func (cl *Client) ListBank() (banks []Bank, err error) {
	var (
		logp = `ListBank`
		req  = CreateRequest(cl.opts)
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
		var cmp int = strings.Compare(banks[x].Code, banks[y].Code)
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
	if res.Code != ResCodeSuccess {
		return nil, fmt.Errorf(`%s: %s: %s`, logp, res.Code, res.Desc)
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

	return res, nil
}
