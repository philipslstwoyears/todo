package main

import "fmt"

func main() {
	var tests int
	fmt.Scan(&tests)
	for i := 0; i < tests; i++ {
		var str string
		fmt.Scan(&str)
		first := str[0]
		flag := true
		for indx, c := range str {
			if !(string(first) == string(c) || str[indx-1] == first && str[indx+1] == first) {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
