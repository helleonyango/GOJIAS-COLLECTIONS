package models

import "time"

type Order struct {
    ID              string
    CustomerID      string
    TailorID        string
    MeasurementID   string
    FabricChoice    string
    StyleDescription string
    Status          string
    Price           float64
    DepositPaid     bool
    BalancePaid     bool
    PlacedAt        time.Time
    ExpectedReady   time.Time
}
func NewOrder(customerID string, tailorID string, measurementID string, fabricChoice string, styleDescription string) Order {
    return Order{
        CustomerID:       customerID,
        TailorID:         tailorID,
        MeasurementID:    measurementID,
        FabricChoice:     fabricChoice,
        StyleDescription: styleDescription,
        Status:           "requested",
        PlacedAt:         time.Now(),
    }
}