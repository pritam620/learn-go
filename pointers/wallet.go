package pointers

type Wallet struct {
	balance int
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.balance = wallet.balance + amount
}

func (wallet Wallet) Balance() int {
	return wallet.balance
}
