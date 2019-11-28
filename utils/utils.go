package utils

import (
    "os"
    "image"
    "image/jpeg"
    "image/color"
    "path/filepath"
    "fmt"
    "strings"
    "math"
    consts "../consts"
)

func Check(err error) {
  if err != nil {
    panic(err)
  }
}

func Resize(size int) int {
  if (size % 8 != 0) {
    return int(math.Ceil(float64(size) / 8.0) * 8.0)
  }
  return size
}

func NewImgPath(originalPath string, suffix string) string {
  ext := filepath.Ext(originalPath)
  name := strings.TrimSuffix(filepath.Base(originalPath), ext)
  dirPath := filepath.Join(filepath.Dir(originalPath), fmt.Sprintf("%s-output", name))

  // Put output in originalDirPath/{name}-output/suffix.ext
  os.MkdirAll(dirPath, os.ModePerm)
  newPath := fmt.Sprintf("%s/%s%s", dirPath, suffix, ext)
  return newPath
}


func DecodeJpeg(path string) image.Image {
  f, err := os.Open(path)
  Check(err)
  defer f.Close()

  // img is readonly
  img, format, err := image.Decode(f)
  if format != "jpeg" { panic("Only jpeg images are supported") }
  Check(err)
  return img
}

func EncodeJpeg(img image.Image, path string) {
  fg, err := os.Create(path)
  defer fg.Close()
  Check(err)
  err = jpeg.Encode(fg, img, nil)
  Check(err)
}

func ToBlackAndWhite(originalImg image.Image, newImg *image.RGBA, size image.Point) {
  for x := 0; x < size.X; x++ {
    for y := 0; y < size.Y; y++ {
      pixel := originalImg.At(x,y)
      originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

      r := float64(originalColor.R)
      g := float64(originalColor.G)
      b := float64(originalColor.B)
      gray := uint8((r + g + b) / 3)

      newColor := color.RGBA {
        R: gray, G: gray, B: gray, A: originalColor.A,
      }

      newImg.Set(x, y, newColor)
    }
  }
}

func PrintBlock(b *consts.Block) {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      fmt.Printf("%d\t", b[i + 8 * j])
    }
    fmt.Printf("\n")
  }
}

func BitSize(val int32) uint8 {
  var b int32 = 1
  var count uint8 = 0
  val = int32(math.Abs(float64(val)))

  if val == 1 {
    return 1
  }

  for b < val {
    b = b << 1
    count++
  }

  return count
}
