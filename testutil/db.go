package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 3306
	if _, difined := os.LookupEnv("CI"); difined {
		port = 3306
	}

	db, err := sql.Open("mysql", fmt.Sprintf("lab:lab@tcp(127.0.0.1:%d)/lab?parseTime=true", port))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "mysql")
}
