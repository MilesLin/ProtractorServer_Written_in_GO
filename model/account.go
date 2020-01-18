package model

type Account struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

var Accounts []Account = []Account{
	Account{
		Id:        1,
		FirstName: "John J.",
		LastName:  "John",
		Username:  "John",
		Password:  "123456",
	},
}
