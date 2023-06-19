package main

import (
	"fmt"
	"path"
)

// path package provides utility functions for working with url path strings
func main() {
	var dir, file string
	dir, file = path.Split("go_basics/type.txt")

	fmt.Println("dir :", dir)
	fmt.Println("file :", file)

	// Output of the above is
	// dir : go_basics/
	// file : type.txt

	_, file = path.Split("go_basics/type.txt") // _ can be used to remove the output of dir

	fmt.Println("file :", file)
	// Output of the above is
	// file : type.txt
}

/*
func split(path string) (dir, file string)
{
	//dir output and file both are string
	// example if path=assets/agreement.pdf then dir = assets/ and file = agreement.pdf


}
*/
