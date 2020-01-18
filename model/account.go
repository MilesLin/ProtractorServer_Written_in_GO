package model

type Account struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
	Password  string
}


var Accounts []Account = []Account{
	Account{
		Id:1,
		FirstName:"John J.",
		LastName:"John",
		Username:"John",
		Password:"123456",
	},
}