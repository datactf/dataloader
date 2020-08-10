package data

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"time"
)

type Account struct {
	Id     string  `json:"id" bson:"_id" faker:"uuid_digit"`
	Logins []Login `json:"logins" bson:"logins"`
}

type User struct {
	Email    string `json:"email" bson:"email" faker:"email"`
	UserName string `json:"username" bson:"username" faker:"username"`
}

type Login struct {
	Website    string    `json:"website" bson:"website" faker:"email`
	MacAddress string    `json:"mac" bson:"mac" faker:"mac_address"`
	IpAddress  string    `json:"ip" bson:"ip" faker:"ipv6"`
	Date       time.Time `json:"date" bson:"login_date"`
	User       User      `json:"user" bson:"user"`
}

//GenerateAccount generates a fake Account structure.
func GenerateAccount() (*Account, error) {
	account := Account{}
	err := faker.FakeData(&account)
	if err != nil {
		fmt.Println("could not generate account data")
		return nil, fmt.Errorf("could not generate account data %s", err)
	}
	return &account, nil
}

//GenerateAccounts generates an Accounts array of size `n`
func GenerateAccounts(n int) []Account {
	accounts := make([]Account, n)
	for i := 0; i < n; i++ {
		acc, err := GenerateAccount()
		if err != nil {
			continue
		}
		accounts[i] = *acc
	}
	return accounts
}
