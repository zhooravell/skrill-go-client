package skrill

import (
	"errors"
	"strings"
)

// Type for language.
// Skrill supports the following languages (2-character ISO codes).
// https://en.wikipedia.org/wiki/ISO_639-1
// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
type Language string

// Validate Skrill language
func (rcv Language) Valid() error {
	data := strings.TrimSpace(string(rcv))

	switch data {
	case
		"BG",
		"ZH",
		"CS",
		"DA",
		"NL", // from the endonym Nederlands
		"EN",
		"FI",
		"FR",
		"EL",
		"IT",
		"JA",
		"PL",
		"PT",
		"RO",
		"RU",
		"ES",
		"SV",
		"TR":
		return nil
	}

	return errors.New("not accepted language by Skrill")
}
