package account

import "testing"

func TestTransactionValue(t *testing.T) {
	trans := Transaction{100}

	if trans.value != 100 {
		t.Error("The value stored is not the one given.")
	}
}
