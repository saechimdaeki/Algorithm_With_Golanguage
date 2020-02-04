package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//(a Account)가 receiver역할 Deposit amount on account
func (a *Account) Deposit(amount int) {
	//fmt.Println("deposint", amount)
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("돈을 더이상 빼낼수없다 너는 돈이없어 ㅠㅠㅠ")

// 돈빼내기
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//change owneraccount
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//owner name
func (a *Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account \n He(she) have money :", a.Balance())
}
