// package for run core program
package main
// include other package 
import "fmt"
// func == function
// core function 
func main() {
// var [name] [type] = value
// s: value
	// var m map[type key] type var
	m := map[string]string{
		"hello":"world",
		"hello2":"world",
	}
	m["wingyplus"] = "maxbeef"
	fmt.Println(m)
}
