package frontend

import (
    "fmt"
    "strings"
)

type TokenType string

const (
    KEYWORD TokenType = "KEYWORD"
    IDENT   TokenType = "IDENT"
    STRING  TokenType = "STRING"
    NUMBER  TokenType = "NUMBER"
    SYMBOL  TokenType = "SYMBOL"
)

type Token struct {
    Type  TokenType
    Value string
}

func Lex(input string) []Token {
    words := strings.Fields(input)
    tokens := []Token{}

    for _, w := range words {
        switch w {
        case "program", "func", "class", "namespace", "print":
            tokens = append(tokens, Token{Type: KEYWORD, Value: w})
        case "{", "}", "(", ")", ";":
            tokens = append(tokens, Token{Type: SYMBOL, Value: w})
        default:
            tokens = append(tokens, Token{Type: IDENT, Value: w})
        }
    }
    return tokens
}

func main() {
    code := `program { func main() { print("Hello, C#+!") } }`
    tokens := Lex(code)
    for _, t := range tokens {
        fmt.Println(t)
    }
}
