package main

import (
	"fmt"
	"github.com/markustenghamn/stringtotype"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println(stringtotype.GetType(arg))
}
