
package main
import "fmt"

func main() {
	var n*int = new(int)
	*n = 1
	fmt.Println(*n)
}