package awesomeProject

import (
	"errors"
	"sync"
)

// Кошелек для хранения биткойнов
type Wallet struct {
	balance float64
	mutex   sync.Mutex
}

// Ошибка при попытке снятия средств, когда баланс недостаточен
var ErrInsufficientFunds = errors.New("недостаточно средств на кошельке")

// Вносит деньги на кошелек
func (w *Wallet) Deposit(amount float64) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.balance += amount
}

// Снимает деньги с кошелька
func (w *Wallet) Withdraw(amount float64) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Получает баланс кошелька
func (w *Wallet) Balance() float64 {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.balance
}
