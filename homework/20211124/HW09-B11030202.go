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

var students = make(map[string]string)

func showMenu() {
	fmt.Println(strings.Repeat("=", 10), "資管系學生手冊", strings.Repeat("=", 10))
	fmt.Printf("(%d)\t%s\n", 1, "顯示學生清單")
	fmt.Printf("(%d)\t%s\n", 2, "新增學生")
	fmt.Printf("(%d)\t%s\n", 3, "刪除學生")
	fmt.Printf("(%d)\t%s\n", 0, "離開")
	fmt.Println(strings.Repeat("=", 36))
}

func selectMode() int {
	var selected int
	fmt.Print("請輸入功能代碼數字：")
	fmt.Scanln(&selected)
	return selected
}

func showStudentsList() {
	defer handleException()

	if len(students) <= 0 {
		panic("目前尚未有學生資料")
	}

	var i int = 0
	for id, name := range students {
		fmt.Printf("(%d)\t%s\t%s\n", i+1, id, name)
		i++
	}
}

func addStudent() {
	defer handleException()

	var id string
	var name string
	fmt.Print("請輸入學生「學號」：")
	fmt.Scanln(&id)
	fmt.Print("請輸入學生「姓名」：")
	fmt.Scanln(&name)
	if _, ok := students[id]; ok {
		// 學號已存在
		panic("此學號已經存在，無法建立")
	}
	students[id] = name
}

func delStudent() {
	defer handleException()

	var id string
	fmt.Print("請輸入學生「學號」：")
	fmt.Scanln(&id)
	if _, ok := students[id]; !ok {
		panic("此學號尚不存在，無法刪除")
	}

	delete(students, id)
	fmt.Println("刪除成功")
}

func showError(message string) {
	fmt.Printf("發生錯誤。原因：%s\n", message)
}

func handleException() {
	if r := recover(); r != nil {
		fmt.Println("發生錯誤。原因：", r)
	}
}

func main() {
	for {
		showMenu()
		mode := selectMode()
		switch mode {
		case 1:
			showStudentsList()
			break
		case 2:
			addStudent()
			break
		case 3:
			delStudent()
			break
		case 0:
			return
		default:
			showError("無此選項")
			break
		}
		fmt.Println()
	}
}
