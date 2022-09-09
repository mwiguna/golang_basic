package main

/**
	Masukkan file project ke gopath (mirip htdocs)

	Package harus sesuai dengan namafolder
	import package : namaprojectmodule/namahelper
	gunakan package : namahelper.namafunction
	tidak ada boleh func sama di satu package

	Penamaan fungsi dan variabel : 
	awali huruf besar = public ke package lain, 
	kecil = protected package yg sama aja (walau sudah import package)

	init() = consctrutor, gunakan import _ jika mau import tanpa mengakses fungsi di package tsb
*/

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func CheckHandler(w http.ResponseWriter, r* http.Request){
	// Function
	// fmt.Println(sumAll("Banyak jadi array : ", 15, 20, 30))
	// paramFunc("Biodata Diri", getBiodata)
}
  
func main()  {
	callRegex()
	fmt.Println("Server Running")
	//  --- Routing

	route := mux.NewRouter()
	route.HandleFunc("/", CheckHandler).Methods("GET")

	// Run server
	http.ListenAndServe("localhost:8000", route)
}
  
//  Go Run .