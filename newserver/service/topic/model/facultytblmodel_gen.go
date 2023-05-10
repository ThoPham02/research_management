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
	facultyTblFieldNames          = builder.RawFieldNames(&FacultyTbl{}, true)
	facultyTblRows                = strings.Join(facultyTblFieldNames, ",")
	facultyTblRowsExpectAutoSet   = strings.Join(stringx.Remove(facultyTblFieldNames), ",")
	facultyTblRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(facultyTblFieldNames, "id"))
)

type (
	facultyTblModel interface {
		Insert(ctx context.Context, data *FacultyTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FacultyTbl, error)
		Update(ctx context.Context, data *FacultyTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFacultyTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FacultyTbl struct {
		Id   int64  `db:"id"`
		Name string `db:"name"`
	}
)

func newFacultyTblModel(conn sqlx.SqlConn) *defaultFacultyTblModel {
	return &defaultFacultyTblModel{
		conn:  conn,
		table: `"public"."faculty_tbl"`,
	}
}

func (m *defaultFacultyTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFacultyTblModel) FindOne(ctx context.Context, id int64) (*FacultyTbl, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", facultyTblRows, m.table)
	var resp FacultyTbl
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

func (m *defaultFacultyTblModel) Insert(ctx context.Context, data *FacultyTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2)", m.table, facultyTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name)
	return ret, err
}

func (m *defaultFacultyTblModel) Update(ctx context.Context, data *FacultyTbl) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, facultyTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name)
	return err
}

func (m *defaultFacultyTblModel) tableName() string {
	return m.table
}
