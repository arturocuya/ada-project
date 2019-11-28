package main

import (
	"fmt"
	huffman "./huffman"
	consts "./consts"
)

func main(){
	sortedValues := []huffman.NodeData{'A', 'B', 'D', 'C'}
	// bitArray := []consts.HuffmanEdge{0,0,1,1,0,0,1,1,0,1,0,1,1,0,0}
	bitArray := []consts.HuffmanEdge{0,0,0,0,0,1,0,1,1} // C D B A
	hf := huffman.NewHuffmanTree(sortedValues)

	ans := hf.ReadBitArray(bitArray)
	for _, char := range ans{
		fmt.Printf("%c ", char)
	}
}