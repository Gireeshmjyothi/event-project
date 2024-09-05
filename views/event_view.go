package views

import (
	"encoding/json"
	"net/http"
)

type EventView struct{}

func (v *EventView) RenderJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (v *EventView) RenderError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	v.RenderJson(w, map[string]string{"error": message})
}

func (v *EventView) RenderMessage(w http.ResponseWriter, message string) {
	v.RenderJson(w, map[string]string{"message": message})
}
