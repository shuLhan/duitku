// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// [AccountLink] Parameter for payment methods that use OVO Account Link and
// Shopee Account Link.
//
// [AccountLink]: https://docs.duitku.com/api/en/#account-link
type AccountLink struct {
	// [REQ] Credential Code provide by Duitku.
	CredentialCode string `json:"credentialCode"`

	// [REQ] Mandatory for OVO payment.
	OVO AccountLinkOvo `json:"ovo"`

	// [REQ] Mandatory for Shopee payment.
	Shopee AccountLinkShopee `json:"shopee"`
}

type AccountLinkOvo struct {
	PaymentDetails []OvoPaymentDetail `json:"paymentDetails"`
}

// [AccountLinkOvo] payment detail with OVO.
//
// [AccountLinkOvo]: https://docs.duitku.com/api/en/#ovo-detail
type OvoPaymentDetail struct {
	// [REQ] Type of your payment.
	PaymentType string `json:"paymentType"`

	// [REQ] Transaction payment amount.
	Amount int64 `json:"amount"`
}

// [AccountLinkShopee] payment detail with Shopee.
//
// [AccountLinkShopee]: https://docs.duitku.com/api/en/#shopee-detail
type AccountLinkShopee struct {
	// [REQ] Voucher code.
	PromoIDs string `json:"promo_ids"`

	// [REQ] Used for shopee coin from linked ShopeePay account.
	// Set true when pay transaction would like to use coins (Only for
	// ShopeePay account link).
	UseCoin bool `json:"useCoin"`
}
