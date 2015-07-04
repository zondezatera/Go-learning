package main

import (
	"contact"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/contacts", addressBookHandler)
	http.ListenAndServe(":3000", nil)
}

func addressBookHandler(q http.ResponseWriter, r *http.Request) {
	contacts := []contact.Contact{}
	if r.Method == "POST" {
			var c contact.Contact
			err := json.NewDecoder(r.Body).Decode(&c)
			if err != nil {
				fmt.Println(err)
			}
			contacts = append(contacts, c)
			fmt.Println(contacts)
		}
}