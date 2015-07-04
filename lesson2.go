
package main
import "fmt"

// Class contact {
//  private String name;
// }

type Contact struct {
	Name string
}

func main() {
	contact := Contact {
		Name:"wingyplus",
	}
	fmt.Println(contact);
}