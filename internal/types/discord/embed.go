package discord

type Embed struct {
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
}

type EmbedField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
