package struct_nostruct

type DataType int

const (
	Int DataType = iota
	String
)

type Data struct {
	Type DataType
	I    int
	S    string
}

func NewInt(i int) *Data {
	return &Data{
		Type: Int,
		I:    i,
		S:    "",
	}
}

func NewString(s string) *Data {
	return &Data{
		Type: String,
		I:    0,
		S:    s,
	}
}
