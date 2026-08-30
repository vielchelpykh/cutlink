package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type LinkModel struct {
	FullLink  string
	ShortLink string
	CreatedAt time.Time
	Pressed   int
}

func NewLink(FullLink string, ShortLink string) LinkModel {
	return LinkModel{
		FullLink:  FullLink,
		ShortLink: ShortLink,
		CreatedAt: time.Now(),
		Pressed:   0,
	}
}

type RepositoryModel struct {
	Ctx  context.Context
	Conn *pgx.Conn
}

func NewRepository(ctx context.Context, conn *pgx.Conn) *RepositoryModel {
	return &RepositoryModel{
		Ctx:  ctx,
		Conn: conn,
	}
}

func (repo *RepositoryModel) FindLink(link string, isFull bool) bool {
	var sqlQuery string

	if isFull {
		sqlQuery = `
	SELECT COUNT(*) FROM links
	WHERE full_link=$1;
	`
	} else {
		sqlQuery = `
	SELECT COUNT(*) FROM links
	WHERE short_link=$1;
	`
	}

	var count int
	row := repo.Conn.QueryRow(repo.Ctx, sqlQuery, link)

	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return false
	}
	return true
}

func (repo *RepositoryModel) InsertLink(link LinkModel) {
	sqlQuery := `
	INSERT INTO links (full_link, short_link, created_at, pressed)
	VALUES($1, $2, $3, $4);
	`

	if _, err := repo.Conn.Exec(
		repo.Ctx,
		sqlQuery,
		link.FullLink,
		link.ShortLink,
		link.CreatedAt,
		link.Pressed,
	); err != nil {
		panic(err)
	}
}

func (repo *RepositoryModel) IncreasePressedLink(short_link string) {
	sqlQuery := `
	UPDATE links
	SET pressed = pressed + 1
	WHERE short_link = $1;
	`
	if _, err := repo.Conn.Exec(repo.Ctx, sqlQuery, short_link); err != nil {
		panic(err)
	}
}

func (repo *RepositoryModel) GetInfoLink(shortLink string) *LinkModel {
	sqlQuery := `
	SELECT full_link, short_link, created_at, pressed
	FROM links
	WHERE short_link = $1;
	`

	row, err := repo.Conn.Query(repo.Ctx, sqlQuery, shortLink)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	var linkInfo LinkModel
	if err := row.Scan(
		&linkInfo.FullLink,
		&linkInfo.ShortLink,
		&linkInfo.CreatedAt,
		&linkInfo.Pressed); err != nil {
		panic(err)
	}
	return &linkInfo
}

func (repo *RepositoryModel) GetAllInfo() []LinkModel {
	sqlQuery := `
	SELECT full_link, short_link, created_at, pressed
	FROM links
	ORDER by id ASC
	`
	rows, err := repo.Conn.Query(repo.Ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	linksList := make([]LinkModel, 0)
	for rows.Next() {
		var link LinkModel
		if err := rows.Scan(
			&link.FullLink,
			&link.ShortLink,
			&link.CreatedAt,
			&link.Pressed); err != nil {
			panic(err)
		}
		linksList = append(linksList, link)
	}
	return linksList
}
