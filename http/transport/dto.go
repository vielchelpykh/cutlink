package transport

import (
	"time"
)

type LinkDTO struct {
	FullLink string `json:"full-link"`
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
