package main

import (
	
	"fmt"
	//"../../../Desktop/Study_all_by_myself_-GO-/BankProject_Go/banking"  //나의 폴더.......... 
	"github.com/saechimdaeki/study_Go_lang/banking"
)
//import "../../../Desktop/Study_all_by_myself_-GO-/BankProject_Go/banking"  //나의 폴더.......... 
//import "github.com/saechimdaeki/study_Go_lang/BankProject_Go"
//C:\Users\anima\Desktop\Study_all_by_myself_-GO-

func main() {
	
	account:=banking.Account{Owner:"kimjunseong" , Balance:10000}
	fmt.Println(account)
}
