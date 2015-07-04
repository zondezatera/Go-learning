
package main
import "fmt"

// Class contact {
//  public String name;
// }

type Contact struct {
	Name string
	Tel string
}

// funct name (input arugment) return type
func say(c Contact) string {
	return fmt.Sprintf("Hello %s",c.Name)
}

func main() {
	contact := Contact {
		Name:"wingyplus",
	}
	fmt.Println(say(contact));
}