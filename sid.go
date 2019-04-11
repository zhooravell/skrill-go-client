package skrill

// Type for sid.
type Sid string

// Validate Skrill sid
func (rcv Sid) Valid() error {
	return nil
}
