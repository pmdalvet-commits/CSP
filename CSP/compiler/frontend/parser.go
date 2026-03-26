package frontend

import "fmt"

type Node struct {
    Type string
    Value string
    Children []Node
}

func Parse(tokens []Token) Node {
    root := Node{Type: "Program", Value: "", Children: []Node{}}
    for _, t := range tokens {
        root.Children = append(root.Children, Node{Type: string(t.Type), Value: t.Value})
    }
    return root
}

func main() {
    code := `program { func main() { print("Hello, C#+!") } }`
    tokens := Lex(code)
    ast := Parse(tokens)
    fmt.Println(ast)
}
