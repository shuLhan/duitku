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

const (
	PathDisbursementListBank = `/disbursement/listBank`
)

type Client struct {
	*libhttp.Client

	opts ClientOptions
}

// NewClient create and initialize new Client.
func NewClient(opts ClientOptions) (cl *Client, err error) {
	var (
		httpcOpts = libhttp.ClientOptions{
			ServerUrl: opts.ServerUrl,
		}
	)

	cl = &Client{
		Client: libhttp.NewClient(&httpcOpts),
		opts:   opts,
	}

	return cl, nil
}

// DisbursementListBank fetch list of banks for disbursement.
func (cl *Client) DisbursementListBank() (banks []Bank, err error) {
	var (
		logp = `DisbursementListBank`
		req  = createRequest(cl.opts)
		res  = struct {
			Code string      `json:"responseCode"`
			Desc string      `json:"responseDesc"`
			Data interface{} `json:"Banks"`
		}{}

		httpRes *http.Response
		resBody []byte
	)

	httpRes, resBody, err = cl.PostJSON(PathDisbursementListBank, nil, req)
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
	if res.Code != resCodeSuccess {
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
