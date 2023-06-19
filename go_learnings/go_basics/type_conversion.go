package main

import ( "fmt" )

func main() {
	speed := 100 // int
	force:= 2.5 // float
	// speed = speed * force // we get a compiler error two types
	// speed = speed * int(force) // converts 2.5 to 2 which is not expected
	speed = int(float64(speed) * force)
	fmt.Println(speed)

}