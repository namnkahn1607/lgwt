package pointer_error

import (
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	if got := wallet.Balance(); got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		// Fatal() will stop the test immediately.
		// We don't want to run the next condition if 'got' was a nil.
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(*testing.T) {
		wallet := Wallet{}

		assertNoError(t, wallet.Deposit(Bitcoin(10)))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		assertNoError(t, wallet.Withdraw(Bitcoin(10)))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		assertError(t, wallet.Withdraw(Bitcoin(100)), ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
