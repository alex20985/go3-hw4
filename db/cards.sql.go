// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: cards.sql

package db

import (
	"context"
)

const createCard = `-- name: CreateCard :one
INSERT INTO cards (
	time, 
	date, 
	title, 
	client,
	"user" -- user is a reserved keyword, must use ""
) VALUES (
    $1, 
		$2, 
		$3, 
		$4, 
		$5
) RETURNING id, time, date, title, client, "user"
`

type CreateCardParams struct {
	Time   string
	Date   string
	Title  string
	Client string
	User   string
}

func (q *Queries) CreateCard(ctx context.Context, arg CreateCardParams) (Card, error) {
	row := q.db.QueryRow(ctx, createCard,
		arg.Time,
		arg.Date,
		arg.Title,
		arg.Client,
		arg.User,
	)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Time,
		&i.Date,
		&i.Title,
		&i.Client,
		&i.User,
	)
	return i, err
}

const deleteCard = `-- name: DeleteCard :exec
DELETE FROM cards
WHERE id = $1
`

func (q *Queries) DeleteCard(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCard, id)
	return err
}

const getCard = `-- name: GetCard :one
SELECT id, time, date, title, client, "user" FROM cards
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCard(ctx context.Context, id int64) (Card, error) {
	row := q.db.QueryRow(ctx, getCard, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Time,
		&i.Date,
		&i.Title,
		&i.Client,
		&i.User,
	)
	return i, err
}

const listCards = `-- name: ListCards :many
SELECT id, time, date, title, client, "user" FROM cards
WHERE 
date = $1 OR
date = $2 OR
date = $3 OR
date = $4 OR
date = $5 OR
date = $6 OR
date = $7
ORDER BY date
`

type ListCardsParams struct {
	Date1 string
	Date2 string
	Date3 string
	Date4 string
	Date5 string
	Date6 string
	Date7 string
}

func (q *Queries) ListCards(ctx context.Context, arg ListCardsParams) ([]Card, error) {
	rows, err := q.db.Query(ctx, listCards,
		arg.Date1,
		arg.Date2,
		arg.Date3,
		arg.Date4,
		arg.Date5,
		arg.Date6,
		arg.Date7,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Time,
			&i.Date,
			&i.Title,
			&i.Client,
			&i.User,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCards = `-- name: UpdateCards :one
UPDATE cards
SET 
    time = $2,
    date = $3,
    title = $4,
    client = $5,
		"user" = $6
WHERE id = $1
RETURNING id, time, date, title, client, "user"
`

type UpdateCardsParams struct {
	ID     int64
	Time   string
	Date   string
	Title  string
	Client string
	User   string
}

func (q *Queries) UpdateCards(ctx context.Context, arg UpdateCardsParams) (Card, error) {
	row := q.db.QueryRow(ctx, updateCards,
		arg.ID,
		arg.Time,
		arg.Date,
		arg.Title,
		arg.Client,
		arg.User,
	)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Time,
		&i.Date,
		&i.Title,
		&i.Client,
		&i.User,
	)
	return i, err
}
