package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

// Wallet is a wallet
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit adds money to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw takes money out of the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the amount in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
