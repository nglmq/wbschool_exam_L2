package memory

import (
	"fmt"
	"github.com/nglmq/calendar/internal/app/domain"
	"sync"
	"time"
)

type InMemoryStorage struct {
	Events map[string]domain.Event
	mx     sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Events: make(map[string]domain.Event),
	}
}

func (c *InMemoryStorage) SaveEvent(event domain.Event) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.Events[event.ID] = event
}

func (c *InMemoryStorage) GetEvent(id, userID string) (domain.Event, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	event, exists := c.Events[id]
	if !exists || event.UserID != userID {
		return domain.Event{}, fmt.Errorf("event not found")
	}

	return event, nil
}

func (c *InMemoryStorage) UpdateEvent(event domain.Event) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	if existingEvent, exists := c.Events[event.ID]; !exists || existingEvent.UserID != event.UserID {
		return fmt.Errorf("event not found")
	}

	c.Events[event.ID] = event

	return nil
}

func (c *InMemoryStorage) DeleteEvent(id, userID string) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	if event, exists := c.Events[id]; !exists || event.UserID != userID {
		return fmt.Errorf("event not found")
	}

	delete(c.Events, id)
	return nil
}

func (c *InMemoryStorage) GetEventsByDay(userID, date string) ([]domain.Event, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	var events []domain.Event
	for _, event := range c.Events {
		if event.Date == date && event.UserID == userID {
			events = append(events, event)
		}
	}

	return events, nil
}

func (c *InMemoryStorage) GetEventsForRange(userID string,
	startDate,
	endDate time.Time) ([]domain.Event, error) {

	c.mx.RLock()
	defer c.mx.RUnlock()

	var events []domain.Event
	for _, event := range c.Events {
		eventDate, _ := time.Parse("02-01-2006", event.Date)
		if event.UserID == userID && !eventDate.Before(startDate) && !eventDate.After(endDate) {
			events = append(events, event)
		}
	}

	return events, nil
}
