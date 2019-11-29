package huffman

import (
  // "fmt"
	"../../consts"
	// "../rle"
	"sort"
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

type HfEncodedBlocks struct{
  Blocks [][]consts.HuffmanEdge
  HfTrees []*HuffmanTree
  X uint16
  Y uint16
}

type node struct {
	/* A new node can be created in three ways:
	   - Non-leaf node: Passing children to the newNonLeafNode constructor. Frequecy
	                    is calculated based on the children.
	   - Leaf node: Passing data and frequency to the newLeafNode constructor
	   - Any node: Children or data and frequency can (and must) be initialized later
	             with the default new keyword
	   In all constructors, the frequency value is also needed
	*/
	left      *node
	right     *node
	data      consts.RLETuple
	frequency int
}

func newLeafNode(data consts.RLETuple, freq int) *node {
	n := new(node)
	n.data = data
	n.frequency = freq
	return n
}

func newNonLeafNode(left, right *node) *node {
	n := new(node)
	n.left = left
	n.right = right
	n.frequency = left.frequency + right.frequency
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

/* For debugging
func (n *node) print() {
  fmt.Println("Freq: ", n.frequency)
  if (n.isLeaf()){
    fmt.Println("Data: ", n.data)
    fmt.Println("Is leaf")
    return
  } else{
    fmt.Println("Left: ")
    n.left.print()
    fmt.Println("----")
    fmt.Println("Right: ")
    n.right.print()
  }
}
*/

type HuffmanTree struct {
	root           *node
	addressTable map[consts.RLETuple][]consts.HuffmanEdge
}

/* For debugging
func (hf *HuffmanTree) GetFreqTable() (map[consts.RLETuple]int, []consts.RLETuple){
   return hf.frequencyTable, sortedKeysByVal(hf.frequencyTable)
}
*/

/* For debugging
func (hf *HuffmanTree) Print() (){
   hf.root.print()
}
*/


func NewHuffmanTree(dataFrequences map[consts.RLETuple]int) *HuffmanTree {
	/*  The constructor receives a map with the value to be stored
	    in tree as a key and the frequency of this value.
	    The frequencies are sorted in ascendent order and the most frequent
	    values are added to the right.
	    The address (directions or edges from the root to the value) to each
	    value is stored in the frequencyTable
	*/

	if len(dataFrequences) < 2 {
		panic("HuffmanTree needs a list of two or more elements")
	}

	hf := new(HuffmanTree)
	sd := sortedKeysByVal(dataFrequences) // sorted data by ascendent frequency

	firstLeft := newLeafNode(sd[0], dataFrequences[sd[0]])
	firstRight := newLeafNode(sd[1], dataFrequences[sd[1]])
	hf.root = newNonLeafNode(firstLeft, firstRight)

	for i := 2; i < len(sd); i++ {
		var left, right *node
    // fmt.Println("Adding ", sd[i], " with freq ", dataFrequences[sd[i]])
    // fmt.Println("Comparing to  ", hf.root.frequency)
		if dataFrequences[sd[i]] > hf.root.frequency { // right nodes are more frequent
			left = hf.root
			right = newLeafNode(sd[i], dataFrequences[sd[i]])
		} else {
			left = newLeafNode(sd[i], dataFrequences[sd[i]])
			right = hf.root
		}
		hf.root = newNonLeafNode(left, right) // create new root
	}

  hf.addressTable = make(map[consts.RLETuple][]consts.HuffmanEdge)
  hf.setAllAddresses(hf.root, []consts.HuffmanEdge{})

	return hf
}

func (hf *HuffmanTree) setAllAddresses(curNode *node, curAddress []consts.HuffmanEdge){
  address := []consts.HuffmanEdge{}
  address = append(address, curAddress...)
  
  if (curNode.isLeaf()){
    hf.addressTable[curNode.data] = address
    return
  } else{
    address = append(address, consts.LeftEdge)
    hf.setAllAddresses(curNode.left, address)

    address = address[:len(address)-1]

    address = append(address, consts.RightEdge)
    hf.setAllAddresses(curNode.right, address)
  }
}

func (hf *HuffmanTree) EncodeData(dataList []consts.RLETuple) []consts.HuffmanEdge {
	/* Encode the data with the addresses of the address table
	*/

	var encodedList = []consts.HuffmanEdge{}
	for _, data := range dataList {
    encodedData := hf.addressTable[data]
    if len(encodedData)==0{
      panic("Can't encode element outside the Huffman Tree address table")
    }
    encodedList = append(encodedList, encodedData...)
	}
	return encodedList
}

func (hf *HuffmanTree) DecodeData(bitArray []consts.HuffmanEdge) []consts.RLETuple {
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

	valuesFound := make([]consts.RLETuple, 0)
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
	if curNode != hf.root {
		panic("Can't read bitArray. Incorrect address to HuffmanTree value.")
	}
	return valuesFound
}

func sortedKeysByVal(m map[consts.RLETuple]int) []consts.RLETuple {
	sortedKeys := make([]consts.RLETuple, len(m))
	i := 0
	for key, _ := range m {
		sortedKeys[i] = key
		i++
	}
	sort.Slice(sortedKeys, func(left, right int) bool {
		return m[sortedKeys[left]] < m[sortedKeys[right]]
	})

	return sortedKeys
}

func GetFrequencies(rleList []consts.RLETuple) map[consts.RLETuple]int {
	freq := make(map[consts.RLETuple]int)

	for _, data := range rleList {
		if _, val := freq[data]; !val {
			freq[data] = 1
		} else {
			freq[data] += 1
		}
	}
	return freq
}
