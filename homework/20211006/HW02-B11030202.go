/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞 B11030202
 *   All rights reserved.
 *   NTUST
 */

// BMI計算機

package main

import (
	"fmt"
	"strings"
)

// Account config
const password = "11030202"

// Current money
var current_money float64 = 0

func login() bool {
	// Login
	var password_input string
	fmt.Print("請輸入密碼：")
	fmt.Scanln(&password_input)

	// Check login information
	return password_input == password
}

func atm() {
	var money_save float64
	var money_give float64

	fmt.Print("請輸入整數存款金額：")
	fmt.Scanln(&money_save)
	fmt.Print("請輸入整數提款金額：")
	fmt.Scanln(&money_give)

	current_money += money_save
	current_money -= money_give

	fmt.Printf("最終餘額：%.f\n", current_money)
}

func main() {
	// Title
	fmt.Println(strings.Repeat("=", 20), "台科銀行ATM", strings.Repeat("=", 20))

	// Login
	login_sts := login()
	fmt.Println()

	// If login is successful
	// call atm function
	if login_sts {
		fmt.Println("歡迎使用")
		atm()
	} else {
		fmt.Println("登入失敗")
	}

	// Thanks message
	fmt.Println()
	fmt.Println("謝謝使用~")
}
