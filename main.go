package main

import (
    //"os"
    //"image"
    //"image/color"
    //"log"
    "fmt"
    //"reflect"
    //ut "./utils"
    //dct "./dct"
    //cspace "./colorspace"
    rle "./rle"
    consts "./consts"
)

func main() {
  var b = consts.Block {int32(100),int32(-60),0,6,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,int32(13),int32(-1),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,00,0,0,0,0,0,0,0}

  r := rle.RLE(&b)

  for i := 0; i < len(r); i++{
    fmt.Printf("[%d,%d] ", r[i].Zb, r[i].Val)
  }
  fmt.Printf("\n")

  /*
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
  cspace.ToYCbCr(img, imgYcbcr)
  //ut.EncodeJpeg(imgYcbcr, ut.NewImgPath(imgPath, "ycbcr"))

  // Chroma Subsample
  cspace.ChromaSubsampling(imgYcbcr, imgSubsample)
  //ut.EncodeJpeg(imgSubsample, ut.NewImgPath(imgPath, "subsample"))

  // Split YCbCr channels
  cspace.GetChannelsYCbCr(imgYcbcr, channelsImg)

  dct.Compress(channelsImg[0], channelsImg[0].Bounds().Size())
  ut.EncodeJpeg(channelsImg[0], ut.NewImgPath(imgPath, fmt.Sprintf("dct-%d", 0)))
  */
}
