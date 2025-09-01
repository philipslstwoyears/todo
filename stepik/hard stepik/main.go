package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	uniq := make([]string, 0)
	povtor := make([]string, 0)
	var prev string
	n := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		str := scanner.Text()

		if str == "0" {
			break
		}

		if prev == str {
			povtor = append(povtor, str)
			n[str]++
			prev = str
		} else {
			n[str]++
			prev = str
		}
	}
	for m, i := range n {
		if i == 1 && m != "0" {
			uniq = append(uniq, m)
		}
	}
	fmt.Println("повторы", povtor)
	fmt.Println("уникальные", uniq)
	fmt.Println("кол-во", n)

}
