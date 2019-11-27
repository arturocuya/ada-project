package rle

import (
  "image"
  "image/color"
  "math"
  "fmt"
)

// TODO: Structure this so it holds blocks
type RLETuple struct{
  val uint8
  rep uint32
}

type RLEStruct []RLETuple

func RLE(channel *image.RGBA, size image.Point) /*RLEStruct*/ {
  //var quarterSize int = int(float64(size.X)*float64(size.Y)/4.0)
  rle := make(RLEStruct, 0)
  var prev uint8 = 255

  for line := 1; line <= (size.X + size.Y - 1); line++ {
    startCol := math.Max(0.0, float64(line - size.X))
    count := math.Min(float64(line), math.Min((float64(size.Y) - startCol), float64(size.Y)))

    for i := 0; i < int(count); i++ {
      pixel := channel.At(int(math.Min(float64(size.X), float64(line))) - i - 1, int(startCol) + i).(color.RGBA).R

      if prev != pixel {
        var newTuple RLETuple = RLETuple{pixel, 1}
        rle = append(rle, newTuple)
        prev = pixel
      } else {
        rle[len(rle) - 1].rep += 1
      }
    }
  }

  fmt.Printf("\n")
  fmt.Printf("Reduction: %f \n", 1.0 - float64(len(rle))/float64(size.X*size.Y))
  //return rle
}
