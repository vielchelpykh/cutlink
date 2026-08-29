package transport

import (
	"sync"
	"time"
)

type LinkDTO struct {
	FullLink string `json:"full-link"`
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type Link struct {
	FullLink  string
	ShortLink string
	CreatedAt time.Time
	Pressed   int
}

func NewLink(FullLink string, ShortLink string) Link {
	return Link{
		FullLink:  FullLink,
		ShortLink: ShortLink,
		CreatedAt: time.Now(),
		Pressed:   0,
	}
}

type ListLinks struct {
	List map[string]Link
	Mtx  sync.Mutex
}

func NewList() *ListLinks {
	return &ListLinks{
		List: make(map[string]Link),
	}
}

func (links *ListLinks) AllInfo() map[string]Link {
	allLinks := make(map[string]Link, len(links.List))
	for key, value := range links.List {
		allLinks[key] = value
	}
	return allLinks
}

func (links *ListLinks) ValidateForCreate(link string) bool {
	for _, value := range links.List {
		if value.FullLink == link {
			return true
		}
	}
	return false
}
