package db

import (
    "database/sql"
    "strings"
    "time"

    _ "modernc.org/sqlite"

    "github.com/helleonyango/GOJIAS-COLLECTIONS/backend/models"
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

    customerTable := `
    CREATE TABLE IF NOT EXISTS customers (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT,
        phone TEXT
    );`

    _, err = database.Exec(customerTable)
    if err != nil {
        return err
    }

    measurementTable := `
    CREATE TABLE IF NOT EXISTS measurement_profiles (
        id TEXT PRIMARY KEY,
        customer_id TEXT,
        label TEXT,
        bust REAL,
        waist REAL,
        hips REAL,
        shoulder_width REAL,
        sleeve_length REAL,
        outfit_length REAL,
        last_updated TEXT
    );`

    _, err = database.Exec(measurementTable)
    if err != nil {
        return err
    }

    orderTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id TEXT PRIMARY KEY,
        customer_id TEXT,
        tailor_id TEXT,
        measurement_id TEXT,
        fabric_choice TEXT,
        style_description TEXT,
        status TEXT,
        price REAL,
        deposit_paid INTEGER,
        balance_paid INTEGER,
        placed_at TEXT,
        expected_ready TEXT
    );`

    _, err = database.Exec(orderTable)
    if err != nil {
        return err
    }

    return nil
}

func InsertTailor(database *sql.DB, tailor models.Tailor) error {
    query := `
    INSERT INTO tailors (id, name, phone, email, specialties, portfolio, rating)
    VALUES (?, ?, ?, ?, ?, ?, ?);`

    _, err := database.Exec(query,
        tailor.ID,
        tailor.Name,
        tailor.Phone,
        tailor.Email,
        strings.Join(tailor.Specialties, ","),
        strings.Join(tailor.Portfolio, ","),
        tailor.Rating,
    )

    return err
}
func InsertCustomer(database *sql.DB, customer models.Customer) error {
    query := `
    INSERT INTO customers (id, name, email, phone)
    VALUES (?, ?, ?, ?);`

    _, err := database.Exec(query,
        customer.ID,
        customer.Name,
        customer.Email,
        customer.Phone,
    )

    return err
}
func InsertMeasurementProfile(database *sql.DB, m models.MeasurementProfile) error {
    query := `
    INSERT INTO measurement_profiles (id, customer_id, label, bust, waist, hips, shoulder_width, sleeve_length, outfit_length, last_updated)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

    _, err := database.Exec(query,
        m.ID,
        m.CustomerID,
        m.Label,
        m.Bust,
        m.Waist,
        m.Hips,
        m.ShoulderWidth,
        m.SleeveLength,
        m.OutfitLength,
        m.LastUpdated.Format(time.RFC3339),
    )

    return err
}

func InsertOrder(database *sql.DB, order models.Order) error {
    query := `
    INSERT INTO orders (id, customer_id, tailor_id, measurement_id, fabric_choice, style_description, status, price, deposit_paid, balance_paid, placed_at, expected_ready)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

    _, err := database.Exec(query,
        order.ID,
        order.CustomerID,
        order.TailorID,
        order.MeasurementID,
        order.FabricChoice,
        order.StyleDescription,
        order.Status,
        order.Price,
        order.DepositPaid,
        order.BalancePaid,
        order.PlacedAt.Format(time.RFC3339),
        order.ExpectedReady.Format(time.RFC3339),
    )

    return err
}