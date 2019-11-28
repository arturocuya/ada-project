package compress

import (
  /*"image"
  "image/color"
  "fmt"
  ut "../utils"
  fdct "./fdct"
  idct "./idct"
  rle "./rle"*/
  consts "../consts"
)

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
      b[8*i + j]= int32(float64(b[8*i + j]) / (consts.QuantizationTable[i][j]))
    }
  }
}

func InvQuantize(b *consts.Block) {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      b[8*i + j ]= int32(float64(b[8*i + j]) * (consts.QuantizationTable[i][j]))
    }
  }
}

/*
func Compress(channel *image.RGBA, size image.Point) consts.BlocksRLE{
  fmt.Println("COMPRESSING")
  fmt.Printf("\n")
  numXBlocks := size.X / 8
  numYBlocks := size.Y / 8

  rleBlocks := make(consts.BlocksRLE, 0)

  for xBlock := 0; xBlock < numXBlocks; xBlock++{
    for yBlock := 0; yBlock < numYBlocks; yBlock++{
      // Initialize block
      var b consts.Block

      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          b[x+8*y] =  int32(channel.At(x + 8*xBlock, y + 8*yBlock).(color.RGBA).R)
        }
      }

      bx := b
      shiftBlock(&b)
      sb := b
      fdct.Fdct(&b)
      fdctb := b
      quantize(&b)
      qb := b

      if xBlock == 0 && yBlock == xBlock {
        fmt.Println("Before shifting")
        ut.PrintBlock(&bx)
        fmt.Println("After shifting")
        ut.PrintBlock(&sb)
        fmt.Println("After fdct")
        ut.PrintBlock(&fdctb)
        fmt.Println("After quantizing")
        ut.PrintBlock(&qb)
      }

      rleBlocks = append(rleBlocks, rle.RLE(&b))

      // Reconstruct image from block
      for x := 0; x < 8; x++ {
        for y := 0; y < 8; y++{
          channel.Set(x + 8*xBlock, y + 8*yBlock,color.Gray{uint8(b[x+8*y])})
        }
      }
    }
  }

  return rleBlocks
}

func Decompress(rleBlocks consts.BlocksRLE) {
  fmt.Println("DECOMPRESSING")
  fmt.Printf("\n")

  for i := 0; i < len(rleBlocks); i++ {
    b := rle.InvRLE(rleBlocks[i])

    ibx := b

    invQuantize(&b)
    iqb := b
    idct.Idct(&b)
    idctb := b
    invShiftBlock(&b)
    isb := b

    if i == 0 {
      fmt.Println("Before")
      ut.PrintBlock(&ibx)
      fmt.Println("After inverting quantizing")
      ut.PrintBlock(&iqb)
      fmt.Println("After idct")
      ut.PrintBlock(&idctb)
      fmt.Println("After inverting shifting")
      ut.PrintBlock(&isb)
    }
  }
}
*/
