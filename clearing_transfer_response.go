package duitku

// ClearingTransferResponse contains response from Clearing Transfer.
type ClearingTransferResponse struct {
	Type string `json:"type"`

	RtolTransferResponse
}
