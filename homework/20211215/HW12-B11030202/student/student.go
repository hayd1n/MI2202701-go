/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package student

import "fmt"

type Student struct {
	info Info
}

func (s *Student) setId(id string) {
	s.info.Id = id
}

func (s *Student) setName(name string) {
	s.info.Name = name
}

func (s *Student) setGender(gender bool) {
	s.info.IsMale = gender
}

func (s *Student) SetData() {
	fmt.Println("-------------")
	fmt.Printf("%d\t%s\n", 1, "學號")
	fmt.Printf("%d\t%s\n", 2, "姓名")
	fmt.Printf("%d\t%s\n", 3, "性別")
	fmt.Println("-------------")
	var selected int
	fmt.Print("請輸入選項代碼：")
	fmt.Scanln(&selected)
	switch selected {
	case 1:
		var id string
		fmt.Print("請輸入學生學號：")
		fmt.Scanln(&id)
		s.setId(id)
	case 2:
		var name string
		fmt.Print("請輸入學生姓名：")
		fmt.Scanln(&name)
		s.setName(name)
	case 3:
		var gender_input string
		fmt.Print("請輸入學生性別(男/女)：")
		fmt.Scanln(&gender_input)
		switch gender_input {
		case "男":
			s.setGender(true)
		case "女":
			s.setGender(false)
		default:
			fmt.Println("錯誤！原因：輸入內容不合法。")
		}
	default:
		fmt.Println("錯誤！原因：無此選項。")
	}
}

func (s *Student) ShowInfo() {
	var id string = "尚未設定"
	if s.info.Id != "" {
		id = s.info.Id
	}
	var name string = "尚未設定"
	if s.info.Name != "" {
		name = s.info.Name
	}
	fmt.Println("-------------")
	fmt.Printf("學號：%s\n", id)
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("姓別：%s\n", s.info.GetGender())
	fmt.Println("-------------")
}
