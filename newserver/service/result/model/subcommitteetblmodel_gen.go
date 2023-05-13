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
	subcommitteeTblFieldNames          = builder.RawFieldNames(&SubcommitteeTbl{}, true)
	subcommitteeTblRows                = strings.Join(subcommitteeTblFieldNames, ",")
	subcommitteeTblRowsExpectAutoSet   = strings.Join(stringx.Remove(subcommitteeTblFieldNames), ",")
	subcommitteeTblRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(subcommitteeTblFieldNames, "id"))
)

type (
	subcommitteeTblModel interface {
		Insert(ctx context.Context, data *SubcommitteeTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SubcommitteeTbl, error)
		Update(ctx context.Context, data *SubcommitteeTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSubcommitteeTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SubcommitteeTbl struct {
		Id       int64          `db:"id"`
		Name     sql.NullString `db:"name"`
		FacultId int64          `db:"facult_id"`
		EventId  int64          `db:"event_id"`
	}
)

func newSubcommitteeTblModel(conn sqlx.SqlConn) *defaultSubcommitteeTblModel {
	return &defaultSubcommitteeTblModel{
		conn:  conn,
		table: `"public"."subcommittee_tbl"`,
	}
}

func (m *defaultSubcommitteeTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSubcommitteeTblModel) FindOne(ctx context.Context, id int64) (*SubcommitteeTbl, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", subcommitteeTblRows, m.table)
	var resp SubcommitteeTbl
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

func (m *defaultSubcommitteeTblModel) Insert(ctx context.Context, data *SubcommitteeTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, subcommitteeTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.FacultId, data.EventId)
	return ret, err
}

func (m *defaultSubcommitteeTblModel) Update(ctx context.Context, data *SubcommitteeTbl) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, subcommitteeTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.FacultId, data.EventId)
	return err
}

func (m *defaultSubcommitteeTblModel) tableName() string {
	return m.table
}
