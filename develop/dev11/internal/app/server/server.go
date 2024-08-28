package server

import (
	"github.com/nglmq/calendar/internal/app/domain"
	"github.com/nglmq/calendar/internal/app/handlers"
	middleware "github.com/nglmq/calendar/internal/app/middleware/logger"
	"github.com/nglmq/calendar/internal/storage/memory"
	"net/http"
)

func NewServer() (http.Handler, error) {
	storage := memory.NewInMemoryStorage()
	domain.SetStorage(storage)

	r := http.NewServeMux()

	r.HandleFunc("/create_event", handlers.CreateEvent)
	r.HandleFunc("/update_event", handlers.UpdateEvent)
	r.HandleFunc("/delete_event", handlers.DeleteEvent)
	r.HandleFunc("/events_for_day", handlers.EventsForDay)
	r.HandleFunc("/events_for_week", handlers.EventsForWeek)
	r.HandleFunc("/events_for_month", handlers.EventsForMonth)

	return middleware.RequestLogger(r), nil
}
