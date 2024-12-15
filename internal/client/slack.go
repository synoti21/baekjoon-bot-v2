package client

import "github.com/synoti21/baekjoon-slack-bot/internal/types/slack"

type SlackClient interface {
	GetUserID(user string) (string, error)
	GetChannelID(channel string) (string, error)
	SendBlockMsg(msg slack.BlockMessage, ID string) error
	SendAttachment(attach slack.Attachment, ID string) error
	GetConversationIDOfUser(userID string) (string, error)
}

type slackClient struct {
	botToken string
}

var _ SlackClient = (*slackClient)(nil)

func NewSlackClient() *slackClient {
	//TODO: implement this
	return &slackClient{}
}

func (s *slackClient) GetUserID(user string) (string, error) {
	//TODO: implement this
	return "", nil
}

func (s *slackClient) GetChannelID(channel string) (string, error) {
	//TODO: implement this
	return "", nil
}

func (s *slackClient) SendBlockMsg(msg slack.BlockMessage, ID string) error {
	//TODO: implement this
	return nil
}

func (s *slackClient) SendAttachment(attach slack.Attachment, ID string) error {
	//TODO: implement this
	return nil
}

func (s *slackClient) GetConversationIDOfUser(userID string) (string, error) {
	//TODO: implement this
	return "", nil
}
