package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[:3] ==", s[:3])
	fmt.Println("s[4:] ==", s[4:])
	fmt.Println("s[4:4] ==", s[4:4])

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] is %d\n", i, s[i])
	}
}
