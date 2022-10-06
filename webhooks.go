package whatsapp

import (
	conn "github.com/BendeckDev/go-connector"
	"github.com/labstack/echo/v4"
)

// Register webhooks to Echo server
func RegisterWebhook(s *conn.Server) *conn.Server {

	s.Server.POST("/webhooks/whatsapp", func(c echo.Context) error {
		eventHandler.onNewMessage(Message{
			Text: "Test",
		})
		return nil
	})

	return s
}
