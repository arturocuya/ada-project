package rle

import (
  "image"
  "image/color"
  "math"
  //"fmt"
  consts "../consts"
)

// TODO: Structure this so it holds blocks
type RLETuple struct{
  val uint8
  rep uint32
}

type RLEStruct []RLETuple

func RLE(b *consts.Block) RLEStruct {
  rle := make(RLEStruct, 0)
  var prev uint8 = 255

  for line := 1; line <= 15; line++ {
    startCol := math.Max(0.0, float64(line - 8))
    count := math.Min(float64(line), math.Min((8.0 - startCol), 8.0))

    for i := 0; i < int(count); i++ {
      pixel := b[(int(math.Min(8.0, float64(line))) - i - 1 + 8 * int(startCol) + i)]

      if prev != pixel {
        var newTuple RLETuple = RLETuple{pixel, 1}
        rle = append(rle, newTuple)
        prev = pixel
      } else {
        rle[len(rle) - 1].rep += 1
      }
    }
  }

  return rle
}
