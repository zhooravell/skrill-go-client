package skrill

import (
	"errors"
	"strings"
)

// Type for Skrill recurring payment id (rec_payment_id).
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Wallet_Checkout_Guide.pdf
type RecurringPaymentID string

// Validate recurring payment ID
func (rcv RecurringPaymentID) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("recurring payment ID should not be blank")
	}

	return nil
}
