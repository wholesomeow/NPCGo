package npcapi

import "time"

type Response struct {
	Status    string
	Message   string
	Data      interface{}
	Error     interface{}
	Timestamp time.Time
}
