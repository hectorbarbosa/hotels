package main

import (
	"fmt"
	"strings"
)

func PrintTable(n int) {
	var sb strings.Builder
	var s string

	sb.WriteString("  ")
	for i := 1; i <= n; i++ {
		s = fmt.Sprintf("%3d", i)
		sb.WriteString(s)
	}
	fmt.Printf("%s\n", sb.String())
	sb.Reset()

	for i := 1; i <= n; i++ {
		sb.WriteString(fmt.Sprintf("%2d", i))
		for j := 1; j <= n; j++ {
			s = fmt.Sprintf("%3d", j*i)
			sb.WriteString(s)
		}
		fmt.Printf("%s\n", sb.String())
		sb.Reset()
	}
}

func main() {
	PrintTable(10)
}
