package main

import(
	"contact"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// (response,request)
	// POST /contacts
	// { "name" :"",tel: "",email:"" }
	// contacts := []contact.Contact{}
	http.HandleFunc("/contacts",func(q http.ResponseWriter,r *http.Request) {
		var c contact.Contact
		json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(q)
		})
	http.ListenAndServe(":3000",nil)
}