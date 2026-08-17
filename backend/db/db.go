package db

import (
    "database/sql"

    _ "modernc.org/sqlite"
)

func InitDB(path string) (*sql.DB, error) {
    database, err := sql.Open("sqlite", path)
    if err != nil {
        return nil, err
    }

    err = database.Ping()
    if err != nil {
        return nil, err
    }

    return database, nil
}
func CreateTables(database *sql.DB) error {
    tailorTable := `
    CREATE TABLE IF NOT EXISTS tailors (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        phone TEXT,
        email TEXT,
        specialties TEXT,
        portfolio TEXT,
        rating REAL
    );`

    _, err := database.Exec(tailorTable)
    if err != nil {
        return err
    }

    return nil
}