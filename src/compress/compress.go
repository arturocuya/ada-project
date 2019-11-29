package compress

import (
  "image"
  "image/color"
  //ut "../utils"
  dct "./dct"
  rle "./rle"
  consts "../consts"
  huffman "./huffman"
)

var smooth float64 = 3

func ShiftBlock(b *consts.Block) {
  for x := 0; x < 8; x++{
    for y := 0; y < 8; y++{
      b[x + 8*y] -= 128
    }
  }
}

func InvShiftBlock(b *consts.Block) {
  for x := 0; x < 8; x++{
    for y := 0; y < 8; y++{
      b[x + 8*y] += 128
    }
  }
}

func Quantize(b *consts.Block) {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      b[8*i + j]= int32(float64(b[8*i + j]) / (consts.QuantizationTable[i][j]/smooth))
    }
  }
}

func InvQuantize(b *consts.Block) {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      b[8*i + j ]= int32(float64(b[8*i + j]) * (consts.QuantizationTable[i][j]/smooth))
    }
  }
}

func Compress(channel *image.RGBA, size image.Point) huffman.HfEncodedBlocks{
  numXBlocks := size.X / 8
  numYBlocks := size.Y / 8

  var encodedBlocks huffman.HfEncodedBlocks
  encodedBlocks.X = uint16(numXBlocks)
  encodedBlocks.Y = uint16(numYBlocks)

  for xBlock := 0; xBlock < numXBlocks; xBlock++{
    for yBlock := 0; yBlock < numYBlocks; yBlock++{
      // Initialize block
      var b consts.Block

      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          b[8*x+y] =  int32(channel.At(x + 8*xBlock, y + 8*yBlock).(color.RGBA).R)
        }
      }

      ShiftBlock(&b)
      dct.Fdct(&b)
      Quantize(&b)

      var hf *huffman.HuffmanTree
      var encodedData []consts.HuffmanEdge
      // Encode with Huffman
      rleList := rle.RLE(&b)
      if (len(rleList)<2){
      	hf = nil
      	encodedData = []consts.HuffmanEdge{}
      } else {
      	frequencies := huffman.GetFrequencies(rleList)
	      hf = huffman.NewHuffmanTree(frequencies)
	      encodedData = hf.EncodeData(rleList)
      }

      encodedBlocks.Blocks = append(encodedBlocks.Blocks, encodedData)
      encodedBlocks.HfTrees = append(encodedBlocks.HfTrees, hf)

      // Reconstruct image from block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          channel.Set(x + 8*xBlock, y + 8*yBlock, color.Gray{uint8(b[x+8*y])})
        }
      }
    }
  }
  return encodedBlocks
}

// Do not touch for other reason than complete destruction
func Decompress(encodedBlocks huffman.HfEncodedBlocks) *image.RGBA{
  rect := image.Rect(0, 0, int(encodedBlocks.X)*8, int(encodedBlocks.Y)*8)
  finalImg := image.NewRGBA(rect)

  for xBlock := 0; xBlock < int(encodedBlocks.X); xBlock++{
    for yBlock := 0; yBlock < int(encodedBlocks.Y); yBlock++{
      encodedBlock := encodedBlocks.Blocks[xBlock*int(encodedBlocks.Y) + yBlock]
      hf := encodedBlocks.HfTrees[xBlock*int(encodedBlocks.Y) + yBlock]

      var rleList []consts.RLETuple
      // Decode with huffman
      if hf == nil{
      	rleList = []consts.RLETuple{ consts.RLETuple{0,0,0} }
      } else{
      	rleList = hf.DecodeData(encodedBlock)
      }
      
      b := rle.InvRLE(rleList)

      InvQuantize(&b)
      dct.Idct(&b)
      InvShiftBlock(&b)

      // Reconstruct image from block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          finalImg.Set(x + 8*xBlock, y + 8*yBlock, color.Gray{uint8(b[8*x+y])})
        }
      }
    }
  }
  return finalImg
}
