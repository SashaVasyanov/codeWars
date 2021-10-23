package centy

import "fmt"

func GivenCent(v int) int {
	cent := (v-1)/100 + 1

	fmt.Println(cent)
	return cent
}
