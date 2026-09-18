package models

import "github.com/google/uuid"

func NewTailor(name string, phone string, email string) Tailor {
    return Tailor{
        ID:    uuid.New().String(),
        Name:  name,
        Phone: phone,
        Email: email,
    }
}