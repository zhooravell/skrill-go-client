package skrill

// The detail1_description combined with the detail1_text is shown in the more information field of the merchant account history CSV file.
type DetailDescription string

// Validate Skrill description detail
func (rcv DetailDescription) Valid() error {
	return nil
}
