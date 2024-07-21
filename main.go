package main

import "fmt"

func main() {

	price := 0
	sabet := 1000000
	var sod int
	for i := 0; i <= 1*12; i++ {
		for i := 0; i <= 2*12; i++ {
			sod = (price / 100) * 1
			price += sod
			price += sabet

			fmt.Println(price, sod, i/12, i)
		}
	}
}
