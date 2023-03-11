package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kwtryo/user-management-kawata-api/clock"
	"github.com/kwtryo/user-management-kawata-api/config"
	_ "github.com/lib/pq"
)

// const (
// 	// MySQLの重複エラーコード
// 	ErrCodeMySQLDuplicateEntry = 1062
// )

// var (
// 	ErrNotFound     = errors.New("not found")
// 	ErrAlreadyEntry = errors.New("duplicate entry")
// )

type Repository struct {
	Clocker clock.Clocker
}

func New(cfg *config.Config) (*sqlx.DB, func(), error) {
	driver := "postgres"
	db, err := sql.Open(driver, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	))
	if err != nil {
		return nil, nil, err
	}

	// sql.Openは接続確認が行われないため、ここで確認する。
	if err := db.Ping(); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db, driver)
	return xdb, func() { _ = db.Close() }, nil
}

type DBConnection interface {
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

var (
	_ DBConnection = (*sqlx.DB)(nil)
	_ DBConnection = (*sqlx.Tx)(nil)
)
