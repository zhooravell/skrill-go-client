package skrill

import (
	"errors"
	"strings"
)

// Type for Skrill On demand node (ondemand_note).
type RecurringBillingNote string

// Validate Skrill description detail
func (rcv RecurringBillingNote) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("recurring billing note should not be blank")
	}

	return nil
}
