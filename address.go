// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// Address contains detailed [address] of customer.
//
// [address]: https://docs.duitku.com/api/en/#address
type Address struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
}
