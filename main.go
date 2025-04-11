package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/papaya147/expenses/api"
	"github.com/papaya147/expenses/db/sqlc"
)

func main() {
	conn, err := migrateUp("db", "file://db/migration")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := sqlc.New(conn)

	s := api.NewServer(queries)

	port := 8080

	fmt.Printf("starting server on port %d \n", port)
	err = s.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}

func dbFileName(dbName string) string {
	return fmt.Sprintf("data/%s.sqlite", dbName)
}

func migrateDown(dbName string) {
	_ = os.Remove(dbFileName(dbName))
}

func migrateUp(dbName, migrationSource string) (*sql.DB, error) {
	path := dbFileName(dbName)

	parentPath := filepath.Join(path, "..")
	_ = os.MkdirAll(parentPath, os.ModePerm)

	conn, err := sql.Open("sqlite3", dbFileName(dbName))
	if err != nil {
		return nil, err
	}

	dbSource := fmt.Sprintf("sqlite3://%s", dbFileName(dbName))

	migration, err := migrate.New(migrationSource, dbSource)
	if err != nil {
		return nil, err
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return conn, nil
}

func browserOpen(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
