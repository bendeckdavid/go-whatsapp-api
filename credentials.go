package whatsapp

import (
	"fmt"
	"net/http"
	"os"

	conn "github.com/BendeckDev/go-connector"
)

var credentials *Credentials

type Credentials struct {
	PhoneID string
	Token   string
	Version string
}

// Read credentials from environment variables
func getCredentials() *Credentials {

	if credentials == nil {
		credentials = Credentials{
			PhoneID: os.Getenv("WSP_PHONE_ID"),
			Token:   os.Getenv("WSP_TOKEN"),
			Version: os.Getenv("WSP_VERSION"),
		}.Validate()
	}

	return credentials
}

// Make a Request with the Instance
func (cd Credentials) Request(req conn.Request) conn.Response {

	if req.Type == nil {
		req.Type = &conn.Get
	}
	req.Endpoint = fmt.Sprintf("https://graph.facebook.com/%v/%s%s", cd.Version, cd.PhoneID, req.Endpoint)
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
		Endpoint: "/",
	}); req.Status != http.StatusOK {
		panic("invalid Whatsapp credentials")
	}

	return &cd
}
