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
	var m map[string]string = map[string]string{
		"hello":"world",
		"hello2":"world",
	}
	fmt.Println(m)
}
