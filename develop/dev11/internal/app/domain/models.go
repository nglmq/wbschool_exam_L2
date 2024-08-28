package domain

import (
	"errors"
	"fmt"
	"time"
)

type Event struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
}

type Storage interface {
	SaveEvent(event Event)
	GetEvent(id, userID string) (Event, error)
	UpdateEvent(event Event) error
	DeleteEvent(id, userID string) error
	GetEventsByDay(userID, date string) ([]Event, error)
	GetEventsForRange(userID string, startDate, endDate time.Time) ([]Event, error)
}

var storage Storage

func SetStorage(s Storage) {
	storage = s
}

func CreateEvent(event Event) error {
	storage.SaveEvent(event)

	return nil
}

func UpdateEvent(event Event) error {
	err := storage.UpdateEvent(event)
	if err != nil {
		return fmt.Errorf("error updating event")
	}
	return nil
}

func DeleteEvent(id, userID string) error {
	err := storage.DeleteEvent(id, userID)
	if err != nil {
		return fmt.Errorf("error deleting event")
	}
	return nil
}

func GetEventsForDay(userID, date string) ([]Event, error) {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	return storage.GetEventsByDay(userID, date)
}

func GetEventsForWeek(userID, startDate string) ([]Event, error) {
	parsedDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	return storage.GetEventsForRange(userID, parsedDate, parsedDate.AddDate(0, 0, 7))
}

func GetEventsForMonth(userID, startDate string) ([]Event, error) {
	parsedDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	return storage.GetEventsForRange(userID, parsedDate, parsedDate.AddDate(0, 1, 0))
}
