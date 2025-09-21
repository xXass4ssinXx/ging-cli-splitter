package main

import (
	"fmt"
	"os"
	"wrap/block"
	"wrap/pad"
)

func main() {
	// op1, err := block.ParseOpcode("02", "OP_PUSH_2")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// }
	// int1, err := block.ParseInt16("97 00", "151")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// }
	// op2, err := block.ParseOpcode("1a", "OP_PUSH_26")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// }
	str1, err := block.ParseString("2f 70 6f 67 6f 6c 6f 20 2d 20 66 6f 73 73 20 69 73 20 66 72 65 65 64 6f 6d 2f", "/pogolo - foss is freedom/")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	str2, err := block.ParseString("2f 70 6f 67 6f 6c 6f 20 2d 20 66 6f 73 73 20 69 73 20 66 72 65 65 64 6f 6d 2f 2f 70 6f 67 6f 6c 6f 20 2d 20 66 6f 73 73 20 69 73 20 66 72 65 65 64 6f 6d 2f 2f 70 6f 67 6f 6c 6f 20 2d 20 66 6f 73 73 20 69 73 20 66 72 65 65 64 6f 6d 2f 2f 70 6f 67 6f 6c 6f 20 2d 20 66 6f 73 73 20 69 73 20 66 72 65 65 64 6f 6d 2f", "/pogolo - foss is freedom//pogolo - foss is freedom//pogolo - foss is freedom//pogolo - foss is freedom/")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	// op3, err := block.ParseOpcode("93", "OP_ADD")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// }
	// unk1, err := block.ParseUnknown("59 12 12 00 00 00 00", "[error]")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// }

	// blocks := []block.Block{&op1, &int1, &op2, &str1, &op3, &unk1}
	blocks := []block.Block{&str1, &str2}

	// Represents "minium padding"
	const PAD_AROUND = 1
	// Your algo will tell me the length of the current line, accounting for the separators between blocks
	const LENGTH_LINE_CURRENT = 60
	// Your algo will also tell me the max length of the line, probably subtracting an additional 2 for the end separators
	const LENGTH_LINE_MAX = 80
	splitBlock := blocks[1].SplitByLines(LENGTH_LINE_CURRENT, LENGTH_LINE_MAX-PAD_AROUND*2)
	var headersToPrint []string
	var dataToPrint []string
	for _, block := range splitBlock {
		headersToPrint = append(headersToPrint, block.Header())
		dataToPrint = append(dataToPrint, block.Data())
	}

	spaceRemaining := LENGTH_LINE_MAX - LENGTH_LINE_CURRENT
	for line := range headersToPrint {
		maxLength := max(len(headersToPrint[line]), len(dataToPrint[line]))
		paddedHeader := pad.PadAround(headersToPrint[line], maxLength+2)
		paddedData := pad.PadAround(dataToPrint[line], maxLength+2)
		fmt.Printf("---\nLine %d (%d remaining; header %d, data %d; padded header %d, padded data %d)\nheader: |%s|\ndata:   |%s|\npadded header: |%s|\npadded data:   |%s|\n",
			line, spaceRemaining,
			len(headersToPrint[line]), len(dataToPrint[line]),
			len(paddedHeader), len(paddedData),
			headersToPrint[line], dataToPrint[line],
			paddedHeader, paddedData,
		)
		spaceRemaining = LENGTH_LINE_MAX
	}
}
