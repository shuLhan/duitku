package duitku

// RtolTransferResponse contains response from online transfer.
type RtolTransferResponse struct {
	Purpose string `json:"purpose"`

	RtolInquiryResponse

	UserID int64 `json:"userId"`
}
