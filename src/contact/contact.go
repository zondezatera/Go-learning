package contact

import "fmt"

type Contact struct {
	Name string `json:"name"`
	Tel string `json:"tel"`
	Email string `json:"email"`
}

func Say(c Contact) string {
	return fmt.Sprintf("Hello %s",c.Name)
}
