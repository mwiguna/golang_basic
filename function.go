package main

import (
    "fmt"
    "strconv"
)

//var args (selalu diujung, banyak param bisa dijadikan array, atau masukkan slice/array langsung juga bisa)

func sumAll(judul string, numbers ...int) int {
	fmt.Print(judul)
	
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// Func parameter 

func paramFunc(judul string, cetakBiodata func(string, int) string){ // atau bisa buat type NamaAlias func(string int) string
	fmt.Println(judul)
	fmt.Println(cetakBiodata("Reni", 15))
}

func getBiodata(nama string, umur int) string {
	return nama + " " + strconv.Itoa(umur)
}


// Catatan
// defer namaFunction() tambahkan di line awal suatu function, maka ini akan selalu dipanggil setelah func selesai, misal log (destruct in function)
// panic == die
// recover() == menangkap parameter panic. **Tangkap di dalam func defer, karena kode fungsi dibawah panic itu di die. Tapi selain difungsi jadi undie

