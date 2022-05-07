package Moudle

import "fmt"

type accounts struct {
	user_name string
	pwd       string
	banlase   float64
	deposit   float64
}

func NewAccounts(user_name string, pwd string, deposit float64) *accounts {

	/*
		要求：
			 长度不可低于（6||10，密码必须大于6位数
	*/

	if len(user_name) < 6 || len(user_name) > 10 {

		fmt.Println("名称输入不合规")
		return nil
	}

	if len(pwd) < 6 {
		fmt.Println("密码小于约定位数")
		return nil
	}

	if deposit < 100 {
		fmt.Println("存款数必须大于约定位数")
		return nil
	}

	return &accounts{
		user_name: user_name,
		pwd:       pwd,
		banlase:   0,
		deposit:   deposit,
	}

}

// 存款

func (a *accounts) SetBalance(money float64, pwd string) {

	if pwd != a.pwd {
		fmt.Println("passwd 不正确")
		return
	}

	if money <= 0 {
		panic("不能等于0")
		return

	} else {
		a.deposit += money
		fmt.Printf("存款成功，当前余额为：%v", a.deposit)
	}

	fmt.Println("\t")
}

// 查看余额
func (a *accounts) GetDeposit(user_name string, pwd string) {

	if pwd != a.pwd {
		fmt.Println("passwd 不正确")
		return
	}

	fmt.Printf("当前用户：%v, 账户余额为:%v", user_name, a.deposit)

	fmt.Println("\t")

}

// 取款
func (a *accounts) SetMoney(user_name string, pwd string, money float64) {

	if pwd != a.pwd {
		fmt.Println("passwd 不正确")
		return
	} else {
		fmt.Printf("欢迎用户登陆")
	}

	if money <= 0 {
		panic("取款金额不能等于0")
		return
	} else {
		a.deposit -= money
		fmt.Printf("当前用户余额为: %v", a.deposit)
	}

	fmt.Println("\t")
}
