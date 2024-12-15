package slack

type BlockMessage struct {
	// Required fields
	Channel string   `json:"channel"`
	Blocks  []*Block `json:"blocks"`
	Text    string   `json:"text"`
	Files   []string `json:"files,omitempty"`
}

type Block struct {
	Type     string      `json:"type"`
	Text     *BlockText  `json:"text,omitempty"`
	Fields   []BlockText `json:"fields,omitempty"`
	Title    *BlockText  `json:"title,omitempty"`
	File     *SlackFile  `json:"slack_file,omitempty"`
	AltText  string      `json:"alt_text,omitempty"`
	Elements []Element   `json:"elements,omitempty"`
}

type BlockText struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

type Element struct {
	Type     string     `json:"type"`
	Text     *BlockText `json:"text,omitempty"`
	Value    string     `json:"value,omitempty"`
	URL      string     `json:"url,omitempty"`
	ActionID string     `json:"action_id,omitempty"`
}

type SlackFile struct {
	ID string `json:"id"`
}
