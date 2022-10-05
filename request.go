package whatsapp

import conn "github.com/BendeckDev/go-connector"

type Request struct {
	To       string
	Lang     string
	typeOf   string
	text     *textContent
	template *templateContent
}

func NewRequest(To string, Lang string) *Request {
	return &Request{
		To:   To,
		Lang: Lang,
	}
}

func (r *Request) Text(text string) *Request {

	// Content
	r.typeOf = "text"
	r.text = &textContent{
		Preview: false,
		Text:    text,
	}

	return r
}

func (r *Request) Template(name string, vars ...string) *Request {

	var parameters []parameterContent

	// Build body parameters
	for _, v := range vars {
		parameters = append(parameters, parameterContent{
			Type: "text",
			Text: v,
		})
	}

	// Content
	r.typeOf = "template"
	r.template = &templateContent{
		Name: name,
		Components: []struct {
			Type       string             "json:\"type\""
			Parameters []parameterContent "json:\"parameters\""
		}{{Type: "body", Parameters: parameters}},
	}

	return r
}

func (r *Request) Send(cd Credentials) conn.Response {
	return cd.Request(buildRequest(*r))
}
