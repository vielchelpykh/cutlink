package transport

import (
	"cutlink/http/repository"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HandlerList struct {
	Repo *repository.RepositoryModel
}

func NewHandlerList(repo *repository.RepositoryModel) *HandlerList {
	return &HandlerList{
		Repo: repo,
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

func GenerateShortLink(repo *repository.RepositoryModel, newFullLink string) string {
	elements := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*?+=-{}()"
	for {
		newLen := rand.Intn(len(newFullLink) / 2)
		var shortLink string
		for i := 0; i < newLen; i++ {
			shortLink += string(elements[rand.Intn(len(elements))])
		}

		sqlQuery := `
		SELECT COUNT(*) FROM links
		WHERE short_link=$1;
		`
		row := repo.Conn.QueryRow(repo.Ctx, sqlQuery, shortLink)
		var count int
		err := row.Scan(&count)
		if err != nil {
			panic(err)
		}
		if count == 0 {
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

	if handlerList.Repo.FindLink(newFullLink.FullLink, true) {
		CreateHTTPError("your link just created", w, http.StatusConflict)
		return
	}

	newShortLink := GenerateShortLink(handlerList.Repo, newFullLink.FullLink)
	newLinkInfo := repository.NewLink(newFullLink.FullLink, newShortLink)
	handlerList.Repo.InsertLink(newLinkInfo)

	b, _ := json.MarshalIndent(newLinkInfo, "", "    ")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (handlerList *HandlerList) HandlerFollowLink(w http.ResponseWriter, r *http.Request) {
	shortLink, _ := mux.Vars(r)["title"]

	if !handlerList.Repo.FindLink(shortLink, false) {
		CreateHTTPError("link not found", w, http.StatusNoContent)
		return
	}

	handlerList.Repo.IncreasePressedLink(shortLink)
	linkInfo := handlerList.Repo.GetInfoLink(shortLink)

	w.Header().Set("Location", linkInfo.FullLink)
	w.WriteHeader(http.StatusFound)
	w.Write([]byte{})
}

func (handlerList *HandlerList) HandlerGetStatistic(w http.ResponseWriter, r *http.Request) {
	shortLink, _ := mux.Vars(r)["title"]
	if !handlerList.Repo.FindLink(shortLink, false) {
		CreateHTTPError("link not found", w, http.StatusNoContent)
		return
	}

	linkInfo := handlerList.Repo.GetInfoLink(shortLink)

	b, _ := json.MarshalIndent(linkInfo, "", "    ")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (handlerList *HandlerList) HandlerGetAllInfo(w http.ResponseWriter, r *http.Request) {
	list := handlerList.Repo.GetAllInfo()
	b, _ := json.MarshalIndent(list, "", "    ")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
