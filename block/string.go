package block

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unicode/utf8"
)

type BlockString struct {
	data string
}

// Assumes `header` is a hex string containing the bytes of `data`
func ParseString(header string, data string) (BlockString, error) {
	headerClean := strings.ReplaceAll(header, " ", "")
	dataHex := hex.EncodeToString([]byte(data))
	if !strings.EqualFold(headerClean, dataHex) {
		return BlockString{}, fmt.Errorf("header (%q) does not match hex encoding of data (%q)", header, data)
	}

	return BlockString{data: data}, nil
}

// Space-separated hex string
func (b *BlockString) Header() string {
	hexStr := hex.EncodeToString([]byte(b.data))
	var sb strings.Builder
	for i := 0; i < len(hexStr); i += 2 {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(hexStr[i : i+2])
	}

	return sb.String()
}

func (b *BlockString) HeaderLength() int {
	return max(len(b.data)*3-1, 0)
}

// Original string data
func (b *BlockString) Data() string {
	return b.data
}

func (b *BlockString) DataLength() int {
	return utf8.RuneCountInString(b.data)
}

func (b *BlockString) DataType() DataType {
	return DATATYPE_STRING
}

func (b *BlockString) SplitByLines(lengthLineCurrent int, lengthLineMax int) []Block {
	// It is impossible for the header to be shorter than the data afaik, so we only need to compute by header
	var splitData []string
	data := []rune(b.data)
	// Remember that each rune costs 3 columns of header, unless it is the first (where it only costs 2)
	// I compensate for this by adding 1 to lengthLineMax.
	lengthLineMax += 1
	maxRunesCurrentLine := (lengthLineMax - lengthLineCurrent) / 3
	// Second term has lengthLineMax/3 added; for data of length 27 and line length 15, I'll want to create one slice at [0,15] and a second at [15,27] (clamping 30 to 27), so the loop needs to go "one past"
	for i := maxRunesCurrentLine; i < len(data)+lengthLineMax/3; i += lengthLineMax / 3 {
		dataStart := max(i-lengthLineMax/3, 0) // Previous iteration, or start of string
		dataEnd := min(i, len(data)-1)         // Current iteration, or end of string
		splitData = append(splitData, string(data[dataStart:dataEnd]))
	}

	var blocks []Block
	if len(splitData) == 0 {
		blocks = append(blocks, b)
	} else {
		for _, d := range splitData {
			blocks = append(blocks, &BlockString{
				data: d,
			})
		}
	}
	return blocks
}
