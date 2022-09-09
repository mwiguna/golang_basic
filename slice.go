package main

import (
	"fmt"
)

func arrayBiasa(){
	var names = [...]string{"andi", "ferdi", "tiqa"} // ... bisa diisi panjang index

	fmt.Println(names)
}

func sliceTes(){
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)
}

func sliceMerubahSemua(){
	// Semua Slice dengan alamat yang sama akan berubah
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var bFruits = fruits[1:4]
	var baFruits = bFruits[0:1]

	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(baFruits) // [pinnaple]
}

func sliceBaru(){
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fruits = append(fruits, "baru")

	fmt.Println(fruits)
}


func mapBaru(){
	var bulans = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	bulans["maret"] = 30

	fmt.Println(bulans)

	// bisa delete, isExist
}


func sliceMapBaru(){
	var bulans = []map[string]int{
		{"januari":  50, "februari": 40},
		{"januari":  10, "februari": 30},
	}
	bulans[0]["februari"] = 15

	bulans = append(bulans, map[string]int{"januari":  3, "februari": 5})

	fmt.Println(bulans)
}