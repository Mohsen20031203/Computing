package main

import (
	"fmt"
)

func main() {

	price := 0
	var sod int
	sabet := 1000000
	for i := 0; i <= 2*12; i++ {
		sod = (price / 100) * 1
		price += sod
		if i%13 == 0 && i != 0 {

			sabet += (sabet / 100) * 40
		}
		price += sabet

		fmt.Println(price, sod, i/12, i)
	}
}
