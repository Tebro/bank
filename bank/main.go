package main

import (
	"bufio"
	"fmt"
	"github.com/Tebro/bank/customer"
	"os"
	"strconv"
	"strings"
)

func QueryUserForString(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	return strings.TrimSuffix(text, "\n")
}

func QueryUserForFloat(message string) float64 {
	text := QueryUserForString(message)
	i, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	return i
}
func main() {

	first := QueryUserForString("Enter first name.")
	last := QueryUserForString("Enter last name.")

	c := customer.NewCustomer(first, last)

	for {
		addAccount := QueryUserForString("Add account? (y/N)")
		if strings.ContainsAny(addAccount, "yY") {
			a := c.AddAccount()
			amount := QueryUserForFloat("How much to add to account?")
			a.AddTransaction(amount)
		} else {
			break
		}
	}

	balance := c.GetBalance()
	fmt.Printf("%v %v has %v account(s) with a total value of: %v\n",
		c.GetFirst(),
		c.GetLast(),
		c.GetNumAccounts(),
		balance)
}
