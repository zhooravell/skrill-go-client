package skrill

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const secretWordMaxLength = 10

// Type for Skrill secret word.
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type SecretWord string

// Validate Skrill secret word
func (rcv SecretWord) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("secret word should not be blank")
	}

	if len(data) > secretWordMaxLength {
		return errors.New(fmt.Sprintf("the length of Skrill secret word should not exceed %d characters", secretWordMaxLength))
	}

	// special characters regexp
	scp, err := regexp.Compile(`[^a-zA-Z0-9\s]`)

	if err != nil {
		return err
	}

	if scp.MatchString(data) {
		return errors.New("special characters are not permitted in Skrill secret word")
	}

	return nil
}
