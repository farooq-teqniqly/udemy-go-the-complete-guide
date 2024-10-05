package note

import "encoding/json"

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func New(title, body string) *Note {
	return &Note{
		Title: title,
		Body:  body,
	}
}

func (note *Note) JSONEncode() ([]byte, error) {
	jsonData, err := json.MarshalIndent(note, "", "  ")

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
