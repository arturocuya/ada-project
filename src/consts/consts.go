package consts

type YCbCr int

const (
  Y   YCbCr = 0
  Cb  YCbCr = 1
  Cr  YCbCr = 2
)

const blockSize = 64 // A DCT block is 8x8.
type Block [blockSize]int32

type RLETuple struct{
  // Zeros before the value
  Zb int
  Size uint8
  Val int32
}

type RLEList []RLETuple

type BlocksRLE []RLEList

// Psychovisually-tuned quantization table
// Extracted from https://es.coursera.org/lecture/dsp/7-6-the-jpeg-compression-algorithm-Q6hgv @ 2:45

// Note: Arrays can't be constant
// See https://stackoverflow.com/a/13140094

var QuantizationTable = [8][8] float64 {
  {16, 11, 10, 16, 24, 40, 51, 61},
  {12, 12, 14, 19, 26, 58, 60, 55},
  {14, 13, 16, 24, 40, 57, 69, 56},
  {14, 17, 22, 29, 51, 87, 80, 62},
  {18, 22, 37, 56, 68, 109, 103, 77},
  {24, 35, 55, 64, 81, 104, 113, 92},
  {49, 64, 78, 87, 103, 121, 120, 101},
  {72, 92, 95, 98, 112, 100, 103, 99},
}
