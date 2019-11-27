package utils

import (
    "os"; "image"; "image/jpeg"; "image/color"
    "path/filepath"; "fmt"; "log"; "strings"; "../consts"
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

      componentY := originalColor.R
      componentCb := originalColor.G
      componentCr := originalColor.B

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

func ChromaSubsampling(originalImg image.Image, newImg *image.RGBA){
  /* This implementation corresponds to 4:2:2. It is only taking the average of
     TWO horizontal consecutive pixels
  */
  size := originalImg.Bounds().Size()
  oldImgColor := getColor(originalImg)
  const mcuWidth = 16 // MCU width
  const mcuHeight = 8 // MCU height

  // Iterate through each MCU in image 
 for y_mcu:=0; y_mcu<size.Y; y_mcu+=mcuHeight{
   for x_mcu:=0; x_mcu<size.X; x_mcu+=mcuWidth{
     // Iterate through MCU in image
     for y:=y_mcu; y<y_mcu+mcuHeight; y++{
       for x:=x_mcu; x<x_mcu+mcuWidth; x++{
         /* Replace the chromance of two consecutive horizontal pixels 
            for its average.

            In other words...
            Pixel A and pixel B are two consecutive horizontal pixels. Replace its
            chromance (Cb and Cr) for AVERAGE(A.cb, B.cb) and AVERAGE(A.cr, B.cr).

            The local index (index inside its MCU) of A is going to be even (0, 2, 3... )
            The local index (index inside its MCU) of B is going to be odd (1, 3, 5... )
         */
         var pixel [4]float64
         if x-x_mcu%2==0{
           // Get chromance average and put it in A
           pixel = oldImgColor(x, y)
           pixel[consts.Cb] = (oldImgColor(x*2, y)[consts.Cb] + oldImgColor(x*2+1, y)[consts.Cb])/2
           pixel[consts.Cr] = (oldImgColor(x*2, y)[consts.Cr] + oldImgColor(x*2+1, y)[consts.Cr])/2
         } else{
           // B just copies the chromance of the previous horizontal pixel (A)
           pixel = oldImgColor(x-1, y)
         }

         newImg.Set(x,y, color.RGBA {
           R: uint8(pixel[consts.Y]), G: uint8(pixel[consts.Cb]), B: uint8(pixel[consts.Cr]), A: uint8(pixel[3]),
         })
       }
     }
   }
  } 
}



func getColor(img image.Image) func(x, y int) [4]float64{
  return func(x, y int) [4]float64{
    pixel := img.At(x,y)
    intCol := color.RGBAModel.Convert(pixel).(color.RGBA)
    var floatColor [4]float64
    for i, intRGB := range [4]uint8{intCol.R, intCol.G, intCol.B, intCol.A}{
      floatColor[i] = float64(intRGB)
    }
    return floatColor
  }
}

