// package for run core program
package main
// include other package 
import "fmt"
// func == function
// core function 
func main() {
// var [name] [type] = value
// s: value
	var s string =  "hello Golang"
	var b = s[0:5]
	// first letter uppercase meaning public var
	fmt.Println(len(b))
}
