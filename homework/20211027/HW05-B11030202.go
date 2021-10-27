/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package main

import (
	"fmt"
	"strings"
)

var functions = [...]string{
	"管理 - 新增電影",
	"管理 - 檢視菜單",
	"銷售",
	"退出",
}

var menu = make(map[string]int)

func showMenu() {
	fmt.Println("..... Movies .....")
	if len(menu) > 0 {
		i := 0
		for name, price := range menu {
			fmt.Printf("(%d)\t%s ... %d\n", i+1, name, price)
			i++
		}
	} else {
		fmt.Println("目前沒有任何電影")
	}
	fmt.Println("..................")
}

func addMovie(name string, price int) {
	if _, exist := menu[name]; exist {
		showError("「" + name + "」名稱重複")
	} else {
		menu[name] = price
	}
}

func selectMovie(index int) {
	i := 0
	for name, price := range menu {
		if index == i {
			fmt.Printf("「%s」購買成功。請收款 %d 元\n", name, price)
			return
		}
		i++
	}
	showError("無此選項")
}

func showError(reason string) {
	fmt.Printf("執行錯誤。原因：%s\n", reason)
}

func main() {

	for {
		// Header
		fmt.Println(strings.Repeat("=", 10), "NTUST影廳 POS系統", strings.Repeat("=", 10))

		// Functions menu
		for i, name := range functions {
			fmt.Printf("(%d)\t%s\n", i+1, name)
		}

		// Select function
		var selected int
		fmt.Print("請輸入代號：")
		fmt.Scanln(&selected)
		switch selected {
		case 1:
			var name string
			var price int
			fmt.Print("請輸入要新增的電影名稱：")
			fmt.Scanln(&name)
			fmt.Print("請輸入要新增的電影價格：")
			fmt.Scanln(&price)
			addMovie(name, price)
		case 2:
			showMenu()
			break
		case 3:
			var id int
			fmt.Print("請輸入電影編號：")
			fmt.Scanln(&id)
			selectMovie(id - 1)
			break
		case 4:
			fmt.Println("感謝使用，祝您生意興隆")
			return
		default:
			showError("無此選項")
			break
		}

		// Footer
		fmt.Println(strings.Repeat("=", 39))
		fmt.Println()
	}

}
