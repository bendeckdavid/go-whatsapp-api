package whatsapp

import conn "github.com/BendeckDev/go-connector"

type Request struct {
	To          string
	credentials *Credentials
	typeOf      string
	text        *textContent
	template    *templateContent
}

// Create new whatsapp request
func NewRequest(To string) *Request {
	return &Request{
		To: To,
	}
}

// Set custom credentials for a request
func (r *Request) WithCredentials(cd *Credentials) *Request {
	r.credentials = cd
	return r
}

// Set text content for request
func (r *Request) Text(text string) *Request {

	// Content
	r.typeOf = "text"
	r.text = &textContent{
		Preview: false,
		Text:    text,
	}

	return r
}

// Set template content for request
func (r *Request) Template(name string, lang string, bodyVars ...string) *Request {

	var parameters []parameterContent

	// Build body parameters
	for _, v := range bodyVars {
		parameters = append(parameters, parameterContent{
			Type: "text",
			Text: v,
		})
	}

	// Content
	r.typeOf = "template"
	r.template = &templateContent{
		Name: name,
		Language: struct {
			Code string "json:\"code\""
		}{Code: lang},
		Components: []struct {
			Type       string             "json:\"type\""
			Parameters []parameterContent "json:\"parameters\""
		}{{Type: "body", Parameters: parameters}},
	}

	return r
}

// Make whatsapp request
func (r *Request) Send() conn.Response {

	if r.credentials == nil {
		r.credentials = credentials
	}

	return r.credentials.Request(buildRequest(*r))
}
