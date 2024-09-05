package controllers

import (
	"encoding/json"
	"coffee-project/models"
	"coffee-project/views"
	"net/http"
	"strconv"
)

type EventController struct {
	EventManager models.EventManagerInterface
	View         *views.EventView
}

func NewEventController(eventManager models.EventManagerInterface) *EventController {
	return &EventController{
		EventManager: eventManager,
		View:         &views.EventView{},
	}
}

func (c *EventController) ListEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		c.View.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	events := c.EventManager.ListEvents()
	c.View.RenderJson(w, events)
}

func (c *EventController) ViewEventDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		c.View.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	eventIndexStr := r.URL.Query().Get("eventIndex")
	eventIndex, err := strconv.Atoi(eventIndexStr)
	if err != nil {
		c.View.RenderError(w, http.StatusBadRequest, "Invalid event index")
		return
	}

	event, err := c.EventManager.GetEvent(eventIndex)
	if err != nil {
		c.View.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	c.View.RenderJson(w, event)
}

func (c *EventController) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		c.View.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.View.RenderError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	eventIndexStr := r.URL.Query().Get("eventIndex")
	eventIndex, err := strconv.Atoi(eventIndexStr)
	if err != nil {
		c.View.RenderError(w, http.StatusBadRequest, "Invalid event index")
		return
	}

	err = c.EventManager.RegisterUser(eventIndex, user)
	if err != nil {
		c.View.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	c.View.RenderMessage(w, "User registered successfully!")
}
