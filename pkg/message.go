package minify

import "fmt"

type Message struct {
	Code string
	Text string
}

func (m Message) String() string {
	return fmt.Sprintf("%v: %v", m.Code, m.Text)
}

func (m Message) Error() string {
	return fmt.Sprintf("%v: %v", m.Code, m.Text)
}
