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
	for k,v := range m {
		if k == "hello" {
			fmt.Println("Hello Gopher")
		}
		fmt.Println(k,v)
	}
}
