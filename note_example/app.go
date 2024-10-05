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
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}
