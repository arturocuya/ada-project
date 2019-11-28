package huffman

import (
	"../consts"
	// "../rle"
)

/*
type RLETuple struct{
  // Zeros before the value
  Zb int
  Size uint8
  Val int32
}
type RLEList []RLETuple
*/

type NodeData rune // rle.RLETuple
type node struct {
	/* A new node can be created in three ways:
	   - Non-leaf node: Passing children to the newNonLeafNode constructor
	   - Leaf node: Passing data to the newLeafNode constructor
	   - Any node: Children or data can (and must) be initialized later
	             with the default new keyword
	*/
	left  *node
	right *node
	data  NodeData
}

func newLeafNode(data NodeData) *node {
	n := new(node)
	n.data = data
	return n
}

func newNonLeafNode(left, right *node) *node {
	n := new(node)
	n.left = left
	n.right = right
	return n
}

func (n *node) isLeaf() bool {
	/* All huffman nodes have exactly two children or non.
	   This means that a left can not exist without a right
	   or viceversa. Therefore, to check if a node is a leaf
	   is trivial to look at the left or rigth. There is no
	   point in checking both
	*/
	return n.left == nil // also could be right==nil
}

type HuffmanTree struct {
	root *node
}

func NewHuffmanTree(sortedList []NodeData) *HuffmanTree {
	/* The relevance of the RLETuple (value) in the Huffman Tree is
	   defined by its proximity to the bottom right corner. Notice that
	   after the DCT this corner is the least important in the MCU.
	   The constructor receives a RLE List which should be already sorted
	   (in descending order of importance) after applying RLE in zig-zag
     to the MCU.
	*/

	if len(sortedList) < 2 {
		panic("HuffmanTree needs a list of one or more elements")
	}

	hf := new(HuffmanTree)

	hf.root = newNonLeafNode(newLeafNode(sortedList[len(sortedList)-1]),
		newLeafNode(sortedList[len(sortedList)-2]))
	for i := len(sortedList) - 3; i >= 0; i-- {
		left := hf.root
		right := newLeafNode(sortedList[i])
		hf.root = newNonLeafNode(left, right)
	}

  return hf
}

func (hf *HuffmanTree) ReadBitArray(bitArray []consts.HuffmanEdge) []NodeData {
	/* Huffman iterates through Bit Array, which represents a 'map'
	   with the directions (bit) to Huffman values (leafs). After
	   iterating, a list of the values found is returned. If at the
     end the current node is not a leaf, the bitArray is incorrect
     and an error is raised.

	   To navigate through a Huffman Tree:
	   1. Start with the root (parent).
	   2. Choose if going to the left or the right side (children). Going
	      to the left is represented as 0 and going to the right as 1.
	   3. If you reached a leaf, you found a value.
	      3.1. Add the value to the list of values found.
	      3.2. Return to the root.
	   4. Continue the search

	*/

	valuesFound := make([]NodeData, 0)
	curNode := hf.root
	for _, edge := range bitArray {
		if edge == consts.LeftEdge {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}

		if curNode.isLeaf() {
			valuesFound = append(valuesFound, curNode.data)
			curNode = hf.root // leaf reached, return to root
		}
	}
  if (curNode!=hf.root) { panic("Can't read bitArray. Incorrect address to HuffmanTree value.") }
	return valuesFound
}
