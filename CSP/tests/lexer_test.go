package tests

import (
    "testing"
    "CSP/compiler/frontend"
)

func TestLexer(t *testing.T) {
    code := `program { func main() { print("Hello") } }`
    tokens := frontend.Lex(code)
    if len(tokens) == 0 {
        t.Error("Lexer não gerou tokens")
    }
}
