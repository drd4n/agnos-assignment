package database

import (
	"embed"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func RunSQLMigrations(db *gorm.DB) error {
	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			migrationSQL, err := migrationFiles.ReadFile("migrations/" + entry.Name())
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
			}

			if err := db.Exec(string(migrationSQL)).Error; err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", entry.Name(), err)
			}

			fmt.Printf("✓ Migration executed: %s\n", entry.Name())
		}
	}

	return nil
}
