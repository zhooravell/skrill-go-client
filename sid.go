package skrill

import (
	"errors"
	"strings"
)

// Type for sid.
type Sid string

// Validate Skrill sid
func (rcv Sid) Valid() error {
	data := strings.TrimSpace(string(rcv))

	if "" == data {
		return errors.New("sid should not be blank")
	}

	return nil
}
