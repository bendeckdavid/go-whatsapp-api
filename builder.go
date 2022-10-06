package whatsapp

import conn "github.com/BendeckDev/go-connector"

func buildRequest(r Request) conn.Request {

	return conn.Request{
		Endpoint: getCredentials().PhoneID + "/messages",
		Type:     &conn.Post,
		Body: struct {
			Product   string           `json:"messaging_product"`
			Recipient string           `json:"recipient_type"`
			To        string           `json:"to"`
			Type      string           `json:"type"`
			Template  *templateContent `json:"template,omitempty"`
			Text      *textContent     `json:"text,omitempty"`
		}{
			Product:   "whatsapp",
			Recipient: "individual",
			To:        r.To,
			Type:      r.typeOf,
			Template:  r.template,
			Text:      r.text,
		},
	}
}
