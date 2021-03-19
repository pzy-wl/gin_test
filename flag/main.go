package main

import (
	"flag"
	"fmt"
)

var int = flag.Int("int", 1234, "help message for flagname")
var bool = flag.Bool("bool", false, "yes or no")
var string = flag.String("string", "nothing", "this is a string")

func main() {
	flag.Parse()
	fmt.Println("int:", *int)
	fmt.Println("bool:", *bool)
	fmt.Println("string:", *string)

}
