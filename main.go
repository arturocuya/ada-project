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

  imgSubsample := image.NewRGBA(rect)

  // Convert to YCbCr
  ut.ToYCbCr(img, imgYcbcr)
  //ut.EncodeJpeg(imgYcbcr, ut.NewImgPath(imgPath, "ycbcr"))

  // Chroma Subsample
  ut.ChromaSubsampling(imgYcbcr, imgSubsample)
  // ut.EncodeJpeg(imgSubsample, ut.NewImgPath(imgPath, "subsample"))

  // Split YCbCr channels
  ut.GetChannelsYCbCr(imgYcbcr, channelsImg)

  /* Save channels in file
  for i:=0; i<3; i++ {
    ut.EncodeJpeg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("ycbcr-%d", i+1)))
  }
  */

  dct.DCT(channelsImg[0], channelsImg[0].Bounds().Size())
  ut.EncodeJpeg(channelsImg[0], ut.NewImgPath(imgPath, fmt.Sprintf("dct-%d", 0)))

}
