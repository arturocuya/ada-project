package consts

type YCbCr int

const (
  Y   YCbCr = 0
  Cb  YCbCr = 1
  Cr  YCbCr = 2
)

const blockSize = 64 // A DCT block is 8x8.
type Block [blockSize]int32
