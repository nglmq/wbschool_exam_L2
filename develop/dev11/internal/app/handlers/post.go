package handlers

import (
	"encoding/json"
	"github.com/nglmq/calendar/internal/app/domain"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event domain.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if event.UserID == "" {
		http.Error(w, "missing 'user_id' parameter", http.StatusBadRequest)
		return
	}

	err := domain.CreateEvent(event)
	if err != nil {
		http.Error(w, "error creating event", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event created successfully"})
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event domain.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if event.UserID == "" {
		http.Error(w, "missing 'user_id' parameter", http.StatusBadRequest)
		return
	}

	err := domain.UpdateEvent(event)
	if err != nil {
		http.Error(w, "error updating event", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event updated successfully"})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID     string `json:"id"`
		UserID string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if input.UserID == "" {
		http.Error(w, "missing 'user_id' parameter", http.StatusBadRequest)
		return
	}

	err := domain.DeleteEvent(input.ID, input.UserID)
	if err != nil {
		http.Error(w, "error deleting event", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event deleted successfully"})
}
