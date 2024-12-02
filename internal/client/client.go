package client

type Client interface {
	GetUserID()
	GetChannelID()
	SendTextMsg()
	SendBlockMsg()
}
