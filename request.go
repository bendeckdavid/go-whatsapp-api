package whatsapp

type Request struct {
	To       string
	Lang     string
	typeOf   string
	text     *textContent
	template *templateContent
}

func (r *Request) Text(text string) {

	// Content
	r.typeOf = "text"
	r.text = &textContent{
		Preview: false,
		Text:    text,
	}
}

func (r *Request) Template(name string, vars ...string) {

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
}

func (r *Request) Send(cd Credentials) {
	cd.Request(buildRequest(*r))
}
