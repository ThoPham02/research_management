// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: department.sql

package db

import (
	"context"
)

const createDepartment = `-- name: CreateDepartment :one
INSERT INTO
    departments (name, faculty_id, user_id)
VALUES
    ($1, $2, $3) RETURNING id, name, faculty_id, user_id, created_at, updated_at, deleted_at
`

type CreateDepartmentParams struct {
	Name      string `json:"name"`
	FacultyID int64  `json:"faculty_id"`
	UserID    int64  `json:"user_id"`
}

func (q *Queries) CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (Department, error) {
	row := q.db.QueryRowContext(ctx, createDepartment, arg.Name, arg.FacultyID, arg.UserID)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FacultyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteDepartment = `-- name: DeleteDepartment :one
UPDATE
    departments
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING id, name, faculty_id, user_id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteDepartment(ctx context.Context, id int64) (Department, error) {
	row := q.db.QueryRowContext(ctx, deleteDepartment, id)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FacultyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDepartment = `-- name: GetDepartment :one
SELECT
    id, name, faculty_id, user_id, created_at, updated_at, deleted_at
FROM
    departments
WHERE
    user_id = $1
`

func (q *Queries) GetDepartment(ctx context.Context, userID int64) (Department, error) {
	row := q.db.QueryRowContext(ctx, getDepartment, userID)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FacultyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listDepartments = `-- name: ListDepartments :many
SELECT
    id, name, faculty_id, user_id, created_at, updated_at, deleted_at
FROM
    departments
WHERE faculty_id = $1
`

func (q *Queries) ListDepartments(ctx context.Context, facultyID int64) ([]Department, error) {
	rows, err := q.db.QueryContext(ctx, listDepartments, facultyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Department{}
	for rows.Next() {
		var i Department
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.FacultyID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const updateDepartment = `-- name: UpdateDepartment :one
UPDATE
    departments
SET
    name = $2,
    faculty_id = $3,
    updated_at = NOW()
WHERE
    id = $1 RETURNING id, name, faculty_id, user_id, created_at, updated_at, deleted_at
`

type UpdateDepartmentParams struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FacultyID int64  `json:"faculty_id"`
}

func (q *Queries) UpdateDepartment(ctx context.Context, arg UpdateDepartmentParams) (Department, error) {
	row := q.db.QueryRowContext(ctx, updateDepartment, arg.ID, arg.Name, arg.FacultyID)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FacultyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
