//Задача: Банковский счет
//Создайте интерфейс BankAccount с методами Deposit(amount float64) и Withdraw(amount float64) error.
//Реализуйте интерфейс в структуре SimpleAccount.
//Deposit(amount float64) должен увеличивать баланс счета на переданную сумму.
//Withdraw(amount float64) должен уменьшать баланс счета на переданную сумму.
//	Если сумма для снятия превышает баланс, верните ошибку.
//В функции main создайте экземпляр SimpleAccount, внесите на счет депозит,
//выполните несколько операций снятия и выводите текущий баланс после каждой операции.

package main

type BankAccount interface {
	Plus() int
	Minus() int
	Total() int
	Amount() int
}
type SimpleAccount int

func (s *SimpleAccount) Deposit() int {
	return s.Deposit() + s.Amount()
}

func (s *SimpleAccount) Withdraw() int {
	return s.Deposit()
}
func (s SimpleAccount) Amount() int {

}
