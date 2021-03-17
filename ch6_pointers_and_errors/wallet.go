package pointers_and_errors

type Wallet struct {
	// 소문자로 썼으면 외부에 공개되지 않는 private 변수
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
