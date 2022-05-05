package main

import (
	"awesomeProject/Moudle"
)

func main() {
	res := Moudle.NewAccounts("zhangsan", "123123", 200)
	res.SetMoney("zhangsan", "123123", 100)
	res.SetBalance("zhangsan", 1000, "123123")
	res.GetDeposit("zhangsan", "123123")
}
