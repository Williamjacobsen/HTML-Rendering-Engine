package dom

// NodeType represents the type of a DOM node.
type NodeType int

const (
    ElementNode NodeType = iota
    TextNode
    DocumentNode
    // Add more node types as needed
)

// Node represents one element or text in the DOM tree.
type Node struct {
    Type     NodeType           // Element, Text, Document, etc.
    Tag      string             // "div", "p", "body" (empty for text nodes)
    Attrs    map[string]string  // attributes like id="hero"
    Text     string             // only for text nodes
    Children []*Node
    Parent   *Node
}