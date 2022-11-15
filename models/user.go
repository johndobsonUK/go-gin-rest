package models

type Address struct {
	State 		string `json:"state" bson:"state"`
	City  		string `json:"city" bson:"city"`
	Street  	string `json:"street" bson:"street"`
	PostCode  	string `json:"postcode" bson:"postcode"`
}

type User struct {
	Name  	string 		`json:"name" bson:"user_name"`
	Age  	string 		`json:"age" bson:"user_age"`
	Address Address  	`json:"address" bson:"user_address"`
}