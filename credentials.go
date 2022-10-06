package whatsapp

import (
	"fmt"
	"net/http"
	"os"

	conn "github.com/BendeckDev/go-connector"
)

var credentials *Credentials

type Credentials struct {
	BusinessID string
	Token      string
	PhoneID    string
	Version    string
}

// Set global credentials
func SetCredentials(cd *Credentials) {
	if cd == nil {
		Credentials{
			BusinessID: os.Getenv("WSP_BUSINESS_ID"),
			Token:      os.Getenv("WSP_TOKEN"),
			PhoneID:    os.Getenv("WSP_PHONE_ID"),
			Version:    os.Getenv("WSP_VERSION"),
		}.Save()
	} else {
		cd.Save()
	}
}

// Make a Request with the Instance
func (cd Credentials) Request(req conn.Request) conn.Response {

	if req.Type == nil {
		req.Type = &conn.Get
	}
	req.Endpoint = fmt.Sprintf("https://graph.facebook.com/%s/%s", cd.Version, req.Endpoint)
	req.Headers = append(req.Headers, conn.Header{
		Name:  "Authorization",
		Value: "Bearer " + cd.Token,
	})

	return req.Make()
}

// Save credentials to be used globally
func (cd Credentials) Save() {
	credentials = cd.Validate()
}

// Validate credentials
func (cd Credentials) Validate() *Credentials {

	if req := cd.Request(conn.Request{
		Endpoint: "/" + cd.BusinessID,
	}); req.Status != http.StatusOK {
		panic("invalid Whatsapp credentials")
	}

	if req := cd.Request(conn.Request{
		Endpoint: "/" + cd.PhoneID,
	}); req.Status != http.StatusOK {
		panic("invalid phone number ID")
	}

	return &cd
}
