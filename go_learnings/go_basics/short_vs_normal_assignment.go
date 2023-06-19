package main

var version string // packaged scoped variable

func main() {
	// score:=0 // Dont use short declaration here

	var score int // already score =0 which is initialized

	var (
		video string

		duration int
		current  int
	)

	// SHORT DECLARATION

	// if you know the initial value declare short

	width, height := 20, 175

	width = 50 // assigns 50 to width

	color := "red" // new varibale: color

}

/*
o/p
$ go run short_vs_normal_assignment.go
# command-line-arguments
.\short_vs_normal_assignment.go:11:3: video declared and not used
.\short_vs_normal_assignment.go:13:3: duration declared and not used
.\short_vs_normal_assignment.go:14:3: current declared and not used
.\short_vs_normal_assignment.go:21:2: width declared and not used
.\short_vs_normal_assignment.go:21:9: height declared and not used
.\short_vs_normal_assignment.go:25:2: color declared and not used
*/
