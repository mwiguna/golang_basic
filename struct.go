package main

import (
	"fmt"
)

type User struct {
	nama, email string
	umur int
}


// Cara agar function bisa dipanggil oleh struct
func (user User) sayHello() {
	fmt.Println("Halo", user.nama)
}



func callStruct(){

	// -------- Ada berbagai cara mengisi

	tiqa := User{
		nama: "Tiqa",
		email: "tiqa@",
		umur: 15,
	}

	rian := User{"Rian", "rian@gmai.com", 18}

	var prabowo User
	prabowo.nama = "Prabowo"
	prabowo.email = "prabowo@gmail.com"
	prabowo.umur = 50

	fmt.Println(tiqa)
	fmt.Println(rian)
	fmt.Println(prabowo)

	//  ------ End Cara Mengisi

	// Struct memanggil fungsion sayHello
	tiqa.sayHello()
}

