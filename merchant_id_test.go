package skrill

import "testing"

func TestMerchantID_Valid(t *testing.T) {
	id := MerchantID(1)

	if err := id.Valid(); err != nil {
		t.Error(err)
	}
}
