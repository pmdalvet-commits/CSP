#[derive(Debug)]
pub enum NodeType {
    Program,
    Func,
    Class,
    Print,
    Ident,
    Number,
    String,
}

#[derive(Debug)]
pub struct Node {
    pub node_type: NodeType,
    pub value: String,
    pub children: Vec<Node>,
}
