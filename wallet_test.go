package awesomeProject

import "testing"

func TestWallet_Deposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	assertBalance(t, wallet, 10.0)
}

func TestWallet_WithdrawSufficientFunds(t *testing.T) {
	wallet := Wallet{balance: 20}
	err := wallet.Withdraw(10)

	assertNoError(t, err)
	assertBalance(t, wallet, 10.0)
}

func TestWallet_WithdrawInsufficientFunds(t *testing.T) {
	wallet := Wallet{balance: 20}
	err := wallet.Withdraw(30)

	assertError(t, err, ErrInsufficientFunds)
	assertBalance(t, wallet, 20.0)
}

// Вспомогательная функция для проверки баланса кошелька
func assertBalance(t *testing.T, wallet Wallet, want float64) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("ожидалось %v, получено %v", want, got)
	}
}

// Вспомогательная функция для проверки наличия ошибки
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("ожидалась ошибка, но не было получено никакой ошибки")
	}
	if got != want {
		t.Errorf("ожидалась ошибка '%v', получено '%s'", want, got)
	}
}

// Вспомогательная функция для проверки отсутствия ошибки
func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("не ожидалась ошибка, но получено:", got)
	}
}
