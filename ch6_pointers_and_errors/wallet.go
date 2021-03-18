package pointers_and_errors

type Bitcoin int

type Wallet struct {
	// 소문자로 썼으면 외부에 공개되지 않는 private 변수
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
