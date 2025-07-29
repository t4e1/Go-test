package accounts

import (
	"errors"
	"fmt"
)

// Account struct
// private struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// construct fucnction
// Go에서는 constructor를 지원하지 않기 떄문에 함수로 비슷한 기능을 만들어 동작시킨다.
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on you account (method)
// Receiver (a)
// struct의 첫 글자를 따서 짓는게 컨벤션임
// Go 는 function이나 method를 실행할 때 객체의 카피본을 가져온다.
// 그래서 (a Account)로 리시버를 줘버리면, 실제 사용되는 Account가 그 Account의 copy를 만들어오기 때문에 변경 사항이 있으면 적용이 안됨.
// 필요한 경우 *를 사용해서 직접 지정해줘야함
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw your balance
// Error Control
// GO 에는 try-catch 등의 에러를 잡아주는 기능이 없음.
// 직접 에러 핸들링을 해줘야함
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
