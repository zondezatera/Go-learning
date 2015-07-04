
package main
import "fmt"
import "contact" // GOPATH/src/contact

func main() {
	c := contact.Contact {
		Name:"wingyplus",
	}
	fmt.Println(contact.Say(c));
}