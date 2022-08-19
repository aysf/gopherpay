package main

import (
	"fmt"
	"strings"
	"tesgo/gopherpay/payment/oo"
)

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	cash := &oo.Cash{}
	cash.ProcessPayment(amount)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	credit := oo.CreateCreditAccount(
		"Don Krieg",
		"1111-2222-3333-4444",
		5,
		2023,
		123,
	)

	fmt.Println("Paying with credit card")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit())
	credit.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", credit.AvailableCredit())
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := oo.CreateCheckingAccount(
		"Don Krieg",
		"10011000",
		"0343433334",
	)

	fmt.Println("Payment with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance())
	checking.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", checking.Balance())

	fmt.Println("hmmm.., not enough in the account. Nothing we can do!")
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
}
