package block

import (
	"fmt"
	"strconv"
)

var OP_CODES = map[uint8]string{
	0x02: "OP_PUSH_2",
	0x1a: "OP_PUSH_26",
	0x93: "OP_ADD",
}

type BlockOpcode struct {
	data uint8
}

// Assumes header is a hex string (1 byte) and data is a human-readable opcode
func ParseOpcode(header string, data string) (BlockOpcode, error) {
	if len(header) != 2 {
		return BlockOpcode{}, fmt.Errorf("header must be a single hex byte (2 hex chars), got: %q", header)
	}
	dataUint64, err := strconv.ParseUint(header, 16, 8)
	if err != nil {
		return BlockOpcode{}, fmt.Errorf("header is not a valid hex byte: %v", err)
	}
	dataUint8 := uint8(dataUint64)

	return BlockOpcode{
		data: dataUint8,
	}, nil
}

// Hex string (1 byte)
func (b *BlockOpcode) Header() string {
	return fmt.Sprintf("%02x", b.data)
}

func (b *BlockOpcode) HeaderLength() int {
	return 2
}

// Human-readable opcode
func (b *BlockOpcode) Data() string {
	data, ok := OP_CODES[b.data]
	if !ok {
		data = "OP_UNK"
	}
	return data
}

func (b *BlockOpcode) DataLength() int {
	return len(b.Data())
}

func (b *BlockOpcode) DataType() DataType {
	return DATATYPE_OPCODE
}
