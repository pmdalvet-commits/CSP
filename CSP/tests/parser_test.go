package tests

import (
    "testing"
    "CSP/compiler/frontend"
)

func TestParser(t *testing.T) {
    code := `program { func main() { print("Hello") } }`
    tokens := frontend.Lex(code)
    ast := frontend.Parse(tokens)
    if len(ast.Children) == 0 {
        t.Error("Parser não gerou AST")
    }
}
