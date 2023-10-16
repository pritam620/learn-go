package pointers

import "fmt"

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type stringer interface {
	String() string
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance = wallet.balance + amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) {
	wallet.balance -= amount
}

func (wallet Wallet) Balance() Bitcoin {
	return wallet.balance
}
