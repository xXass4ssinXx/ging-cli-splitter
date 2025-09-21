package block

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type BlockUnknown struct {
	header string
	data   string
}

// header and data can both be arbitrary strings, given that header is a space-separated hex string
func ParseUnknown(header string, data string) (BlockUnknown, error) {
	hexBytes := strings.Split(header, " ")
	for _, hexByte := range hexBytes {
		bytes, err := hex.DecodeString(hexByte)
		if err != nil {
			return BlockUnknown{}, err
		} else if len(bytes) != 1 {
			return BlockUnknown{}, fmt.Errorf("invalid hex byte in BlockUnknown header: '%s' (full header: '%s')", hexByte, header)
		}
	}

	return BlockUnknown{
		header: header,
		data:   data,
	}, nil
}

// Returns the original header
func (b *BlockUnknown) Header() string {
	return b.header
}

// Returns the original data string
func (b *BlockUnknown) Data() string {
	return b.data
}

func (b *BlockUnknown) DataType() DataType {
	return DATATYPE_UNK
}
