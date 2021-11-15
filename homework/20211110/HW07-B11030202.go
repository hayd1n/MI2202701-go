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

// 成績結構體
type Score struct {
	name         string
	chineseScore int
	englishScore int
	mathScore    int
	totalScore   int
	avg          float32
}

var scores []Score

// 顯示菜單
func showMenu() {
	fmt.Println(strings.Repeat("=", 20), "NTUST Score System", strings.Repeat("=", 20))
	fmt.Printf("%d\t%s\n", 1, "匯入學生")
	fmt.Printf("%d\t%s\n", 2, "輸入成績")
	fmt.Printf("%d\t%s\n", 3, "匯出成績")
	fmt.Printf("%d\t%s\n", 0, "結束")
	fmt.Println(strings.Repeat("=", 60))
}

// 選擇功能
func selectFunc() int {
	var selected int
	fmt.Print("請輸入選項：")
	fmt.Scanln(&selected)
	return selected
}

var importedStudents bool = false

// 匯入學生
func importStudents() {
	if !importedStudents {
		// 根據 檔案名稱 呼叫 ReadFile()
		importByte, err := ioutil.ReadFile("StudentList.txt")
		if err != nil { // 若收到的 err 有錯誤
			fmt.Print(err) // 印出錯誤原因
		}
		// 將 byte格式的文字 轉換成 string字串
		importString := string(importByte)

		// 將 string 切割成 array (依據 換行 符號切割)
		// 換行符號可能為 \n 或 \r\n
		importArray_line := strings.Split(importString, "\r\n")

		// 產生對應大小的陣列
		scores = make([]Score, len(importArray_line))

		// 遍歷學生名單
		for i, name := range importArray_line {
			scores[i] = Score{
				name: name,
			}
		}
		importedStudents = true
		fmt.Println("執行成功")
	} else {
		showError("已經匯入過了")
	}
}

// 匯入成績
func importScores() {
	if importedStudents {
		var courses = map[string]int{"國文": 0, "英文": 0, "數學": 0}
		for i, student := range scores {
			fmt.Printf("〔%s〕\n", student.name)
			for courseName := range courses {
				courses[courseName] = inputScore("\t請輸入" + courseName + "成績：")
			}
			scores[i].chineseScore = courses["國文"]
			scores[i].englishScore = courses["英文"]
			scores[i].mathScore = courses["數學"]
		}
		fmt.Println("執行完成")
	} else {
		showError("尚未匯入學生資料")
	}
}

func inputScore(message string) int {
	var score int
	fmt.Printf(message)
	fmt.Scanln(&score)
	if score < 0 || score > 100 {
		score = 0
	}
	return score
}

// 匯出成績
func exportScores() {
	if importedStudents {
		var exportString string = ""

		for i, score := range scores {
			var totalScore int = score.chineseScore + score.englishScore + score.mathScore
			var avg float32 = float32(totalScore) / 3
			scores[i].totalScore = totalScore
			scores[i].avg = avg
			formatStr := "〔%s〕\n\t請輸入國文成績：%d\n\t請輸入英文成績：%d\n\t請輸入數學成績：%d\n總分：%d / 平均：%.2f\n\n"
			exportString += fmt.Sprintf(
				formatStr,
				score.name,
				score.chineseScore,
				score.englishScore,
				score.mathScore,
				totalScore,
				avg,
			)
			fmt.Printf(
				formatStr,
				score.name,
				score.chineseScore,
				score.englishScore,
				score.mathScore,
				totalScore,
				avg,
			)
		}

		// 轉換成Byte型別
		data_byte := []byte(exportString)

		// 輸出檔案
		ioutil.WriteFile("ScoreBook.txt", data_byte, 0644)

		fmt.Println("執行完成")
	} else {
		showError("尚未匯入學生資料")
	}
}

// 顯示錯誤
func showError(message string) {
	fmt.Println("[Error]", message)
}

func main() {
	for {
		showMenu()
		switch selectFunc() {
		case 0:
			return
		case 1:
			importStudents()
			break
		case 2:
			importScores()
		case 3:
			exportScores()
		default:
			showError("無此選項")
		}
	}
}
