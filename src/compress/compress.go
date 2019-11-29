package compress

import (
  "image"
  "image/color"
  // "fmt"
  //ut "../utils"
  dct "./dct"
  rle "./rle"
  consts "../consts"
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

func Compress(channel *image.RGBA, size image.Point) consts.RLEBlocks{
  numXBlocks := size.X / 8
  numYBlocks := size.Y / 8

  var rleBlocks consts.RLEBlocks
  rleBlocks.X = uint16(numXBlocks)
  rleBlocks.Y = uint16(numYBlocks)

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

      rleBlocks.Blocks = append(rleBlocks.Blocks, rle.RLE(&b))

      // Reconstruct image from block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          channel.Set(x + 8*xBlock, y + 8*yBlock, color.Gray{uint8(b[x+8*y])})
        }
      }
    }
  }

  return rleBlocks
}

// Do not touch for other reason than complete destruction
func Decompress(rleBlocks consts.RLEBlocks) *image.RGBA{
  rect := image.Rect(0, 0, int(rleBlocks.X)*8, int(rleBlocks.Y)*8)
  finalImg := image.NewRGBA(rect)

  // fmt.Println(rleBlocks.X)
  // fmt.Println(rleBlocks.Y)

  for xBlock := 0; xBlock < int(rleBlocks.X); xBlock++{
    for yBlock := 0; yBlock < int(rleBlocks.Y); yBlock++{
      rleBlock := rleBlocks.Blocks[xBlock*int(rleBlocks.Y) + yBlock]
      b := rle.InvRLE(rleBlock)

      InvQuantize(&b)
      dct.Idct(&b)
      InvShiftBlock(&b)

      // Reconstruct image from block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          // fmt.Printf("Building pixel (%d, %d)\n", x + 8*xBlock, y + 8*yBlock)
          finalImg.Set(x + 8*xBlock, y + 8*yBlock, color.Gray{uint8(b[8*x+y])})
        }
      }
    }
  }
  return finalImg
}
