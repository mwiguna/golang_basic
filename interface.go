package main

import (
	"errors"
	"fmt"
)

// Deklarasi interface (rancangan aja, ntah apa gunanya buat ini)
type HasName interface {
	GetName() string
}

// Penggunaan interface, tapi gak bisa dipanggil, jika tidak buat implemen fungsinya lagi
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// implemen fungsi dari penggunaan interface, sehingga sayhello otomatis jadi aktif. untuk struct person
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func callInterface() {
	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}



//  ------------- Interface kosong == dynamic


func Ups(i int) interface{} { // bisa return tipe data apapun
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func callInterfaceKosong(){
	var data interface{} = Ups(3)
	fmt.Println(data)

	//atau 
	var data_alternatif string = Ups(3).(string) // convert to string
	fmt.Println(data_alternatif)

	//jika ragu apa tipedatanya, if else aja Ups(3).(type) == string / int / bool
}





// ---------------- Error = object error itu bawaan, tinggal panggil

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else {
		result := nilai / pembagi
		return result, nil
	}
}

func callError() {
	hasil, err := Pembagi(100, 0)
	if err == nil{
		fmt.Println("Hasil", hasil)
	}else {
		fmt.Println("Error", err.Error())
	}
}