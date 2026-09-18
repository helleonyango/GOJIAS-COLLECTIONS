package models

import (
    "time"

    "github.com/google/uuid"
)

type MeasurementProfile struct {
    ID            string
    CustomerID    string
    Label         string
    Bust          float64
    Waist         float64
    Hips          float64
    ShoulderWidth float64
    SleeveLength  float64
    OutfitLength  float64
    LastUpdated   time.Time
}

func NewMeasurementProfile(customerID string, label string) MeasurementProfile {
    return MeasurementProfile{
        ID:         uuid.New().String(),
        CustomerID: customerID,
        Label:      label,
    }
}