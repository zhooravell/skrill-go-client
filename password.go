package skrill

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const passwordMinLength = 8

// Type for Skrill API/MQI password.
// - At least 8 characters long
// - At least 1 alphabetic character (A-Z)
// - At least 1 non-alphabetic character (0-9, ., +, etc.)
// - Cannot be the same as your email address
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type Password string

// Validate Skrill password
func (rcv Password) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if passwordMinLength > len(data) {
		return errors.New(fmt.Sprintf("skrill API/MQI password is too short. It should have %d characters or more", passwordMinLength))
	}

	// letter regexp
	lr, err := regexp.Compile(`\pL`)

	if err != nil {
		return err
	}

	if !lr.MatchString(data) {
		return errors.New("skrill API/MQI password must include at least one letter")
	}
	// number regexp
	nr, err := regexp.Compile(`\pN`)

	if err != nil {
		return err
	}

	if !nr.MatchString(data) {
		return errors.New("skrill API/MQI password must include at least one number")
	}

	return nil
}
