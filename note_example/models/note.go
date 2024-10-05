package note

type Note struct {
	Title string
	Body  string
}

func New(title, body string) *Note {
	return &Note{
		Title: title,
		Body:  body,
	}
}
