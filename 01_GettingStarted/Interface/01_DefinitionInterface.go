package main

// BankAccount represents a bank account with a balance.
type bankAccount struct {
	balance float64
}

/*
Experience 2:
Chúng ta có thể không cần define interface cho BankAccount struct bởi vì mục đích
interface lúc này là giúp ta có thể dễ quản lý và sử dụng ở các package khác.
*/
type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
}
