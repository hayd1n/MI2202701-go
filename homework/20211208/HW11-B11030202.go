/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package main

import "fmt"

type Product interface {
	setData()
	showData()
	getDescription() string
}

type Noodle struct {
	name  string
	price int
}

func (n *Noodle) setData() {
	fmt.Print("請輸入商品名稱：")
	fmt.Scanln(&n.name)
	fmt.Print("請輸入商品價格：")
	fmt.Scanln(&n.price)
}

func (n Noodle) showData() {
	fmt.Printf("%s ($%d)", n.name, n.price)
}

func (n Noodle) getDescription() string {
	return "打開包裝 > 加入調料與熱水 > 靜置三分鐘 > 享受美味餐點"
}

type Drink struct {
	name  string
	price int
}

func (d *Drink) setData() {
	fmt.Print("請輸入商品名稱：")
	fmt.Scanln(&d.name)
	fmt.Print("請輸入商品價格：")
	fmt.Scanln(&d.price)
}

func (d Drink) showData() {
	fmt.Printf("%s ($%d)", d.name, d.price)
}

func (d Drink) getDescription() string {
	return "用力拉起拉環 > 飲用好喝的飲料"
}

var list = make([]Product, 0)

func showMenu() {
	fmt.Println("============= NTUST =============")
	fmt.Printf("%d)\t%s\n", 1, "顯示清單")
	fmt.Printf("%d)\t%s\n", 2, "新增項目")
	fmt.Printf("%d)\t%s\n", 0, "離開")
	fmt.Println("=================================")
}

func addProduct() {
	var p Product
	fmt.Println("--------------")
	fmt.Printf("%d)\t%s\n", 1, "泡麵")
	fmt.Printf("%d)\t%s\n", 2, "飲料")
	fmt.Println("--------------")
	var selected int
	fmt.Print("請輸入商品代碼：")
	fmt.Scanln(&selected)
	switch selected {
	case 1:
		p = &Noodle{}
		break
	case 2:
		p = &Drink{}
		break
	default:
		showError("無此商品代碼")
		return
	}
	p.setData()
	list = append(list, p)
}

func showList() {
	for i, p := range list {
		fmt.Printf("(%.2d)\t", i+1)
		p.showData()
		fmt.Printf("\t使用說明：%s\n", p.getDescription())
	}
}

func showError(message string) {
	fmt.Printf("錯誤！原因：%s\n", message)
}

func main() {
	for {
		showMenu()

		var selected int
		fmt.Print("請輸入代碼：")
		fmt.Scanln(&selected)

		switch selected {
		case 1:
			showList()
			break
		case 2:
			addProduct()
			break
		case 0:
			return
		default:
			showError("無此選項")
			break
		}
	}
}
