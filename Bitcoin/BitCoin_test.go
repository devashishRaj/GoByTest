package bitcoin

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		/*
			When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
			You can see that the addresses of the two balances are different. So when we change the
			value of the balance inside the code, we are working on a copy of what came from the test.
			Therefore the balance in the test is unchanged.

		*/

		//fmt.Printf("Address of wallet in test is %p ", &wallet.balance)

		if got != want {
			t.Errorf("Got %s wanted %s", got, want)
		}
	})
	t.Run("Withdrawl" ,func(t *testing.T) {
		
	})
}

// To make Bitcoin you just use the syntax Bitcoin(999).
//By doing this we're making a new type and we can declare methods on them.
//This can be very useful when you want to add some domain specific functionality on top of existing types.
