package models

import (
	"errors"
)

type Event struct {
	ID    int
	Name  string
	Users []User
}

type User struct {
	Name  string
	Email string
	Phone string
}

type EventManagerInterface interface {
	ListEvents() []Event
	GetEvent(index int) (*Event, error)
	RegisterUser(eventIndex int, user User) error
}

type EventManager struct {
	events []Event
}

func NewEventManager() *EventManager {
	return &EventManager{
		events: []Event{
			{ID: 0, Name: "Volleyball", Users: []User{}},
			{ID: 1, Name: "Baseball", Users: []User{}},
			{ID: 2, Name: "Table Tennis", Users: []User{}},
		},
	}
}

func (e *EventManager) ListEvents() []Event {
	return e.events
}

func (e *EventManager) GetEvent(index int) (*Event, error) {
	if index < 0 || index >= len(e.events) {
		return nil, errors.New("event not found")
	}
	return &e.events[index], nil
}

func (e *EventManager) RegisterUser(eventIndex int, user User) error {
	if eventIndex < 0 || eventIndex >= len(e.events) {
		return errors.New("event not found")
	}
	e.events[eventIndex].Users = append(e.events[eventIndex].Users, user)
	return nil
}
