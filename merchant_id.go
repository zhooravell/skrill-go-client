package skrill

// Type for merchant id (merchant_id).
// Unique ID of your Skrill account.
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type MerchantID int

// Validate Skrill merchant ID
func (rcv MerchantID) Valid() error {
	return nil
}
