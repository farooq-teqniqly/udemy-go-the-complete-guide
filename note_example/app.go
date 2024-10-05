package main

import (
	"bufio"
	"fmt"
	"os"
	"teqniqly.com/note_example/models"
)

func main() {
	title := getUserInput("Note title:")
	body := getUserInput("Note body:")

	newNote := note.New(title, body)
	_, _ = fmt.Printf("Your noted titled %s has the following content:\n", newNote.Title)
	fmt.Println(newNote.Body)

	json, err := newNote.JSONEncode()

	if err != nil {
		fmt.Println("Error converting note to JSON:", err)
		os.Exit(-1)
	}

	err = writeToFile(json, "note.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Note saved to file!")
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func writeToFile(content []byte, filename string) error {
	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		return err
	}

	return nil
}
