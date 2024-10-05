package note

import "encoding/json"

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

func (note *Note) Json() (string, error) {
	jsonData, err := json.Marshal(note)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
