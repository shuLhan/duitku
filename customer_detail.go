// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// [CustomerDetail] detail of customer information for payment to merchant.
//
// [CustomerDetail]: https://docs.duitku.com/api/en/#customer-detail
type CustomerDetail struct {
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	Email           string  `json:"email"`
	PhoneNumber     string  `json:"phoneNumber"`
	BillingAddress  Address `json:"billingAddress"`
	ShippingAddress Address `json:"shippingAddress"`
}
