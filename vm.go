package struct_nostruct

type VM struct {
	sp       int
	variable map[string]int // VariableName: sp
	stack    []*Data
}
