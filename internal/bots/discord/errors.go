package discord

type InitError struct {
	err error
}

type UserNotFoundError struct {
	err error
}

type ChannelNotFoundError struct {
	err error
}

func (e *InitError) Error() string {
	return e.err.Error()
}

func (e *UserNotFoundError) Error() string {
	return e.err.Error()
}

func (e *ChannelNotFoundError) Error() string {
	return e.err.Error()
}
