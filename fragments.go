package whatsapp

type (
	textContent struct {
		Preview bool   `json:"preview_url"`
		Text    string `json:"body"`
	}

	templateContent struct {
		Name     string `json:"name"`
		Language struct {
			Code string `json:"code"`
		} `json:"language"`
		Components []struct {
			Type       string             `json:"type"`
			Parameters []parameterContent `json:"parameters"`
		} `json:"components"`
	}

	parameterContent struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
)
