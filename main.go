package main

import (
    "os"
    "image"
    "image/color"
    "log"
    "fmt"
    //"reflect"
    ut "./utils"
    dct "./dct"
)

func main() {
  // The image path must be passed as an argument
  if len(os.Args) < 2 { log.Fatalln("Image path is required") }
  imgPath := os.Args[1]
  img := ut.DecodeJpeg(imgPath)

  // Create empty images
  size := img.Bounds().Size()
  rect := image.Rect(0, 0, ut.Resize(size.X), ut.Resize(size.Y))
  imgYcbcr := image.NewRGBA(rect)

  var channelsImg [3]*image.RGBA
  for i:=0; i<3; i++ {
    channelsImg[i] = image.NewRGBA(rect)
  }

  // Convert to YCbCr
  ut.ToYCbCr(img, imgYcbcr)
  //ut.EncodeJpeg(imgYcbcr, ut.NewImgPath(imgPath, "ycbcr"))

  // Split YCbCr channels
  ut.GetChannelsYCbCr(imgYcbcr, channelsImg)

  dct.DCT(channelsImg[0], channelsImg[0].Bounds().Size())
    ut.EncodeJpeg(channelsImg[0], ut.NewImgPath(imgPath, fmt.Sprintf("ycbcr-%d", 0)))

  for x:=0; x<size.X; x++{
    for y:=0; y<size.Y;y++{
      fmt.Printf("%d",channelsImg[0].At(x,y).(color.RGBA).R)
      fmt.Printf("\t")
    }
    fmt.Printf("\n")
  }

  /*for i:=0; i<3; i++ {
    ut.EncodeJpeg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("ycbcr-%d", i+1)))
  }

  for i:=0; i<3; i++ {
    ut.EncodeJpeg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("dct-%d", i+1)))
  }
 */
}
