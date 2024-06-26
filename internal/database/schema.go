package database

import (
	"database/sql"
	"time"
)

type File struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Type      string    `json:"type"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}

type FileToAdd struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

func InitSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS files (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name VARCHAR(255) UNIQUE NOT NULL,
	    path TEXT NOT NULL,
	    type TEXT NOT NULL,
	    size INTEGER NOT NULL,
	    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(schema)
	return err
}
