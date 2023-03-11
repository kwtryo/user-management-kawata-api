package testutil

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

func OpenDbForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	cfg := CreateConfigForTest(t)
	driver := "postgres"
	db, err := sql.Open(driver, fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	))
	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, driver)
}
