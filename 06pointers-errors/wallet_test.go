package bitcoin

import (
	"testing"
	// "fmt"
)

func TestBitcoin(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		// fmt.Printf("address of balance in test is: %p \n", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})
	t.Run("withdrawal with error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(15))

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Error aaya hi nahi")
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
