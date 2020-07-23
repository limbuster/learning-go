package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but did not want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		want := Bitcoin(10)
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient fund", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
