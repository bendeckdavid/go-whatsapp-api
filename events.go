package wsp

type Message struct {
	From string
	Text string
}

var eventHandler = struct {
	onNewMessage func(msg Message)
}{}

func NewMessageCallback(callback func(msg Message)) {
	eventHandler.onNewMessage = callback
}
