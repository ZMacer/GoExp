package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Printf("Hello %v\n", World)
	fmt.Println("Hello ", World)
	fmt.Println("Happy", Pi, "day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
