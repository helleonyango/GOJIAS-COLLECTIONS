package models

type Tailor struct {
    Name        string
    Phone       string
    Email       string
    Specialties []string
    Portfolio   []string
    Rating      float64
}
func NewTailor(name string, phone string, email string) Tailor {
    return Tailor{
        Name:  name,
        Phone: phone,
        Email: email,
    }
}