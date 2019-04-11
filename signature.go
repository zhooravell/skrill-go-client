package skrill

// Type for Skrill signature (md5sig or sha2sig).
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type Signature string

// Validate signature
func (rcv Signature) Valid() error {
	return nil
}
