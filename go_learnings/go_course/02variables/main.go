package main

import "fmt"

var jwttoken = "dfgdfgdfg"

const LoginToken = "dfgdfgdf" // Public and it defined with capital letter in the start

const(
    website = "http://www.websits.com";
    website1 = "http://www.websits.com";
    website3 = "http://www.websits.com"
)

func main() {
    var username string = "vamsi"
    fmt.Println(username)
    fmt.Printf("variable is of type: %T \n", username)

    var smallVal float32 = 256.11111111111111111111111111111111111111111111111
    fmt.Println(smallVal)
    fmt.Printf("Variable is of type: %T \n", smallVal)

    // default values and some aliases
    var anotherval int
    fmt.Println(anotherval)
    fmt.Printf("Another variable is of type: %T \n", anotherval)

    //implicit type
    var website = "http://www.google.com"
    fmt.Println(website)

    // no var style
    numberOfUsers := 30000
    fmt.Println(numberOfUsers)

}