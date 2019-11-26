package main

import (
    "os"; "image"; "log"; "fmt";
    ut "./utils"
)

func main() {
  // The image path must be passed as an argument
  if len(os.Args) < 2 { log.Fatalln("Image path is required") }
  imgPath := os.Args[1]
  img := ut.DecodeJpeg(imgPath)

  // Create empty images
  size := img.Bounds().Size()
  rect := image.Rect(0, 0, size.X, size.Y)
  imgYcbcr := image.NewRGBA(rect)

  var channelsImg [3]*image.RGBA
  for i:=0; i<3; i++ {
    channelsImg[i] = image.NewRGBA(rect)
  }

  // Convert to YCbCr
  ut.ToYCbCr(img, imgYcbcr)
  ut.EncodeJpeg(imgYcbcr, ut.NewImgPath(imgPath, "ycbcr"))

  // Split YCbCr channels
  ut.GetChannelsYCbCr(imgYcbcr, channelsImg)
  for i:=0; i<3; i++ {
    ut.EncodeJpeg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("ycbcr-%d", i+1)))
  }
}
