package wsp

import "time"

type Message struct {
	ID        string
	From      string
	Text      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var eventHandler = struct {
	onNewMessage   func(msg Message)
	onStatusChange func(msg Message)
}{}

// Set function to be called when a new message arrived
func NewMessageCallback(callback func(msg Message)) {
	eventHandler.onNewMessage = callback
}

// Set function to be called when a message status changed
func StatusChangeCallback(callback func(msg Message)) {
	eventHandler.onStatusChange = callback
}
