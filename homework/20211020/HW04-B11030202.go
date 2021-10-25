/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package main

import "fmt"

func main() {
	var students, courses int

	fmt.Print("請輸入學生人數：")
	fmt.Scanln(&students)
	fmt.Print("請輸入科目數量：")
	fmt.Scanln(&courses)

	score := make2DArray(students, courses)

	fmt.Println()
	fmt.Println("====== 成績輸入 =====")

	for i := 0; i < students; i++ {
		for j := 0; j < courses; j++ {
			for {
				fmt.Printf("請輸入第%d位學生的第%d成績：", i+1, j+1)
				fmt.Scanln(&score[i][j])
				if score[i][j] < 0 || score[i][j] > 100 {
					errorMessage("成績超出範圍(0~100)")
				} else {
					break
				}
			}
		}
		fmt.Println()
	}

	fmt.Println()

	fmt.Println("====== 判讀結果 =====")
	var max_student int
	var max_score int
	for i := range score {
		var student_score int = 0
		for j := range score[i] {
			student_score += score[i][j]
		}
		if student_score > max_score {
			max_student = i
			max_score = student_score
		}
	}
	fmt.Printf("第一名：同學%d(%d分)", max_student+1, max_score)
}

func make2DArray(a int, b int) [][]int {
	array := make([][]int, a)
	for i := range array {
		array[i] = make([]int, b)
	}
	return array
}

func errorMessage(msg string) {
	fmt.Printf("輸入錯誤。原因：%s\n\n", msg)
}
