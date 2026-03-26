package main

import (
    "fmt"
    "CSP/compiler/frontend"
)

func main() {
    code := `program { func main() { print("Hello, C#+!") } }`
    tokens := frontend.Lex(code)
    ast := frontend.Parse(tokens)
    fmt.Println("AST gerada:", ast)
}
