package client

type Slack interface {
	GetUserID(user string) (string, error)
	GetChannelID(channel string) (string, error)
	SendBlockMsg(msg BlockMessage, ID string) error
	SendAttachment(attach Attachment, ID string) error
	GetConversationIDOfUser(userID string) (string, error)
}

type slack struct {
	botToken string
}

type Attachment struct {
}

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

var _ Slack = (*slack)(nil)

func newSlackClient() *slack {
	//TODO: implement this
	return &slack{}
}

func (s *slack) GetUserID(user string) (string, error) {
	//TODO: implement this
	return "", nil
}

func (s *slack) GetChannelID(channel string) (string, error) {
	//TODO: implement this
	return "", nil
}

func (s *slack) SendBlockMsg(msg BlockMessage, ID string) error {
	//TODO: implement this
	return nil
}

func (s *slack) SendAttachment(attach Attachment, ID string) error {
	//TODO: implement this
	return nil
}

func (s *slack) GetConversationIDOfUser(userID string) (string, error) {
	//TODO: implement this
	return "", nil
}
