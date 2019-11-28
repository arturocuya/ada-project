package dct

import (
  "image"
  //"image/jpeg"
  "image/color"
  //"fmt"
  //"math"
  fdct "./fdct"
  consts "../consts"
  //ut "../utils"
)

func shiftBlock(b *consts.Block) {
  for x := 0; x < 8; x++{
    for y := 0; y < 8; y++{
      b[x + 8*y] -= 128
    }
  }
}

func quantize(b *consts.Block) {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      newValue := int32(float64(b[i + 8*j]) / (consts.QuantizationTable[i][j]*8))

      if newValue < -255 {
        b[i + 8*j] = -255
      } else
      if newValue > 255 {
        b[i + 8*j] = 255
      } else {
        b[i + 8*j] = newValue
      }
    }
  }
}

func Compress(channel *image.RGBA, size image.Point) {
  numXBlocks := size.X / 8
  numYBlocks := size.Y / 8

  for xBlock := 0; xBlock < numXBlocks; xBlock++{
    for yBlock := 0; yBlock < numYBlocks; yBlock++{
      // Initialize block
      var b consts.Block

      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          b[x+8*y] = int32(channel.At(x + 8*xBlock, y + 8*yBlock).(color.RGBA).R)
        }
      }

      shiftBlock(&b)
      //fmt.Println("Raw block:")
      //ut.PrintBlock(&b)

      fdct.Fdct(&b)
      //fmt.Println("Block after dct")
      //ut.PrintBlock(&b)

      quantize(&b)
      //fmt.Println("Block after quantization")
      //ut.PrintBlock(&b)

      // Reconstruct image from block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          channel.Set(x + 8*xBlock, y + 8*yBlock,color.Gray{uint8(b[x+8*y])})
        }
      }
    }
  }
}
