// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	notificationTblFieldNames          = builder.RawFieldNames(&NotificationTbl{}, true)
	notificationTblRows                = strings.Join(notificationTblFieldNames, ",")
	notificationTblRowsExpectAutoSet   = strings.Join(stringx.Remove(notificationTblFieldNames), ",")
	notificationTblRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(notificationTblFieldNames, "id"))
)

type (
	notificationTblModel interface {
		Insert(ctx context.Context, data *NotificationTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*NotificationTbl, error)
		Update(ctx context.Context, data *NotificationTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultNotificationTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	NotificationTbl struct {
		Id   int64          `db:"id"`
		Name string         `db:"name"`
		Url  sql.NullString `db:"url"`
	}
)

func newNotificationTblModel(conn sqlx.SqlConn) *defaultNotificationTblModel {
	return &defaultNotificationTblModel{
		conn:  conn,
		table: `"public"."notification_tbl"`,
	}
}

func (m *defaultNotificationTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultNotificationTblModel) FindOne(ctx context.Context, id int64) (*NotificationTbl, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", notificationTblRows, m.table)
	var resp NotificationTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNotificationTblModel) Insert(ctx context.Context, data *NotificationTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, notificationTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Url)
	return ret, err
}

func (m *defaultNotificationTblModel) Update(ctx context.Context, data *NotificationTbl) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, notificationTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Url)
	return err
}

func (m *defaultNotificationTblModel) tableName() string {
	return m.table
}
