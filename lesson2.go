
package main
import "fmt"

func main() {
	var a int = 0
	var n*int = &a
	*n = 1
	fmt.Println(a)
}