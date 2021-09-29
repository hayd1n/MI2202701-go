/*
 *   Copyright (c) 2021 CRT_HAO 張皓鈞 B11030202
 *   All rights reserved.
 *   NTUST
 */

// Target outputs
// ====================  台科大電影院  ====================
// 編號    項目   全票     學生    早晚    會員    愛心
// 01      2D    $ 280   $ 240   $ 200   $ 238   $ 140
// 02      3D    $ 340   $ 300   $ 260   $ 289   $ 170
// ======================================================

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Ticket type
	type ticket struct {
		name  string
		price int
	}

	// Tickets type array
	var tickets = [...]ticket{
		{"2D", 280},
		{"3D", 340},
	}

	// Title
	fmt.Println(strings.Repeat("=", 20), "台科大電影院", strings.Repeat("=", 20))

	// Tickets info title
	fmt.Printf("編號\t項目\t全票\t學生\t早晚\t會員\t愛心\t\n")

	// Tickets info
	for i, ticket := range tickets {
		fmt.Printf(
			"%02d\t%s\t$%4d\t$%4d\t$%4d\t$%4.0f\t$%4.0f\n", // format
			i+1,                        // 編號
			ticket.name,                // 票名稱
			ticket.price,               // 全票價
			ticket.price-40,            // 學生價
			ticket.price-80,            // 早晚價
			float64(ticket.price)*0.85, // 會員價
			float64(ticket.price)*0.5,  // 愛心價
		)
	}

	// Footer
	fmt.Println(strings.Repeat("=", 54))

}
