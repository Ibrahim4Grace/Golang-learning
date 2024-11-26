package main

import (
	"bufio"
	"fmt"
	"note/utils"
	"os"
	"strings"
)

func main() {

	title, content := getNotedata()
	userNote, err := utils.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("saving the note succesful.")
}

func getNotedata() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	//Handling long user input
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
