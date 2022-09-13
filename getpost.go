package main

import (
    "bytes"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
)


type Food struct {
	ID int
	Name string
	Category string
}


func GetHttp(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")

	// Get Response
	res, _ := http.Get("http://localhost:8001/foods")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Convert response to struct
	var foods []Food
	json.Unmarshal(body, &foods)

	// Write Response
	response, _ := json.Marshal(foods)
	w.Write(response)
}


func PostHttp(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")

	// Get Post to Byte Struct
	id, _ := strconv.Atoi(r.FormValue("ID"))
	postData, _ := json.Marshal(Food{
		ID: id,
		Name: r.FormValue("Name"),
		Category: r.FormValue("Category"),
	})

	// Send Post
	res, _ := http.Post("http://localhost:8001/foods", "application/json", bytes.NewBuffer(postData))
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Convert response to struct
	var foods Food
	json.Unmarshal(body, &foods)

	// Write Response
	response, _ := json.Marshal(foods)
	w.Write(response)
}