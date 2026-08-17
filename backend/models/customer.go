package models

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
        Name:  name,
        Email: email,
        Phone: phone,
    }
}
