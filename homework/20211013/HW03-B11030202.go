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

// item type
type item struct {
	name  string // 品名
	price int    // 售價
}

// items array
var items = [...]item{
	{"拉麵", 200},
	{"泡芙", 30},
	{"洋芋片", 35},
}

func main() {

	// Header
	fmt.Println(strings.Repeat("=", 20), "米花科技公司食品自動販賣機", strings.Repeat("=", 20))

	// Save money
	var money int = 0
	fmt.Print("請輸入使用者錢包餘額：")
	fmt.Scanln(&money)

	// Items list
	fmt.Println("編號\t品名\t金額")
	for i, itemInfo := range items {
		fmt.Printf("%d\t%s\t%d\n", i+1, itemInfo.name, itemInfo.price)
	}

	// Select which to buy
	var selectIndex int = 0
	fmt.Print("請輸入要購買的食物編號：")
	fmt.Scanln(&selectIndex)
	selectIndex -= 1

	fmt.Println("========== 交易結果 ==========")

	// Buy item
	if selectIndex >= 0 && selectIndex < len(items) {
		if money >= items[selectIndex].price {
			money -= items[selectIndex].price
			fmt.Println("成功。")
			fmt.Printf("品名：%s($%d)\n", items[selectIndex].name, items[selectIndex].price)
			fmt.Printf("帳戶餘額：%d\n", money)
		} else {
			fmt.Println("失敗。")
			fmt.Println("原因：餘額不足！")
		}
	} else {
		fmt.Println("失敗。")
		fmt.Println("原因：商品編號無效！")
	}

	// Footer
	fmt.Println(strings.Repeat("=", 68))
}
