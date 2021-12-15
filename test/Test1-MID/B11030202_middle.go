/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type User struct {
	username string
	name     string
	password string
	admin    bool
}

var userlist = map[string]User{}

var logined_user = ""
var logined = false

func (u User) toString() string {
	return fmt.Sprintf("%-10s\t%s", u.username, u.name)
}

func login() {
	var username string
	var password string
	fmt.Print("請輸入帳號：")
	fmt.Scanln(&username)
	fmt.Print("請輸入密碼：")
	fmt.Scanln(&password)
	if _, ok := userlist[username]; ok {
		if userlist[username].password == password {
			logined_user = username
			logined = true
			fmt.Println("登入成功")
			return
		}
	}
	fmt.Println("登入失敗")
}

func logout() {
	logined_user = ""
	logined = false
}

func showMenu() {
	fmt.Println(strings.Repeat("=", 20), "會員管理系統", strings.Repeat("=", 20))
	if logined {
		fmt.Printf("使用者：%s\n", userlist[logined_user].name)
	} else {
		fmt.Printf("尚未登入\n")
	}
	fmt.Println(strings.Repeat(".", 40))
	if logined {
		if userlist[logined_user].admin {
			fmt.Printf("(%d)\t%s\t%s\n", 1, "使用者清單", "檢視所有使用者的資料")
			fmt.Printf("(%d)\t%s\t%s\n", 2, "新增使用者", "建立使用者帳號")
			fmt.Printf("(%d)\t%s\t%s\n", 3, "修改使用者", "更改使用者資料")
			fmt.Printf("(%d)\t%s\t%s\n", 4, "刪除使用者", "刪除使用者資料")
			fmt.Printf("(%d)\t%s\t%s\n", 5, "備份使用者", "匯出所有帳號為文字檔")
		} else {
			fmt.Printf("(%d)\t%s\t%s\n", 1, "修改使用者", "更改使用者資料")
		}
		fmt.Printf("(%d)\t%s\t\t%s\n", 0, "登出", "登出使用者，回到主選單")
	} else {
		fmt.Printf("(%d)\t%s\t%s\n", 1, "登入", "登入使用者後臺功能")
		fmt.Printf("(%d)\t%s\t%s\n", 0, "離開", "結束使用程式")

	}
	fmt.Println(strings.Repeat("=", 55))
}

func showUsers() []string {
	fmt.Println(strings.Repeat(".", 40))
	var i = 0
	var usernames = make([]string, 0)
	for _, user := range userlist {
		fmt.Printf("%d\t%s\n", i+1, user.toString())
		usernames = append(usernames, user.username)
		i++
	}
	fmt.Println(strings.Repeat(".", 40))
	return usernames
}

func addUser() {
	var username string
	var password string
	var name string
	fmt.Print("請輸入帳號：")
	fmt.Scanln(&username)
	fmt.Print("請輸入密碼：")
	fmt.Scanln(&password)
	fmt.Print("請輸入姓名：")
	fmt.Scanln(&name)
	if _, ok := userlist[username]; ok {
		fmt.Println("新增失敗。原因：使用者帳號重複")
		return
	}

	userlist[username] = User{
		username: username,
		password: password,
		name:     name,
	}

	fmt.Println("新增成功")
}

func editUser() {
	if userlist[logined_user].admin {
		var number int
		list := showUsers()
		fmt.Print("請輸入使用者編號：")
		fmt.Scanln(&number)
		if number > 0 && number <= len(userlist) {
			var i = 0
			for username, user := range userlist {
				if list[number-1] == username {
					// fmt.Printf("----%s-----", username)
					var password string
					var name string
					fmt.Print("請輸入新密碼：")
					fmt.Scanln(&password)
					fmt.Print("請輸入新姓名：")
					fmt.Scanln(&name)
					userlist[username] = User{
						username: user.username,
						password: password,
						name:     name,
						admin:    user.admin,
					}
					fmt.Println("編輯成功")
					break
				}
				i++
			}
		} else {
			fmt.Println("編輯失敗。原因：無此使用者")
		}
	} else {
		editUser_me()
	}
}

func editUser_me() {
	var password string
	var name string
	fmt.Print("請輸入新密碼：")
	fmt.Scanln(&password)
	fmt.Print("請輸入新姓名：")
	fmt.Scanln(&name)
	userlist[logined_user] = User{
		username: userlist[logined_user].username,
		password: password,
		name:     name,
		admin:    userlist[logined_user].admin,
	}
	fmt.Println("編輯成功")
}

func deleteUser() {
	list := showUsers()
	var number int
	fmt.Print("請輸入使用者編號：")
	fmt.Scanln(&number)
	if number > 0 && number <= len(userlist) {
		var i = 0
		for username := range userlist {
			if list[number-1] == username {
				// fmt.Printf("----%s----", username)
				delete(userlist, username)
				fmt.Println("刪除成功")
				return
			}
			i++
		}
	} else {
		fmt.Println("刪除失敗。原因：無此使用者")
	}
}

func backupUser() {
	var exportString string = ""

	var i = 0
	for _, user := range userlist {
		exportString += fmt.Sprintf(
			"%d\t%s\n",
			i,
			user.toString(),
		)
		i++
	}

	// 轉換成Byte型別
	data_byte := []byte(exportString)

	// 輸出檔案
	err := ioutil.WriteFile("Users.txt", data_byte, 0644)
	if err != nil {
		fmt.Println("匯出失敗。原因：此目錄拒絕存取")
	} else {
		fmt.Println("匯出成功")
	}
}

func main() {
	userlist["B11030202"] = User{
		username: "B11030202",
		password: "0202",
		name:     "張皓鈞",
		admin:    true,
	}
	for {
		showMenu()
		var selected int
		fmt.Print("請輸入選項：")
		fmt.Scanln(&selected)
		if logined {
			if userlist[logined_user].admin {
				switch selected {
				case 1:
					showUsers()
					break
				case 2:
					addUser()
					break
				case 3:
					editUser()
					break
				case 4:
					deleteUser()
					break
				case 5:
					backupUser()
					break
				case 0:
					logout()
					break
				default:
					fmt.Println("失敗。原因：無此選項")
				}
			} else {
				switch selected {
				case 1:
					editUser()
					break
				case 0:
					logout()
					break
				default:
					fmt.Println("失敗。原因：無此選項")
				}
			}
		} else {
			switch selected {
			case 1:
				login()
				break
			case 0:
				fmt.Println("感謝您的使用~")
				return
			default:
				fmt.Println("失敗。原因：無此選項")
			}
		}
	}
}
