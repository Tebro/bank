package account

import "testing"

func TestAccountCreation(t *testing.T) {
	a := NewAccount(1)
	if a.id != 1 {
		t.Error("The id does not match.")
	}
	if len(a.transactions) > 0 {
		t.Error("The list should be empty.")
	}
}

func TestAddingValueToAccount(t *testing.T) {
	a := NewAccount(1)
	a.AddTransaction(100)
	if a.GetBalance() != 100 {
		t.Error("The amount does not match.")
	}
}

func TestSumOfSeveralTransactions(t *testing.T) {
	a := NewAccount(1)

	t1 := 1000.0
	t2 := 321.3
	t3 := 425.93
	sum := t1 + t2 + t3

	a.AddTransaction(t1)
	a.AddTransaction(t2)
	a.AddTransaction(t3)

	returned := a.GetBalance()
	if returned != sum {
		t.Error("The returned sum does not match.")
	}
}
