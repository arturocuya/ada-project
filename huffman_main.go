package main

import (
	"fmt"
	huffman "./huffman"
	// consts "./consts"
)

func main(){
	values := []huffman.NodeData{'A','A','A','A', 'B','B','B', 'D', 'D', 'C'}
	frequencies := huffman.GetFrequencies(values)
	
	hf := huffman.NewHuffmanTree(frequencies)

	data := []huffman.NodeData{'C', 'A', 'B', 'D'}
	fmt.Println("Original data:",  data)
	encodedData := hf.EncodeData(data)
	fmt.Println("Encoded data:",  encodedData)
	fmt.Println("Decoded data:", hf.DecodeData(encodedData))
}