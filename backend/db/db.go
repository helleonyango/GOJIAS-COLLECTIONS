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
