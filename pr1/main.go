package main

import (
	"fmt"
	"log"

	"github.com/t4e1/go-prac1/pr1/accounts"
	mydict "github.com/t4e1/go-prac1/pr1/dict"
)

func main() {

	// account := banking.Account{Owner: "jtw", Balance: 1000} // Owner와 Balance 둘 다 public 이므로 수정이 가능함. Constructor 로 만들어야함
	account := accounts.NewAccount("jtw")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	err2 := dictionary.Add("hello", "Greeting")
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)

	err3 := dictionary.Add("hello", "Greeting")
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(dictionary)
	err4 := dictionary.Update("hello", "Second")
	if err4 != nil {
		fmt.Println(err4)
	}
	word2, _ := dictionary.Search("hello")
	fmt.Println(word2)
}
