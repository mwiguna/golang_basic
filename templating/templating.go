package templating

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
)

//  ------------ HTML

var vHeader, vFooter = "views/template/header.html", "views/template/footer.html"

func Home(w http.ResponseWriter, r* http.Request){
	view := template.Must(template.ParseFiles(vHeader, vFooter, "views/home.html"))

	view.ExecuteTemplate(w, "header.html", nil)
	view.ExecuteTemplate(w, "home.html", nil)
	view.ExecuteTemplate(w, "footer.html", nil)
}
  

func main()  {
	fmt.Println("Server Running")

	//  --- Routing

	route := mux.NewRouter()
	route.HandleFunc("/home", Home).Methods("GET")

	// Run server
	http.ListenAndServe(":8000", route)
}
  