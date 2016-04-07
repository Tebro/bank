package customer

import "testing"

var first string = "Bob"
var last string = "Wood"

func TestCustomerCreation(t *testing.T) {
	c := NewCustomer(first, last)

	if c.GetFirst() != first {
		t.Error("The first name does not match!")
	}

	if c.GetLast() != last {
		t.Error("The last name does not match!")
	}
	if c.GetNumAccounts() != 0 {
		t.Error("The customer does not have 0 accounts.")
	}
}

func TestAddingAccount(t *testing.T) {
	c := NewCustomer(first, last)
	c.AddAccount()
	c.AddAccount()
	if c.GetNumAccounts() != 2 {
		t.Error("The customer does not have the right amount of accounts.")
	}
}

func TestCalculatingCustomerBalance(t *testing.T) {
	var v1, v2 float64 = 100, 250
	c := NewCustomer(first, last)
	a1 := c.AddAccount()
	a1.AddTransaction(v1)
	a2 := c.AddAccount()
	a2.AddTransaction(v2)

	if c.GetBalance() != (v1 + v2) {
		t.Error("The sum returned does not match the entered amounts.")
	}
}
