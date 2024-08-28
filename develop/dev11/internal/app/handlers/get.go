package handlers

import (
	"encoding/json"
	"github.com/nglmq/calendar/internal/app/domain"
	"net/http"
)

// EventsForDay обрабатывает GET запросы для получения событий на день
func EventsForDay(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	if userID == "" || date == "" {
		http.Error(w, "missing 'user_id' or 'date' parameter", http.StatusBadRequest)
		return
	}

	events, err := domain.GetEventsForDay(userID, date)
	if err != nil {
		http.Error(w, "error retrieving events", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]domain.Event{"result": events})
}

// EventsForWeek обрабатывает GET запросы для получения событий на неделю
func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	if userID == "" || date == "" {
		http.Error(w, "missing 'user_id' or 'date' parameter", http.StatusBadRequest)
		return
	}

	events, err := domain.GetEventsForWeek(userID, date)
	if err != nil {
		http.Error(w, "error retrieving events", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]domain.Event{"result": events})
}

// EventsForMonth обрабатывает GET запросы для получения событий на месяц
func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	if userID == "" || date == "" {
		http.Error(w, "missing 'user_id' or 'date' parameter", http.StatusBadRequest)
		return
	}

	events, err := domain.GetEventsForMonth(userID, date)
	if err != nil {
		http.Error(w, "error retrieving events", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]domain.Event{"result": events})
}
