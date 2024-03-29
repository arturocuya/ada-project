package rle

import (
  //"image"
  //"image/color"
  //"fmt"
  consts "../../consts"
  ut "../../utils"
)

func RLE(b *consts.Block) consts.RLEList {
  rle := make(consts.RLEList, 0)

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
        var newTuple consts.RLETuple = consts.RLETuple{zeroCount, ut.BitSize(bValue), bValue}
        rle = append(rle, newTuple)
        zeroCount = 0
      }
    }
  }

  if zeroCount != 0{
    var lastTuple consts.RLETuple = consts.RLETuple{0, 0, 0}
    rle = append(rle, lastTuple)
  }

  return rle
}

func InvRLE(rle consts.RLEList) consts.Block {
  var b consts.Block
  rleIndex := 0
  zerosRight := 0

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

      if (rle[rleIndex].Val != 0) {
        if (zerosRight == 0) {
          bValue = rle[rleIndex].Val
          if rleIndex+1 < len(rle) {
            rleIndex++
            zerosRight = rle[rleIndex].Zb
          }
        } else {
          bValue = 0
          zerosRight--
        }
      } else {
        // Case where only zeros are left
        bValue = 0
      }

      if d&1 == 0 {
        b[(x+k)*8+y-k] = bValue
      } else {
        b[(y-k)*8+x+k] = bValue
      }
    }
  }

  return b
}
