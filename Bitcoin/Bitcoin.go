package bitcoin

import (
	"errors"
	"fmt"
)

// In Go, errors are values, so we can refactor it out into a variable
// and have a single source of truth for it.

var ErrNotEnoughBalance = errors.New("cannot withdraw, insufficient funds")

// Go lets you create new types from existing ones.
//The syntax is type MyName OriginalType

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// This stringer interface is defined in the fmt package and lets you define how your type is printed
// when used with the %s format string in prints.
type Stringer interface {
	String() string
}

// https://gobyexample.com/pointers
// read as "a pointer to a wallet".
func (w *Wallet) Deposit(amount Bitcoin) {
	// These pointers to structs even have their own name: struct pointers
	// and they are automatically dereferenced.
	w.balance += amount
	//fmt.Printf("Address of wallet in deposit is is %p ", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	// you can use :  (*w).balance too
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	/*

		Errors can be nil because the return type of Withdraw will be error, which is an interface.
		If you see a function that takes arguments or returns values that are interfaces,
		they can be nillable.

		NOTE : Like null if you try to access a value that is nil it will throw a runtime panic.
		This is bad! You should make sure that you check for nils.

	*/
	if amount > w.balance {
		// errors.New creates a new error with a message of your choosing.
		return ErrNotEnoughBalance
	}
	w.balance -= amount
	return nil
}
