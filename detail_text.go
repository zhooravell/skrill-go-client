package skrill

import (
	"errors"
	"strings"
)

// The detail1_description combined with the detail1_text is shown in the more information field of the merchant account history CSV file.
type DetailText string

// Validate Skrill text detail
func (rcv DetailText) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("detail text name should not be blank")
	}

	return nil
}
