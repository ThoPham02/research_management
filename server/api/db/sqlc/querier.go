// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetUserByName(ctx context.Context, name string) (User, error)
	GetUserInfo(ctx context.Context, userID int64) (UserInfo, error)
}

var _ Querier = (*Queries)(nil)
