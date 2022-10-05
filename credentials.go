package whatsapp

import (
	"fmt"
	"net/http"

	conn "github.com/BendeckDev/go_connector"
)

type Credentials struct {
	PhoneID string
	Token   string
	Version string
}

// Make a Request with the Instance
func (cd Credentials) Request(req conn.Request) conn.Response {

	(req.Type) = &conn.Get
	req.Endpoint = fmt.Sprintf("https://graph.facebook.com/%v/%s%s", cd.Version, cd.PhoneID, req.Endpoint)
	req.Headers = append(req.Headers, conn.Header{
		Name:  "Authorization",
		Value: "Bearer " + cd.Token,
	})

	return req.Make()
}

// Validate credentials
func (i Credentials) Validate() Credentials {

	if req := i.Request(conn.Request{
		Endpoint: "/",
	}); req.Status != http.StatusOK {
		panic("invalid Whatsapp credentials")
	}

	return i
}
