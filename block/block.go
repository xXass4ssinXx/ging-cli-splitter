package block

type DataType int

const (
	DATATYPE_OPCODE DataType = iota
	DATATYPE_INT16
	DATATYPE_STRING
	DATATYPE_UNK
)

type Block interface {
	Header() string
	HeaderLength() int
	Data() string
	DataLength() int
	DataType() DataType
	SplitByLines(lengthLineCurrent int, lengthLineMax int) []Block
}
