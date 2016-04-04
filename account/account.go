package account

type Account struct {
	key          int
	transactions []Transaction
}

func NewAccount(id int) *Account {
	return &Account{id, []Transaction{}}
}

func (a *Account) GetBalance() float64 {
	sum := 0.0
	for _, t := range a.transactions {
		sum += t.value
	}
	return sum
}

func (a *Account) AddTransaction(value float64) {
	t := Transaction{value}
	a.transactions = append(a.transactions, t)
}
