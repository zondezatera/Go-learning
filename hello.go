// package for run core program
package main
// include other package 
import "fmt"
// func == function
// core function 
func main() {
// var [name] [type] = value
// s: value
	var arr [5]int =  [5]int{1,2,3,4}
	for _, v:= range arr {
		fmt.Println(v)
	}
}
