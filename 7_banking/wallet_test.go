package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})

	// Failing tests:

	t.Run("withdraw exact balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(20))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("deposit and withdraw multiple times", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(15))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("overdraw after multiple transactions", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(30))
		err1 := wallet.Withdraw(Bitcoin(10))
		err2 := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err1)
		assertNoError(t, err2)
	})

	t.Run("concurrent deposits", func(t *testing.T) {
		wallet := Wallet{}
		done := make(chan bool)

		go func() {
			wallet.Deposit(Bitcoin(10))
			done <- true
		}()

		go func() {
			wallet.Deposit(Bitcoin(20))
			done <- true
		}()

		<-done
		<-done

		assertBalance(t, wallet, Bitcoin(30))
	})

	t.Run("concurrent withdrawals", func(t *testing.T) {
		wallet := Wallet{Bitcoin(50)}
		done := make(chan bool)

		go func() {
			wallet.Withdraw(Bitcoin(30))
			done <- true
		}()

		go func() {
			wallet.Withdraw(Bitcoin(20))
			done <- true
		}()

		<-done
		<-done

		assertBalance(t, wallet, Bitcoin(20))
	})

}

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

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
