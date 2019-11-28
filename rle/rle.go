package rle

import (
  //"image"
  //"image/color"
  //"fmt"
  consts "../consts"
  ut "../utils"
)

type RLETuple struct{
  // Zeros before the value
  Zb int
  Size uint8
  Val int32
}

type RLEList []RLETuple

func RLE(b *consts.Block) RLEList {
  rle := make(RLEList, 0)

  var zeroCount = 0

  // Zig zag iteration
  for d := 1; d < 16; d++ {
    x := d - 8
    if x < 0 {
      x = 0
    }
    y := d - 1
    if y > 7 {
      y = 7
    }
    j := 16 - d
    if j > d {
      j = d
    }
    for k := 0; k < j; k++ {
      var bValue int32
      if d&1 == 0 {
        bValue = b[(x+k)*8+y-k]
      } else {
        bValue = b[(y-k)*8+x+k]
      }

      if bValue == 0 {
        zeroCount++
      } else {
        var newTuple RLETuple = RLETuple{zeroCount, ut.BitSize(bValue), bValue}
        rle = append(rle, newTuple)
        zeroCount = 0
      }
    }
  }

  if zeroCount != 0{
    var lastTuple RLETuple = RLETuple{0, 0, 0}
    rle = append(rle, lastTuple)
  }

  return rle
}
