package skrill

import (
	"errors"
	"strings"
)

// Type for Skrill signature (md5sig or sha2sig).
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type Signature string

// Validate signature
func (rcv Signature) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("signature should not be blank")
	}

	if strings.ToUpper(data) != data {
		return errors.New("signature should in uppercase")
	}

	return nil
}
