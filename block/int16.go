package block

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type BlockInt16 struct {
	data int16
}

// Assumes header is a space-separated, little-endian hex string (2 bytes) and data is a decimal string
func ParseInt16(header string, data string) (BlockInt16, error) {
	header = strings.ReplaceAll(header, " ", "")
	if len(header) != 4 {
		return BlockInt16{}, fmt.Errorf("header must be 2 bytes (4 hex chars), got: %q", header)
	}
	bytes, err := hex.DecodeString(header)
	if err != nil {
		return BlockInt16{}, fmt.Errorf("header is not valid hex: %v", err)
	}
	val := int16(bytes[0]) | int16(bytes[1])<<8

	dataVal, err := strconv.ParseInt(data, 10, 16)
	if err != nil {
		return BlockInt16{}, fmt.Errorf("data is not a valid int16: %v", err)
	}
	if int16(dataVal) != val {
		return BlockInt16{}, fmt.Errorf("data (%d) does not match header value (%d)", dataVal, val)
	}

	return BlockInt16{data: val}, nil
}

// Space-separated hex string (2 bytes, little endian)
func (b *BlockInt16) Header() string {
	lo := byte(b.data & 0xff)
	hi := byte((b.data >> 8) & 0xff)
	return fmt.Sprintf("%02x %02x", lo, hi)
}

func (b *BlockInt16) HeaderLength() int {
	return 5
}

// Decimal string
func (b *BlockInt16) Data() string {
	return strconv.FormatInt(int64(b.data), 10)
}

func (b *BlockInt16) DataLength() int {
	return len(b.Data())
}

func (b *BlockInt16) DataType() DataType {
	return DATATYPE_INT16
}
