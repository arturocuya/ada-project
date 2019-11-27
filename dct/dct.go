package dct

import (
  "image"
  //"image/jpeg"
  "image/color"
  //"fmt"
  //"math"
  fdct "./fdct"
)

func shiftBlock(channel *image.RGBA, size image.Point) {
  for x := 0; x < size.X; x++{
    for y := 0; y < size.Y; y++{
      px := channel.At(x,y).(color.RGBA).R
      channel.Set(x,y,color.Gray{uint8(px - 128)})
    }
  }
}

func DCT(channel *image.RGBA, size image.Point) {
  shiftBlock(channel, size)

  numXBlocks := size.X / 8
  numYBlocks := size.Y / 8

  for xBlock := 0; xBlock < numXBlocks; xBlock++{
    for yBlock := 0; yBlock < numYBlocks; yBlock++{
      var b fdct.Block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          b[x+8*y] = int32(channel.At(x + 8*xBlock, y + 8*yBlock).(color.RGBA).R)
        }
      }

      fdct.Fdct(&b)

      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          channel.Set(x,y,color.Gray{uint8(b[x+8*y])})
        }
      }
    }
  }
}
