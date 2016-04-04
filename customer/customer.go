package customer

import (
	"github.com/Tebro/bank/account"
)

type Customer struct {
	firstname string
	lastname  string
	accounts  []*account.Account
}

func NewCustomer(first string, last string) Customer {
	return Customer{first, last, []*account.Account{}}
}

func (c *Customer) AddAccount() *account.Account {
	var a *account.Account = account.NewAccount(len(c.accounts) + 1)
	c.accounts = append(c.accounts, a)
	return a
}

func (c *Customer) GetBalance() float64 {
	sum := 0.0
	for _, a := range c.accounts {
		sum += a.GetBalance()
	}
	return sum
}

func (c *Customer) GetFirst() string {
	return c.firstname
}

func (c *Customer) GetLast() string {
	return c.lastname
}

func (c *Customer) GetNumAccounts() int {
	return len(c.accounts)
}
