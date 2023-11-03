package main

import "fmt"

func main() {
	var a int = 876
	var c bool = true
	for b := 2; b < a; b++ {
		if a%b == 0 {
			c = false
			fmt.Println(a, "是否为素数：", c)
			break
		} else {
			c = true
			fmt.Println(a, "是否为素数：", c)
		}
	}
}
