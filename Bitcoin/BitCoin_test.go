//https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

package bitcoin

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

		/*
			When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
			You can see that the addresses of the two balances are different. So when we change the
			value of the balance inside the code, we are working on a copy of what came from the test.
			Therefore the balance in the test is unchanged.

		*/

		//fmt.Printf("Address of wallet in test is %p ", &wallet.balance)

	})
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw from insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrNotEnoughBalance)
		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Error("Wanted an error but did not get one")
		}
	})
}

// To make Bitcoin you just use the syntax Bitcoin(999).
//By doing this we're making a new type and we can declare methods on them.
//This can be very useful when you want to add some domain specific functionality on top of existing types.

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
