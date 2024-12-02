package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// and struct type defination with json metadata
type Note struct {
	Title     string    `json:"title"` //struct tags to change output
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"create_at"`
}

// content to display
func (note Note) Display() {
	fmt.Printf("Your note title %v has the following content: \n %v", note.Title, note.Content)
}

// writting conent to file in json format
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// saving result with json format
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// checking if there is error
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}
