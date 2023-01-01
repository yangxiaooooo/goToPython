package simple_parser

type NodeInterface interface {
	AddChild(child NodeInterface)
	GetChildren() []NodeInterface
	Attribute() string
}

type SyntaxNode struct {
	T        string
	children []NodeInterface
}

func NewSyntaxNode() *SyntaxNode {
	return &SyntaxNode{
		T: "",
	}
}

func (s *SyntaxNode) AddChild(node NodeInterface) {
	s.children = append(s.children, node)
}

func (s *SyntaxNode) GetChildren() []NodeInterface {
	return s.children
}

func (s *SyntaxNode) Attribute() string {
	if len(s.GetChildren()) == 0 {
		return s.T
	}

	attribute := ""
	for _, child := range s.children {
		attribute = attribute + child.Attribute()
	}

	attribute += s.T
	return attribute
}