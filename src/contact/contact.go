package contact

import "fmt"

type Contact struct {
	Name string `json:"name"`
	Tel string
	Email string
}

func Say(c Contact) string {
	return fmt.Sprintf("Hello %s",c.Name)
}
