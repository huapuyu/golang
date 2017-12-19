package main

import "fmt"
import "golang/mylib"

const age = 23
var name string = "hello world"

type TTT struct {

}

/*
	hello world
 */
func main() {
	fmt.Printf("hello world")
	fmt.Printf("%d", mylib.Sum(1,2))
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
