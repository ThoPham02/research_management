// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: student.sql

package db

import (
	"context"
)

const createStudent = `-- name: CreateStudent :one
INSERT INTO
    students (name, user_id)
VALUES
    ($1, $2) RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

type CreateStudentParams struct {
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent, arg.Name, arg.UserID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteStudent = `-- name: DeleteStudent :one
UPDATE
    students
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteStudent(ctx context.Context, id int64) (Student, error) {
	row := q.db.QueryRowContext(ctx, deleteStudent, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getStudent = `-- name: GetStudent :one
SELECT
    id, name, user_id, created_at, updated_at, deleted_at
FROM
    students
WHERE
    id = $1
`

func (q *Queries) GetStudent(ctx context.Context, id int64) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listStudents = `-- name: ListStudents :many
SELECT
    id, name, user_id, created_at, updated_at, deleted_at
FROM
    students
ORDER BY
    id
LIMIT $1 OFFSET $2
`

type ListStudentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, listStudents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Student{}
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const updateStudent = `-- name: UpdateStudent :one
UPDATE
    students
SET
    name = $2,
    updated_at = NOW()
WHERE
    id = $1 RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

type UpdateStudentParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateStudent, arg.ID, arg.Name)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
