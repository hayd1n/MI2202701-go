/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package main

import (
	"B11030202/student"
	"fmt"
)

func main() {
	var s student.Student
	for {
		fmt.Println("=========================")
		fmt.Printf("%d\t%s\n", 1, "設定資訊")
		fmt.Printf("%d\t%s\n", 2, "檢視資訊")
		fmt.Printf("%d\t%s\n", 0, "離開")
		fmt.Println("=========================")
		var selected int
		fmt.Print("請輸入選項代碼：")
		fmt.Scanln(&selected)
		switch selected {
		case 1:
			s.SetData()
		case 2:
			s.ShowInfo()
		case 0:
			return
		default:
			fmt.Println("錯誤！原因：無此選項。")
		}
	}
}
