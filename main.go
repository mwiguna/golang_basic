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

func CallHandler(w http.ResponseWriter, r* http.Request){
	// Function
	// fmt.Println(sumAll("Banyak jadi array : ", 15, 20, 30))
	// paramFunc("Biodata Diri", getBiodata)
	callRegex()
}

func GetHttpHandler(w http.ResponseWriter, r* http.Request){
	GetHttp(w, r)
}

func PostHttpHandler(w http.ResponseWriter, r* http.Request){
	PostHttp(w, r)
}
  
func main()  {
	fmt.Println("Server Running")
	//  --- Routing

	route := mux.NewRouter()
	route.HandleFunc("/", CallHandler).Methods("GET")
	route.HandleFunc("/getHttp", GetHttpHandler).Methods("GET")
	route.HandleFunc("/postHttp", PostHttpHandler).Methods("POST")

	// Run server
	http.ListenAndServe("localhost:8000", route)
}
  
//  Go Run .