package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Done      bool      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
