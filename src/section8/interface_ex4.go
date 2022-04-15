package main

import (
	"fmt"
	"strconv"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

func structToMsg(arg interface{}) string {
	switch arg.(type) {
	case Account:
		a := arg.(Account)
		return "Account : " + a.number + " : " + strconv.FormatFloat(a.balance*a.interest+a.balance, 'f', -1, 64)
	case *Account:
		a := arg.(*Account)
		return "*Account : " + a.number + " : " + strconv.FormatFloat(a.balance*a.interest+a.balance, 'f', -1, 64)
	default:
		return "ERROR"
	}
}

func calculateBalance(arg interface{}) {
	switch arg.(type) {
	case Account:
		a := arg.(Account)
		a.balance = a.balance*a.interest + a.balance
	case *Account:
		a := arg.(*Account)
		a.balance = a.balance*a.interest + a.balance
	default:
		fmt.Println("ERROR")
	}
}

func main() {
	//ex1

	fmt.Println(structToMsg(Account{number: "245-901", balance: 10000000, interest: 0.015}))
	fmt.Println(structToMsg(&Account{number: "245-902", balance: 12000000, interest: 0.025}))

	var user = new(Account)
	user.number = "245-903"
	user.balance = 15000000
	user.interest = 0.035
	fmt.Println(structToMsg(user))

	//ex2
	// 값 전달, 포인터전달 차이 복습
	a1 := Account{number: "245-901", balance: 10000000, interest: 0.015}
	a2 := &Account{number: "245-902", balance: 12000000, interest: 0.025}

	fmt.Println("before : ", int(a1.balance))
	fmt.Println("before : ", int(a2.balance))
	fmt.Println("before : ", int(user.balance))

	calculateBalance(a1)
	calculateBalance(a2)
	calculateBalance(*user)

	fmt.Println("after : ", int(a1.balance))
	fmt.Println("after : ", int(a2.balance))
	fmt.Println("after : ", int(user.balance))

}
