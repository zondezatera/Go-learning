package main

import(
	"contact"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/contacts",func(q http.ResponseWriter,r *http.Request) {
		var c contact.Contact
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(c)
		})
	http.ListenAndServe(":3000",nil)
}