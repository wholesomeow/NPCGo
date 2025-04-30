package npcapi

import (
	"net/http"
	"time"
)

type Response struct {
	Status    string
	Message   string
	Data      interface{}
	Error     interface{}
	Timestamp time.Time
}

func Response500(msg string) (int, Response) {
	// Return error message if querystring parameters don't pass
	status := http.StatusInternalServerError
	response := Response{
		Status:    http.StatusText(status),
		Message:   msg,
		Timestamp: time.Now(),
	}

	return status, response
}
