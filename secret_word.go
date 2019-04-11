package skrill

// Type for Skrill secret word.
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type SecretWord string

// Validate Skrill secret word
func (rcv SecretWord) Valid() error {
	return nil
}
