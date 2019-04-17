package skrill

import "testing"

func TestRecurringBillingNote_Valid(t *testing.T) {
	rbn := RecurringBillingNote("text")

	if err := rbn.Valid(); err != nil {
		t.Error(err)
	}
}

func TestRecurringBillingNote_ValidEmpty(t *testing.T) {
	rbn := RecurringBillingNote("")

	if err := rbn.Valid(); err == nil {
		t.Error(err)
	}
}
