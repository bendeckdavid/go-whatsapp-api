package wsp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	conn "github.com/bendeckdavid/go-connector"
	"github.com/labstack/echo/v4"
)

const ig = "Event ignored because "

// Verification endpoint
func verifyHook(e *echo.Echo) {

	e.GET("/webhooks/whatsapp", func(c echo.Context) error {

		var hook webhook
		err := c.Bind(&hook)
		if err != nil || hook.Token != os.Getenv("WSP_WEBHOOK_TOKEN") {
			return conn.BuildError(err, http.StatusBadRequest).Send(c)
		}

		return c.String(http.StatusOK, fmt.Sprint(hook.Challenge))
	})

}

// Register webhooks to Echo server
func RegisterWebhook(s *conn.Server) *conn.Server {

	verifyHook(s.Server)
	s.Server.POST("/webhooks/whatsapp", func(c echo.Context) error {

		var ev event
		if err := c.Bind(&ev); err != nil {
			return conn.BuildError(err, http.StatusBadRequest).Send(c)
		}

		if s.Debug {
			debug, _ := json.Marshal(ev)
			log.Println(string(debug))
		}

		for _, entry := range ev.Entry {

			if entry.ID != getCredentials().BusinessID {
				log.Println(ig + "BusinessID mismatch")
				continue
			}

			for _, change := range entry.Changes {

				if change.Field != "messages" {
					log.Println(ig + change.Field + " changes are not implemented")
					continue
				}

				if change.Value.Product != "whatsapp" {
					log.Println(ig + "only listening whatsapp events")
					continue
				}

				for _, msg := range change.Value.Messages {

					msg := Message{
						ID:   msg.ID,
						From: msg.From,
						Text: msg.Text.Body,
					}

					conn.DB.Create(&msg)
					go eventHandler.onNewMessage(msg)
				}

			}

		}

		return nil
	})

	return s
}

type webhook struct {
	Mode      string `query:"hub.mode"`
	Challenge int    `query:"hub.challenge"`
	Token     string `query:"hub.verify_token"`
}

type event struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Value struct {
				Product  string `json:"messaging_product"`
				Messages []struct {
					From string `json:"from"`
					ID   string `json:"id"`
					Text struct {
						Body string `json:"body"`
					} `json:"text,omitempty"`
				} `json:"messages"`
				Statuses []struct {
					ID     string `json:"id"`
					Status string `json:"status"`
				} `json:"statuses,omitempty"`
			} `json:"value"`
			Field string `json:"field"`
		} `json:"changes"`
	} `json:"entry"`
}
