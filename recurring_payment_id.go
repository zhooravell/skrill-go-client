package skrill

// Type for Skrill recurring payment id (rec_payment_id).
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Wallet_Checkout_Guide.pdf
type RecurringPaymentID string

// Validate recurring payment ID
func (rcv RecurringPaymentID) Valid() error {
	return nil
}
