package models

import "github.com/google/uuid"

type Tailor struct {
    ID          string
    Name        string
    Phone       string
    Email       string
    Specialties []string
    Portfolio   []string
    Rating      float64
}

func NewTailor(name string, phone string, email string) Tailor {
    return Tailor{
        ID:    uuid.New().String(),
        Name:  name,
        Phone: phone,
        Email: email,
    }
}