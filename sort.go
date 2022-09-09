package main

import(
	"fmt"
	"sort"
)

type Anggota struct {
	Name string
	Age  int
}

// Hanya buat alias
type AnggotaSlice []Anggota

// Fungsi yang harus dibuat, jadi ini mau filder berdasar umur terkecil
func (value AnggotaSlice) Len() int {
	return len(value)
}

func (value AnggotaSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value AnggotaSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func callSort() {
	Anggotas := []Anggota {
		{"Eko", 30},
		{"Jokok", 35},
		{"Budi", 31},
		{"Rudi", 25},
	}

	fmt.Println(Anggotas)
	sort.Sort(AnggotaSlice(Anggotas)) // ini sorting, kepanggil tiga method diatas, dan sort harus pake alias(data)
	fmt.Println(Anggotas)
}

