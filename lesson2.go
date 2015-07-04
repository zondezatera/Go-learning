
package main
import "fmt"

// Class contact {
//  public String name;
// }

type Contact struct {
	Name string
	Tel string
}

func main() {
	contact := &Contact {
		Name:"wingyplus",
	}
	contact.Tel = "telephone"
	fmt.Println(contact.Tel);
}