package main

import (
	"fmt"
)

func SpellComputers(num int) string {
	if num < 0 {
		return "invalid input"
	}

	n := num
	var add string = "компьютеров"
	// num digits length
	numlen := lenInt(n)

	if numlen > 2 {
		n = n - n/100*100
	}

	if n > 20 {
		n = n - n/10*10
	}
	switch n {
	case 1:
		add = "компьютер"
	case 2, 3, 4:
		add = "компьютера"
	}

	s := fmt.Sprintf("%d %s", num, add)
	return s
}

func lenInt(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}
