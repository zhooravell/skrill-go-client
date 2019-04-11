package skrill

import (
	"errors"
	"strings"
)

// Type for transaction id (transaction_id or mb_transaction_id).
// A unique reference or identification number.
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type Reference string

// Validate reference
func (rcv Reference) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("reference should not be blank")
	}

	return nil
}
