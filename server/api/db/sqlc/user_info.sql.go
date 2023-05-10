// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user_info.sql

package db

import (
	"context"
	"database/sql"
)

const createUserInfo = `-- name: CreateUserInfo :one
INSERT INTO "user_info" (
  "id", "user_id", "name", "email", "phone", "faculty_id", "degree", "year_start", "avata_url", "birthday", "bank_account"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id, user_id, name, email, phone, faculty_id, degree, year_start, avata_url, birthday, bank_account
`

type CreateUserInfoParams struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	FacultyID   int32          `json:"faculty_id"`
	Degree      int32          `json:"degree"`
	YearStart   int32          `json:"year_start"`
	AvataUrl    sql.NullString `json:"avata_url"`
	Birthday    sql.NullString `json:"birthday"`
	BankAccount sql.NullString `json:"bank_account"`
}

func (q *Queries) CreateUserInfo(ctx context.Context, arg CreateUserInfoParams) (UserInfo, error) {
	row := q.db.QueryRowContext(ctx, createUserInfo,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.FacultyID,
		arg.Degree,
		arg.YearStart,
		arg.AvataUrl,
		arg.Birthday,
		arg.BankAccount,
	)
	var i UserInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.FacultyID,
		&i.Degree,
		&i.YearStart,
		&i.AvataUrl,
		&i.Birthday,
		&i.BankAccount,
	)
	return i, err
}

const deleteUserInfo = `-- name: DeleteUserInfo :exec
DELETE FROM "user_info"
WHERE id = $1
`

func (q *Queries) DeleteUserInfo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserInfo, id)
	return err
}

const getStudentByName = `-- name: GetStudentByName :many
SELECT "name", user_id from "user_info"
WHERE "name" like $1
`

type GetStudentByNameRow struct {
	Name   string `json:"name"`
	UserID int32  `json:"user_id"`
}

func (q *Queries) GetStudentByName(ctx context.Context, name string) ([]GetStudentByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getStudentByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetStudentByNameRow{}
	for rows.Next() {
		var i GetStudentByNameRow
		if err := rows.Scan(&i.Name, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserInfo = `-- name: GetUserInfo :one
SELECT id, user_id, name, email, phone, faculty_id, degree, year_start, avata_url, birthday, bank_account FROM "user_info"
WHERE "user_id" = $1 LIMIT 1
`

func (q *Queries) GetUserInfo(ctx context.Context, userID int32) (UserInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, userID)
	var i UserInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.FacultyID,
		&i.Degree,
		&i.YearStart,
		&i.AvataUrl,
		&i.Birthday,
		&i.BankAccount,
	)
	return i, err
}

const listUserInfos = `-- name: ListUserInfos :many
SELECT id, user_id, name, email, phone, faculty_id, degree, year_start, avata_url, birthday, bank_account FROM "user_info"
ORDER BY "name"
`

func (q *Queries) ListUserInfos(ctx context.Context) ([]UserInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUserInfos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserInfo{}
	for rows.Next() {
		var i UserInfo
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.FacultyID,
			&i.Degree,
			&i.YearStart,
			&i.AvataUrl,
			&i.Birthday,
			&i.BankAccount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserInfosByName = `-- name: ListUserInfosByName :many
SELECT id, user_id, name, email, phone, faculty_id, degree, year_start, avata_url, birthday, bank_account FROM "user_info"
WHERE name like $1
`

func (q *Queries) ListUserInfosByName(ctx context.Context, name string) ([]UserInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUserInfosByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserInfo{}
	for rows.Next() {
		var i UserInfo
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.FacultyID,
			&i.Degree,
			&i.YearStart,
			&i.AvataUrl,
			&i.Birthday,
			&i.BankAccount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserInfo = `-- name: UpdateUserInfo :exec
UPDATE "user_info"
  set name = $2,
  email = $3,
  phone = $4,
  faculty_id = $5,
  degree = $6,
  year_start = $7,
  avata_url = $8,
  birthday = $9,
  bank_account = $10
WHERE "user_id" = $1
`

type UpdateUserInfoParams struct {
	UserID      int32          `json:"user_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	FacultyID   int32          `json:"faculty_id"`
	Degree      int32          `json:"degree"`
	YearStart   int32          `json:"year_start"`
	AvataUrl    sql.NullString `json:"avata_url"`
	Birthday    sql.NullString `json:"birthday"`
	BankAccount sql.NullString `json:"bank_account"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateUserInfo,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.FacultyID,
		arg.Degree,
		arg.YearStart,
		arg.AvataUrl,
		arg.Birthday,
		arg.BankAccount,
	)
	return err
}
