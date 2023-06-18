package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("go_basics/type.txt")

	fmt.Println("dir :", dir)
	fmt.Println("file :", file)
}
