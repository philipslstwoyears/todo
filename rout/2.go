package main

import (
	"fmt"
)

func main() {
	var num int
	var tests int
	fmt.Scan(&num, &tests)
	word := make(map[string]int, num)
	for i := 0; i < num; i++ {
		var abc string
		fmt.Scan(&abc)
		word[abc] += 1
	}
	for i := 0; i < tests; i++ {
		test := make(map[string]int)
		var str string
		fmt.Scan(&str)
		for _, i3 := range str {
			test[string(i3)] += 1
		}
		flag := len(word) == len(test)
		for key, value := range word {
			if value != test[key] {
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
