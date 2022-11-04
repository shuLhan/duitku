package duitku

import "github.com/shuLhan/share/lib/math/big"

// Balance contains the current user balances.
type Balance struct {
	Response

	// Current balance before settlement.
	Current *big.Rat `json:"balance"`

	// Effective Balance that can be used for disbursement.
	Effective *big.Rat `json:"effectiveBalance"`

	Email  string `json:"email"`
	UserID int64  `json:"userId"`
}
