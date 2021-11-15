/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */

package main

import (
	"fmt"
	"math"
	"strings"
)

func showMenu() {
	fmt.Println(strings.Repeat("=", 20), "圖形計算機", strings.Repeat("=", 20))
	fmt.Printf("%d\t%s\n", 1, "圓形")
	fmt.Printf("%d\t%s\n", 2, "三角形")
	fmt.Printf("%d\t%s\n", 3, "矩形")
	fmt.Printf("%d\t%s\n", 0, "離開")
	fmt.Println(strings.Repeat("=", 52))
}

var circle_radius int

func calculateCircle() {
	fmt.Print("請輸入「整數」半徑：")
	fmt.Scanln(&circle_radius)
	fmt.Printf("面積：%.2f\n", getCircleArea())
}

func getCircleArea() float64 {
	return math.Pow(float64(circle_radius), 2) * math.Pi
}

var triangle [2]int

func calculateTriangle() {
	fmt.Print("請輸入「整數」底：")
	fmt.Scanln(&triangle[0])
	fmt.Print("請輸入「整數」高：")
	fmt.Scanln(&triangle[1])
	fmt.Printf("面積：%.2f\n", getTriangleArea())
}

func getTriangleArea() float64 {
	return float64(triangle[0]) * float64(triangle[1]) / 2
}

var rectangle [2]int

func calculateRectangle() {
	fmt.Print("請輸入「整數」長：")
	fmt.Scanln(&rectangle[0])
	fmt.Print("請輸入「整數」寬：")
	fmt.Scanln(&rectangle[1])
	fmt.Printf("面積：%d\n", getRectangleArea())
}

func getRectangleArea() int {
	return rectangle[0] * rectangle[1]
}

func showError(message string) {
	fmt.Println("[Error]", message)
}

func main() {
	for {
		showMenu()
		var selected int
		fmt.Print("請輸入選項：")
		fmt.Scanln(&selected)
		fmt.Println()
		switch selected {
		case 0:
			return
		case 1:
			calculateCircle()
			break
		case 2:
			calculateTriangle()
			break
		case 3:
			calculateRectangle()
			break
		default:
			showError("無此選項")
		}
		fmt.Println()
	}
}
