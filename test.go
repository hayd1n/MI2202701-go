/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞
 *   All rights reserved.
 *   CISH Robotics Team
 */
package main

import "fmt"

type test struct {
	name string
	age  int
}

func main() {
	var qwq = test{
		"qwq", 12,
	}
	fmt.Println(qwq.name, qwq.age)
}
