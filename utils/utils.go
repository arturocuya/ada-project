package utils

import (
    "os"; "image"; "image/jpeg"; "image/color"
    "path/filepath"; "fmt"; "log"; "strings"
)

func Check(err error) {
  if err != nil {
    panic(err)
  }
}

func NewImgPath(originalPath string, suffix string) string {
  ext := filepath.Ext(originalPath)
  name := strings.TrimSuffix(filepath.Base(originalPath), ext)
  newPath := fmt.Sprintf("%s/%s_%s%s", filepath.Dir(originalPath), name,   suffix, ext)
  return newPath
}


func DecodeJpeg(path string) image.Image {
  f, err := os.Open(path)
  Check(err)
  defer f.Close()

  // img is readonly
  img, format, err := image.Decode(f)
  Check(err)
  if format != "jpeg" { log.Fatalln("Only jpeg images are supported") }

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

