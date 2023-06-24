package main

import (
    "fmt"
    "os"
)

func main() {
    username := os.Args[1]
    password := os.Args[2]
    fmt.Println(len(os.Args))
    if len(os.Args)==3 {
        if username == "abc" && password == "123" {
            fmt.Println("Username and password are correct")
        } else if username == "abc" && password != "123" {
            fmt.Printf("User %q is denied access as the password entered is wrong", username)
        } else {
            fmt.Printf("User %q is denied access", username)
        }
    } else {
        fmt.Println("Usage: go run main.go <username> <password>")
    }
}