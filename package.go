package main

import (
	"fmt"
	"strconv"
	"container/list"
	"container/ring"
	"time"
	"regexp"
)

func pkgStrconv(){
	umur := "15"
	umurInt, _ := strconv.ParseInt(umur, 10, 64) //param 2 itu bit type : biasa = 10, binary = 2, octal = 8. param 3 itu bit size 64bit. return ke 2 itu err, tapi _ aja
	fmt.Println(umurInt)
	fmt.Println(strconv.FormatInt(umurInt, 10)) //parse int = Itoa, format int = Atoi
}

func callList(){ // double linked list (array tanpa index, harus di foreach)
	datas := list.New()

	datas.PushBack(1)
	datas.PushBack(2)
	datas.PushBack(3)
	datas.PushFront(4)

	dataPertama := datas.Front().Value //Back() terakhir
	fmt.Println(dataPertama)
	
	dataMajuMundur := datas.Front().Next().Next().Prev().Value // index 0, 1, 2, 1
	fmt.Println(dataMajuMundur)

	// Semua
	for data := datas.Front(); data != nil; data = data.Next() {
		fmt.Println("Data " + strconv.Itoa(data.Value.(int)))
	}
}

func callRing(){
	datas := ring.New(3) // Length di tentukan, next & prev di ring akan kembali ke awal/akhir (lingkaran), jadi gak ada front dan back
	datas.Value = 1
	datas = datas.Next()
	datas.Value = 2 // nilai ditambah dgn cara seperti ini, sama seperti push
	datas = datas.Next()
	datas.Value = 3
	datas = datas.Next()

	datas.Do(func(value interface{}) {
		fmt.Println(value)
	})
}



func callTime(){
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" // == "yyyy-mm-dd, gokil formatnya"
	parse, _ := time.Parse(layout, "1990-03-20")
	fmt.Print(parse)
}


func callRegex(){
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", 2) // -1 unlimited
	fmt.Println(search)
}