package models

type Customers struct {
	CustomerId int     `json:"customerid" bson:"customerid"`
	FirstName  string  `json:"firstname" bson:"firstname"`
	LastName   string  `json:"lastname" bson:"lastname"`
	BankId     string  `json:"bankid" bson:"bankid"`
	Balance    float64 `json:"balance" bson:"balance"`
	CreatedAt  string  `json:"createdat" bson:"createdat"`
	UpdatedAt  string  `json:"updatedat" bson:"updatedat"`
	IsActive   string  `json:"isactive" bson:"isactive"`
}
type CustomerResponse struct {
	CustomerId int    `json:"customerid" bson:"customerid"`
	CreatedAt  string `json:"isactive" bson:"isactive"`
}
