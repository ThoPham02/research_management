// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user_info.sql

package db

import (
	"context"
	"database/sql"
)

const createUserInfo = `-- name: CreateUserInfo :exec
INSERT INTO "user_info" (
  user_id, name, description, avata_url, birthday, faculity_id, year_start, bank_account, phone, sex
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
`

type CreateUserInfoParams struct {
	UserID      int64          `json:"user_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	AvataUrl    sql.NullString `json:"avata_url"`
	Birthday    sql.NullTime   `json:"birthday"`
	FaculityID  int64          `json:"faculity_id"`
	YearStart   int64          `json:"year_start"`
	BankAccount sql.NullString `json:"bank_account"`
	Phone       sql.NullString `json:"phone"`
	Sex         sql.NullInt64  `json:"sex"`
}

func (q *Queries) CreateUserInfo(ctx context.Context, arg CreateUserInfoParams) error {
	_, err := q.db.ExecContext(ctx, createUserInfo,
		arg.UserID,
		arg.Name,
		arg.Description,
		arg.AvataUrl,
		arg.Birthday,
		arg.FaculityID,
		arg.YearStart,
		arg.BankAccount,
		arg.Phone,
		arg.Sex,
	)
	return err
}

const getUserInfo = `-- name: GetUserInfo :one
SELECT user_id, name, description, avata_url, birthday, faculity_id, year_start, bank_account, phone, sex FROM "user_info"
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserInfo(ctx context.Context, userID int64) (UserInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, userID)
	var i UserInfo
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.AvataUrl,
		&i.Birthday,
		&i.FaculityID,
		&i.YearStart,
		&i.BankAccount,
		&i.Phone,
		&i.Sex,
	)
	return i, err
}

const updateUserInfo = `-- name: UpdateUserInfo :exec
UPDATE "user_info"
  set name = $2,
  description = $3,
  avata_url = $4,
  birthday = $5,
  faculity_id = $6,
  year_start = $7,
  bank_account = $8,
  phone = $9,
  sex = $10
WHERE user_id = $1
`

type UpdateUserInfoParams struct {
	UserID      int64          `json:"user_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	AvataUrl    sql.NullString `json:"avata_url"`
	Birthday    sql.NullTime   `json:"birthday"`
	FaculityID  int64          `json:"faculity_id"`
	YearStart   int64          `json:"year_start"`
	BankAccount sql.NullString `json:"bank_account"`
	Phone       sql.NullString `json:"phone"`
	Sex         sql.NullInt64  `json:"sex"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateUserInfo,
		arg.UserID,
		arg.Name,
		arg.Description,
		arg.AvataUrl,
		arg.Birthday,
		arg.FaculityID,
		arg.YearStart,
		arg.BankAccount,
		arg.Phone,
		arg.Sex,
	)
	return err
}