package main

import (
    "log"

    "github.com/helleonyango/GOJIAS-COLLECTIONS/backend/db"
    "github.com/helleonyango/GOJIAS-COLLECTIONS/backend/models"
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

    tailor := models.NewTailor("Hellen Onyango", "0712345678", "hellen@gojias.com")

    err = db.InsertTailor(database, tailor)
    if err != nil {
        log.Fatal("failed to insert tailor:", err)
    }

    log.Println("Tailor saved successfully:", tailor.Name)

    row := database.QueryRow("SELECT name, phone, email FROM tailors WHERE id = ?", tailor.ID)

    var name, phone, email string
    err = row.Scan(&name, &phone, &email)
    if err != nil {
        log.Fatal("failed to read tailor back:", err)
    }

    log.Println("Read back from database:", name, phone, email)
}