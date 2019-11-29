package main2

import (
  "fmt"
  consts "./src/consts"
  cmp "./src/compress"
  dct "./src/compress/dct"
  rle "./src/compress/rle"
  ut "./src/utils"
)

func main() {
  var b = consts.Block {
    52, 55, 61, 66, 70, 61, 64, 73,
    63, 59, 55, 90, 109, 85, 69, 72,
    62, 59, 68, 113, 144, 104, 66, 73,
    63, 58, 71, 122, 154, 106, 70, 69,
    67, 61, 68, 104, 126, 88, 68, 70,
    79, 65, 60, 70, 77, 68, 58, 75,
    85, 71, 64, 59, 55, 61, 65, 83,
    87, 79, 69, 68, 65, 76, 78, 94,
  }

  bOg := b
  fmt.Println("Initial block")
  ut.PrintBlock(&b)

  cmp.ShiftBlock(&b)
  fmt.Println("Shifted block")
  ut.PrintBlock(&b)

  dct.Fdct(&b)
  fmt.Println("After fdct")
  ut.PrintBlock(&b)

  cmp.Quantize(&b)
  fmt.Println("After Quantization")
  ut.PrintBlock(&b)

  rleFormat := rle.RLE(&b)
  fmt.Println("In RLE format")
  ut.PrintRLE(rleFormat)

  fmt.Println("==============")
  fmt.Println("Reverse process")

  ib := rle.InvRLE(rleFormat)

  fmt.Println("From RLE back to block")
  ut.PrintBlock(&ib)

  cmp.InvQuantize(&ib)
  fmt.Println("Inverse Quantization")
  ut.PrintBlock(&ib)

  dct.Idct(&ib)
  fmt.Println("Idct")
  ut.PrintBlock(&ib)

  cmp.InvShiftBlock(&ib)
  fmt.Println("Inverse shift block")
  ut.PrintBlock(&ib)

  fmt.Println("Final comparison")
  fmt.Println("Original block")
  ut.PrintBlock(&bOg)
  fmt.Println("Decompressed block")
  ut.PrintBlock(&ib)
}
