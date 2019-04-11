package skrill

// Type for Skrill API/MQI password.
// - At least 8 characters long
// - At least 1 alphabetic character (A-Z)
// - At least 1 non-alphabetic character (0-9, ., +, etc.)
// - Cannot be the same as your email address
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type Password string

// Validate Skrill password
func (rcv Password) Valid() error {
	return nil
}
