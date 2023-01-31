// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// MerchantInquiry define request data for payment using merchant.
type MerchantInquiry struct {
	// [REQ] Transaction number from merchant.
	// Every request for a new transaction must use a new ID.
	MerchantOrderId string `json:"merchantOrderId"`

	// [REQ] PaymentMethod type of payment.
	//
	// Ref: https://docs.duitku.com/api/en/#payment-method
	PaymentMethod string `json:"paymentMethod"`

	// [REQ] Description about product/service on sale.
	//
	// You can fill in ProductDetails with a description of the product or
	// service that you provide.
	// You can also insert your store or brand name for more details.
	// Then in ItemDetails you can fill in product variants or product
	// model details, and other details about the products/services listed
	// in the transaction.
	ProductDetails string `json:"productDetails"`

	// [REQ] Customer's email.
	Email string `json:"email"`

	// [REQ] The name that would be shown at bank payment system.
	CustomerVaName string `json:"customerVaName"`

	// [OPT] Additional parameter to be used by merchant.
	// If its set, the value must be URL encoded.
	AdditionalParam string `json:"additionalParam,omitempty"`

	// [OPT] Customer's username.
	MerchantUserInfo string `json:"merchantUserInfo,omitempty"`

	// [OPT] Customer's phone number.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// [OPT] Customer's details.
	// [REQ] If PaymentMethod is Credit (DN/AT).
	CustomerDetail *CustomerDetail `json:"customerDetail,omitempty"`

	// [OPT] Details for payment method.
	AccountLink *AccountLink `json:"accountLink,omitempty"`

	// [OPT] Detail of product being payed.
	// [REQ] If PaymentMethod is Credit (DN/AT).
	//
	// The total of all price in ItemDetails must exactly match the
	// PaymentAmount.
	ItemDetails []ItemDetail `json:"itemDetails,omitempty"`

	// [REQ] Amount of transaction.
	//
	// Make sure the PaymentAmount is equal to the total Price in the
	// ItemDetails.
	PaymentAmount int64 `json:"paymentAmount"`
}

// merchantInquiry contains internal fields that will be set by client
// during Sign.
type merchantInquiry struct {
	// [REQ] A link for callback transaction.
	// Default to ClientOptions.MerchantCallbackUrl.
	CallbackUrl string `json:"callbackUrl"`

	// [REQ] MerchantCode is a project that use Duitku.
	//
	// You can get this code on every project you register on the
	// [merchant portal].
	// Default to ClientOptions.MerchantCode.
	//
	// [merchant portal]: https://passport.duitku.com/merchant/Project
	MerchantCode string `json:"merchantCode"`

	// [REQ] A link that is used for redirect after exit payment page,
	// being paid or not.
	// Default to ClientOptions.MerchantReturnUrl.
	ReturnUrl string `json:"returnUrl"`

	// [REQ] Transaction security identification code.
	// Formula: MD5(merchantCode + merchantOrderId + paymentAmount + apiKey).
	Signature string `json:"signature"`

	MerchantInquiry

	// [OPT] Transaction expiry period in minutes.
	// If its empty, it will set to [default] based on PaymentMethod.
	//
	// [default]: https://docs.duitku.com/api/en/#expiry-period
	ExpiryPeriod int `json:"expiryPeriod,omitempty"`
}

func (inq *merchantInquiry) sign(opts ClientOptions) {
	var merchant = opts.Merchant(inq.PaymentMethod)

	inq.CallbackUrl = merchant.CallbackUrl
	inq.MerchantCode = merchant.Code
	inq.ReturnUrl = merchant.ReturnUrl

	var (
		plain    = fmt.Sprintf(`%s%s%d%s`, inq.MerchantCode, inq.MerchantOrderId, inq.PaymentAmount, merchant.ApiKey)
		plainmd5 = md5.Sum([]byte(plain))
	)

	inq.Signature = hex.EncodeToString(plainmd5[:])
}
