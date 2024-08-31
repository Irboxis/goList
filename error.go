package goList

const (
	syntax     = 100
	parameters = 200
	others     = 300
)

type error struct {
	code  int
	title string
	msg   string
	sug   string
}

func (err *error) New(message error) *error {
	return &error{
		code:  message.code,
		title: message.title,
		msg:   message.msg,
		sug:   message.sug,
	}
}
