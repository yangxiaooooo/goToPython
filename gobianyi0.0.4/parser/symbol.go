package simple_parser

type Symbol struct {
	VariableName string
	Type         string
}

func NewSymbol(name, var_tyoe string) *Symbol {
	return &Symbol{
		VariableName: name,
		Type:         var_tyoe,
	}
}
