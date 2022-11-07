// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

// Package duitku provide library and HTTP client for [duitku.com].
//
// [duitku.com]: https://docs.duitku.com/disbursement/id/#langkah-awal
package duitku

const (
	ServerUrlLive    = `https://passport.duitku.com/webapi/api`
	ServerUrlSandbox = `https://sandbox.duitku.com/webapi/api`

	hostLive = `passport.duitku.com`
)

const (
	// LLG (Lalu Lintas Giro) is interbank transfer that cover more than
	// 130 bank in Indonesia.
	// The maximal amount transfer is IDR 500.000.000.
	// Transfer process follows the BI (Bank Indonesia) schedule, which is
	// 8.00-15.00 on business days.
	ClearingTypeLLG = `LLG`

	// RTGS (Real Time Gross Settlement) is interbank transfer that cover
	// more than 130 bank in Indonesia.
	// The minimal amount transfer is IDR 100.000.000.
	// Transfer process follows the BI (Bank Indonesia) schedule, which is
	// 8.00-15.00 on business days.
	ClearingTypeRTGS = `RTGS`

	// H2H (Bank Host to Host) Duitku Host to Host connection to bank, to
	// ensure direct connection and better reliability.
	// Currently only support 4 Major banks in Indonesia (BNI, BRI,
	// Mandiri, Permata).
	// Transfer schedule follows the schedule of each bank.
	ClearingTypeH2H = `H2H`

	// BIFAST (Bank Indonesia Fast Payment) National retail payments that
	// can facilitate retail payments in real-time, safe, efficient, more
	// affordable service fees and available at any time (24/7).
	ClearingTypeBIFAST = `BIFAST`
)
