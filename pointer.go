package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func callPointerBasic() {
	// Variabel biasa
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}

	// * untuk inisialisasi bahwa ini tipe data pointer (var name). Value nya pake & 
	var address2 *Address = &address1
	var address3 *Address = &address1

	// perubahan satu persatu akan mempengaruhi semua variabel yang mengacu pointer tersebut (address1)
	address2.City = "Bandung"

	// perubahan semua nilai pointer, harus pakai * diawal
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(*address3) // * untuk mengambil value dari variable pointer (tanpa ampersand &)

	// Cara buat var langsung jadi tipe data pointer tanpa mengacu kemanapun (tidak digunakan pada case ini)
	var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	fmt.Println(address4)

	// Cara buat var langsung jadi tipe data pointer tapi kosong, pakai new
	var address5 *Address = new(Address)
	address5.City = "Jakarta"
	fmt.Println(address5)
}


// ------------ Pointer as parameter



func pointerAsParameter(){
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	var alamatPointer *Address = &alamat // jika tidak mau konvert dulu, bisa langsung passing &alamat ke ubahCountry(&alamat)
	ubahCountry(alamatPointer)
	fmt.Println(alamat)
}

func ubahCountry(address *Address){ // tambahkan *, maka param ini adalah pointer, bisa juga digunakan di method struct
	address.Country = "Indonesia"
}


