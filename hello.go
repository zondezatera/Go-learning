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
		switch {
			case k == "hello":
				fmt.Println("Hello Gopher")
				fallthrough
			case k == "hello2":
				fmt.Println("Hello")
				fallthrough
			default:
				fmt.Println("Default")
		}
		fmt.Println(k,v)
	}
}
