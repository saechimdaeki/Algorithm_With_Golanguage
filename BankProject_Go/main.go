package main

import (
	"fmt"
	//"log"

	"../../../Desktop/Study_all_by_myself_-GO-/BankProject_Go/accounts" //나의 컴퓨타폴더..........
	//"github.com/saechimdaeki/study_Go_lang/banking"
)

//import "../../../Desktop/Study_all_by_myself_-GO-/BankProject_Go/banking"  //나의 폴더..........
//import "github.com/saechimdaeki/study_Go_lang/BankProject_Go"
//C:\Users\anima\Desktop\Study_all_by_myself_-GO-

func main() {

	account := accounts.NewAccount("junSeong")
	//fmt.Println(account)
	account.Deposit(100)
	fmt.Println(account.Balance())
	/*
		err := account.Withdraw(90)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(account.Balance(), account.Owner())
	*/

	fmt.Println(account)
}
