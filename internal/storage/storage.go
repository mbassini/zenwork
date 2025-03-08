package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName            = "zenwork.db"
	dbFilePermissions = 0755 // RWX R-X R-X
)

var db *sql.DB

func InitDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	fmt.Printf("[Home Dir]: %v\n", homeDir)

	dbPath := filepath.Join(homeDir, ".zenwork", dbName)
	fmt.Printf("[DB Path]: %v\n", dbPath)

	err = os.MkdirAll(filepath.Dir(dbPath), dbFilePermissions)
	if err != nil {
		return fmt.Errorf("error creating database file: %w", err)
	}

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}

	if err = runMigrations(); err != nil {
		return fmt.Errorf("migrations failed: %w", err)
	}

	return nil
}

func runMigrations() error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations(
			version INTEGER,
			applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
		`); err != nil {
		return err
	}

	migrationFiles := []struct {
		version  int
		filename string
	}{
		{1, "001_create_tasks_table.sql"},
	}

	for _, mf := range migrationFiles {
		content, err := os.ReadFile(filepath.Join("internal", "storage", "migrations", mf.filename))
		if err != nil {
			return fmt.Errorf("error reading migration %d: %w", mf.version, err)
		}

		tx, err := db.Begin()
		if err != nil {
			return err
		}

		if _, err = tx.Exec(string(content)); err != nil {
			tx.Rollback()
			return fmt.Errorf("error aplying migration %d: %w", mf.version, err)
		}

		if _, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", mf.version); err != nil {
			tx.Rollback()
			return err
		}

		if err = tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}
