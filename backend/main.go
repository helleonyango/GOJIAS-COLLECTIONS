package main

import (
    "log"

    "github.com/helleonyango/GOJIAS-COLLECTIONS/backend/db"
)

func main() {
    database, err := db.InitDB("gojias.db")
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }
    defer database.Close()

    err = db.CreateTables(database)
    if err != nil {
        log.Fatal("failed to create tables:", err)
    }

    log.Println("Database connected and tables ready.")
}