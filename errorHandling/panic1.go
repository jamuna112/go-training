package errorhandling

import "fmt"

/*
Writing a simple banking program where user cannot withdraw more amount than the user has on their account
. Meaning user cannot go negative on their balance. If user try to withdraw more money then program
should panic, critical failure.
*/
type BankAccount struct {
	Balance float32
}

func (ba *BankAccount) WithdrawMoney(withdrawAmount float32) float32 {

	if ba.Balance < withdrawAmount {
		panic("withdraw amount is greater than balance ")
	} else {
		ba.Balance = ba.Balance - withdrawAmount
	}
	fmt.Printf("balance amount: %v\n", ba.Balance)
	return ba.Balance
}
