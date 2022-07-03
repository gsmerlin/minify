package minify

type Error struct {
	Message
}

var Errors = map[string]Error{
	"NOT_FOUND": {
		Message{
			Code: "NOT_FOUND",
			Text: "Page not found",
		},
	},
}

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
