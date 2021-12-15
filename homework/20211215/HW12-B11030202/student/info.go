/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package student

/*
	「Info」結構
	* 名稱開頭「大寫」 - import可以使用
*/
type Info struct {
	Id     string // 學號變數 (名稱開頭「大寫」 - import可以使用)
	Name   string // 姓名變數 (名稱開頭「大寫」 - import可以使用)
	IsMale bool   // 性別變數 (名稱開頭「大寫」 - import可以使用)
}

/*
	取得性別
	* 名稱開頭「大寫」 - import可以使用
*/
func (i Info) GetGender() string {
	if i.IsMale {
		return "男性"
	} else {
		return "女性"
	}
}
