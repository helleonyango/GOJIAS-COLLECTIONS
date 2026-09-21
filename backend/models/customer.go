package models

import "github.com/google/uuid"

type Customer struct {
    ID                  string
    Name                string
    Email               string
    Phone               string
    MeasurementProfiles []MeasurementProfile
    OrderHistory        []Order
}

func NewCustomer(name string, email string, phone string) Customer {
    return Customer{
        ID:    uuid.New().String(),
        Name:  name,
        Email: email,
        Phone: phone,
    }
}
