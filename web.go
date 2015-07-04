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

func addressBookHandler(w http.ResponseWriter, r *http.Request) {
	contacts := []contact.Contact{}
	if r.Method == "POST" {
			var c contact.Contact
			err := json.NewDecoder(r.Body).Decode(&c)
			if err != nil {
				fmt.Println(err)
			}
			contacts = append(contacts, c)
			fmt.Println(contacts)
		} else if r.Method == "GET" {
			name := r.URL.Query().Get("name")
			for _,c := range contacts {
				if c.Name == name {
					json.NewEncoder(w).Encode(c)
					break
				}
			}
		}
}