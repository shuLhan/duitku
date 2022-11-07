package duitku

// ClearingInquiryResponse contains response from calling [Clearing Inquiry
// request].
//
// [Clearing Inquiry request]: https://docs.duitku.com/disbursement/en/#clearing-inquiry-request
type ClearingInquiryResponse struct {
	Type string `json:"type"`

	RtolInquiryResponse
}
