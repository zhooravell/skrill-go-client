package skrill

type payload struct {
	data map[string]interface{}
}

func (rcv *payload) Payload() map[string]interface{} {
	return rcv.data
}
