package main

import (
	"fmt"
)

func forTes(){
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

func foreachTes(){
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	//  Jika tidak mau memanggil index, i ganti _
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// https://dasarpemrogramangolang.novalagung.com/A-array.html
}