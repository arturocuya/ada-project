package utils

import (
  "os"; "image"; "image/jpeg"; "image/color"
  "path/filepath"; "fmt"; "log"; "strings"
  "math"
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

func Resize(size int) int {
  if (size % 8 != 0) {
    return int(math.Ceil(float64(size) / 8.0) * 8.0)
  }
  return size
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

func ToYCbCr(originalImg image.Image, newImg *image.RGBA){
  size := originalImg.Bounds().Size()

  for x := 0; x < size.X; x++{
    for y := 0; y < size.Y; y++{
      pixel := originalImg.At(x,y)
      originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

      r := float64(originalColor.R)
      g := float64(originalColor.G)
      b := float64(originalColor.B)

      componentY :=  uint8(r*0.299  + g*0.587  + b*0.114)
      componentCb := uint8(r*-0.169 + g*-0.331 + b*0.5 + 128)
      componentCr := uint8(r*0.5    + g*-0.419  + b*-0.081 + 128)

      newColor := color.RGBA {
        R: componentY, G: componentCb, B: componentCr, A: originalColor.A,
      }

      newImg.Set(x, y, newColor)
    }
  }
}

func GetChannelsYCbCr(originalImg image.Image, dividedImgs [3]*image.RGBA){
  size := originalImg.Bounds().Size()

  for x := 0; x < size.X; x++{
    for y := 0; y < size.Y; y++{
      pixel := originalImg.At(x,y)
      originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

      componentY := uint8(float64(originalColor.R))
      componentCb := uint8(float64(originalColor.G))
      componentCr := uint8(float64(originalColor.B))

      // Y channel
      dividedImgs[0].Set(x,y, color.RGBA {
        R: componentY, G: componentY, B: componentY, A: originalColor.A,
      })

      // Cb channel
      dividedImgs[1].Set(x,y, color.RGBA {
        R: componentCb, G: componentCb, B: componentCb, A: originalColor.A,
      })

      // Cr channel
      dividedImgs[2].Set(x,y, color.RGBA {
        R: componentCr, G: componentCr, B: componentCr, A: originalColor.A,
      })
    }
  }
}

func initCosineFormula(index int, size int) float64 {
  var r float64
  if index == 0 {
    r = 1 / math.Sqrt(float64(size))
  } else {
    r = math.Sqrt(2) / math.Sqrt(float64(size))
  }

  return r
}

func cosineFormula(i int, j int, size image.Point, value float64) uint8 {
  var ci, cj, dctl, sum float64

  ci = initCosineFormula(i, size.X)
  cj = initCosineFormula(j, size.Y)

  sum = 0

  for k := 0; k < size.X; k++{
    for l := 0; l < size.Y; l++{
      dctl = value /**
             math.Cos((2 * float64(k) + 1) * float64(i) * math.Pi / float64(2 * size.X)) *
             math.Cos((2 * float64(l) + 1) * float64(j) * math.Pi / float64(2 * size.Y))*/
      sum += dctl

    }
  }

  return (uint8) (ci * cj * sum)
}

func DiscreteCosineTransform(dividedImgs [3]*image.RGBA, size image.Point){
  for i := 0; i < 3; i++{
    fmt.Println("Aplying dct to ", i)
    for x := 0; x < size.X; x++{
      for y := 0; y < size.Y; y++{
        fmt.Println("Working on pixel",x,y)
        pixel := dividedImgs[i].At(x,y)
        originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

        r := cosineFormula(x, y, size, float64(originalColor.R))
        g := cosineFormula(x, y, size, float64(originalColor.G))
        b := cosineFormula(x, y, size, float64(originalColor.B))
        /*r := uint8(float64(originalColor.R))
        g := uint8(float64(originalColor.G))
        b := uint8(float64(originalColor.B))*/

        newColor := color.RGBA {
          R: r, G: g, B: b, A: originalColor.A,
        }

        dividedImgs[i].Set(x, y, newColor)
      }
    }
  }
}

