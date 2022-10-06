package whatsapp

type Message struct {
	Text string
}

var eventHandler = struct {
	onNewMessage func(msg Message)
}{}

func NewMessageCallback(callback func(msg Message)) {
	eventHandler.onNewMessage = callback
}
