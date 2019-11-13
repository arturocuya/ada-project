package main

import (
    "os"; "image"; "image/jpeg"; "image/color"
    "path/filepath"; "fmt"; "log"; "strings"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func decodeJpeg(path string) image.Image {
  f, err := os.Open(path)
  check(err)
  defer f.Close()

  // img is readonly
  img, format, err := image.Decode(f)
  check(err)
  if format != "jpeg" { log.Fatalln("Only jpeg images are supported") }

  return img
}

func encodeJpeg(img image.Image, path string) {
  fg, err := os.Create(path)
  defer fg.Close()
  check(err)
  err = jpeg.Encode(fg, img, nil)
  check(err)
}

func newImgPath(originalPath string, suffix string) string {
  ext := filepath.Ext(originalPath)
  name := strings.TrimSuffix(filepath.Base(originalPath), ext)
  newPath := fmt.Sprintf("%s/%s_%s%s", filepath.Dir(originalPath), name, suffix, ext)
  return newPath
}

func toBlackAndWhite(originalImg image.Image, newImg *image.RGBA, size image.Point) {
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

func main() {
  // The image path must be passed as an argument
  if len(os.Args) < 2 { log.Fatalln("Image path is required") }
  imgPath := os.Args[1]
  img := decodeJpeg(imgPath)

  size := img.Bounds().Size()
  rect := image.Rect(0, 0, size.X, size.Y)
  newImg := image.NewRGBA(rect)

  toBlackAndWhite(img, newImg, size)
  encodeJpeg(newImg, newImgPath(imgPath, "gray"))
}
