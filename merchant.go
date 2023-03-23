// SPDX-FileCopyrightText: 2023 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package duitku

// Merchant the merchant configuration for payment.
//
// You can get the values of Merchant from project you registered on the
// [merchant portal].
//
// [merchant portal]: https://passport.duitku.com/merchant/Project
type Merchant struct {
	// The merchant code is the project code obtained from the Duitku
	// merchant page.
	// This code is useful as an identifier of your project in each
	// request using the /merchant/* APIs.
	Code string `ini:"::code"`

	// The API key for signing merchant request.
	ApiKey string `ini:"::api_key"`

	// The URL that will be used by Duitku to confirm payments made by
	// your customers.
	CallbackUrl string `ini:"::callback_url"`

	// The URL that Duitku will direct the customer after the transaction
	// is successful or canceled.
	ReturnUrl string `ini:"::return_url"`
}
