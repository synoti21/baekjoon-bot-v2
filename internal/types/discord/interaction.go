package discord

type InteractionRequest struct {
	Type   int    `json:"type"`
	Token  string `json:"token"`
	Member struct {
		User struct {
			ID string `json:"id"`
		} `json:"user"`
	} `json:"member"`
	Data struct {
		Name    string `json:"name"`
		Options []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"options"`
	} `json:"data"`
}

type InteractionResponse struct {
	Type int                     `json:"type"`
	Data InteractionResponseData `json:"data,omitempty"`
}

type InteractionResponseData struct {
	Content string  `json:"content,omitempty"`
	Embeds  []Embed `json:"embeds,omitempty"`
}
