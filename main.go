package main

import (
	"fmt"

	"github.com/bumbimz/gotools"
	"github.com/bumbimz/gotools/stringhelper"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("Dog"))
	fmt.Println(stringhelper.Concat("Hello", "World", "i", "j", "k"))
}
