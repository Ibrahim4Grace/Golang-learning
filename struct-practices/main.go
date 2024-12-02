package main

import (
	"bufio"
	"fmt"
	note "note/note"
	"os"
	"strings"
)

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("hello")

	title, content := getNotedata()
	userNote, err := note.New(title, content)

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

// To accepts any kind of value int or string
func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("integer", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float64", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("string", stringVal)
		return
	}

	fmt.Println(value)
}

// This can be use in a case you dont know the input type
func add[T int | float64 | string](a, b T) T {
	return a + b
}

func output() {
	result := add(1, 2)
	fmt.Println(result)
}
