package skrill

import "testing"

func TestRecurringPaymentID_Valid(t *testing.T) {
	rbID := RecurringPaymentID("text")

	if err := rbID.Valid(); err != nil {
		t.Error(err)
	}
}

func TestRecurringPaymentID_ValidEmpty(t *testing.T) {
	rbID := RecurringPaymentID("")

	if err := rbID.Valid(); err == nil {
		t.Error(err)
	}
}
