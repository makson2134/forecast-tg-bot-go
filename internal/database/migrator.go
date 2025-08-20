package database
import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
)
func RunMigrations(db *sql.DB) error {
	content, err := os.ReadFile("migrations/001_initial_schema.up.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}
	if _, err := db.Exec(string(content)); err != nil {
		return fmt.Errorf("failed to apply migration: %w", err)
	}
	slog.Info("migrations applied successfully")
	return nil
}
