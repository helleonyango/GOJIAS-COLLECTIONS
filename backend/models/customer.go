package models

type Customer struct {
    ID                  string
    Name                string
    Email               string
    Phone               string
    MeasurementProfiles []MeasurementProfile
    OrderHistory        []Order
}