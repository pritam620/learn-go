package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("expected %s, but go %s", want, got)
	}
}
