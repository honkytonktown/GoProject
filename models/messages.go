package models

import (
	"fmt"

	"github.com/pkg/errors"
)

//message struct
type Message struct {
	Title string
	Body  string
	ID    int
}

var (
	messages []*Message
	mID      = 0
)

//returns array of message pointers
func GetMessages() []*Message {
	return messages
}

//adds a new message to message array
func AddMessage(m Message) (Message, error) {
	if m.ID != 0 {
		return Message{}, errors.New("The id field was not null.")
	}
	m.ID = mID
	mID++
	messages = append(messages, &m)
	return m, nil
}

//brute force searches for id match in messages array
func GetMessageByID(id int) (Message, error) {
	for _, candidate := range messages {
		if candidate.ID == id {
			return *candidate, nil
		}
	}
	return Message{}, fmt.Errorf("User with ID '%v' not found", id)
}

//updates entry by matching id
func UpdateMessage(m Message) (Message, error) {
	for i, candidate := range messages {
		if candidate.ID == m.ID {
			messages[i] = &m
			return m, nil
		}
	}
	return Message{}, fmt.Errorf("A user with ID '%v' was not found", m.ID)
}

//removes user by id
func RemoveMessageByID(id int) error {
	for i, candidate := range messages {
		if candidate.ID == id {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Message with id %v was not found", id)
}
