package skrill

import (
	"errors"
	"strings"
)

const maxCompanyNameLength = 30

// Type for company name (recipient_description).
// A company name to be shown on the Skrill payment page in the logo area if there is no logo_url parameter.
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type CompanyName string

// Validate company name
func (rcv CompanyName) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("company name should not be blank")
	}

	if len(data) > maxCompanyNameLength {
		return errors.New("the length of company name should not exceed 30 characters")
	}

	return nil
}
