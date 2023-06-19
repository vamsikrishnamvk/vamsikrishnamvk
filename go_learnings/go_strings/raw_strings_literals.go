package main

import "fmt"

func main() {
    var s string
    s= "how are you?"
    st := `how are you?`
    fmt.Println(s)
    fmt.Println(st)

    strLiteral := "<html>\n\t<body>\"hello\"</body>\n\n</html"
    fmt.Println(strLiteral)

    strRawLiteral := `
    <html>
        <body>"Hello"</body>
    </html>`
    fmt.Println(strRawLiteral)
}

/*

$ go run raw_strings_literals.go 
how are you?
how are you?
<html>
        <body>"hello"</body>

</html

    <html>
        <body>"Hello"</body>
    </html>

*/
