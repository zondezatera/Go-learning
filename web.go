package main

import (
	"contact"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	contacts := []contact.Contact{}
	http.HandleFunc("/contacts", func(q http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var c contact.Contact
			err := json.NewDecoder(r.Body).Decode(&c)
			if err != nil {
				fmt.Println(err)
			}
			contacts = append(contacts, c)
			fmt.Println(contacts)
		} else {
			fmt.Println("Get Method")
		}
	})
	http.ListenAndServe(":3000", nil)
}
