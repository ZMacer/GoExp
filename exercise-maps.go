package main

import (
	"fmt"
	//"reflect"
	"strings"
)

func WordCount(s string) map[string]int {
	var v []string
	var res = make(map[string]int)
	v = strings.Split(s, " ")
	for i := 0; i < len(v); i++ {
		res[v[i]] += 1
	}
	return res
}

func main() {
	fmt.Println(WordCount("google is smart go go"))
}
