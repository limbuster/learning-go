package main

import (
	"errors"
	"fmt"
)

// Bitcoin value
type Bitcoin int

// ErrInsufficientFunds error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Stringer returns value in string
// type Stringer interface {
// 	String() string
// }

// Wallet store balance
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit amount to balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw amount from balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance returns the current balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
