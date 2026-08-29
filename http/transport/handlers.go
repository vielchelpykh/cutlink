package transport

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HandlerList struct {
	AllLinks *ListLinks
}

func NewHandlerList(list *ListLinks) *HandlerList {
	return &HandlerList{
		AllLinks: list,
	}
}

func CreateHTTPError(message string, w http.ResponseWriter, statusCode int) {
	errDTO := ErrorDTO{
		Message: "your link just created",
		Time:    time.Now(),
	}
	b, _ := json.MarshalIndent(errDTO, "", "    ")
	http.Error(w, string(b), statusCode)
}

func (handlerList *HandlerList) GenerateShortLink(newFullLink LinkDTO) string {
	elements := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*?+=-{}()"
	for {
		newLen := rand.Intn(len(newFullLink.FullLink) / 2)
		var shortLink string
		for i := 0; i < newLen; i++ {
			shortLink += string(elements[rand.Intn(len(elements))])
		}
		if _, ok := handlerList.AllLinks.List[shortLink]; !ok {
			return shortLink
		}
	}
}

func (handlerList *HandlerList) HandlerCreateShortLink(w http.ResponseWriter, r *http.Request) {
	var newFullLink LinkDTO
	if err := json.NewDecoder(r.Body).Decode(&newFullLink); err != nil {
		CreateHTTPError(err.Error(), w, http.StatusBadRequest)
		return
	}

	if handlerList.AllLinks.ValidateForCreate(newFullLink.FullLink) {
		CreateHTTPError("your link just created", w, http.StatusConflict)
		return
	}

	newShortLink := handlerList.GenerateShortLink(newFullLink)
	newLinkInfo := NewLink(newFullLink.FullLink, newShortLink)
	handlerList.AllLinks.List[newShortLink] = newLinkInfo

	b, _ := json.MarshalIndent(newLinkInfo, "", "    ")
	w.Write(b)
	w.WriteHeader(http.StatusCreated)
}

func (handlerList *HandlerList) HandlerFollowLink(w http.ResponseWriter, r *http.Request) {
	title, _ := mux.Vars(r)["title"]
	linkInfo, ok := handlerList.AllLinks.List[title]
	if !ok {
		CreateHTTPError("link not found", w, http.StatusNoContent)
		return
	}
	linkInfo.Pressed++
	handlerList.AllLinks.List[title] = linkInfo
	w.Header().Set("Location", linkInfo.FullLink)
	w.WriteHeader(http.StatusFound)
	w.Write([]byte{})
}

func (handlerList *HandlerList) HandlerGetStatistic(w http.ResponseWriter, r *http.Request) {
	title, _ := mux.Vars(r)["title"]
	linkInfo, ok := handlerList.AllLinks.List[title]
	if !ok {
		CreateHTTPError("link not found", w, http.StatusNoContent)
		return
	}

	b, _ := json.MarshalIndent(linkInfo, "", "    ")
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
