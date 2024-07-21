package main

import "fmt"

func main() {

	price := 0
	sabet := 1000000

	for i := 0; i < 1*12; i++ {
		price += sabet
		fmt.Println(price)
	}
}
