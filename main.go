package main

import (
    "os"
    "image"
    "log"
    //"fmt"
    ut "./src/utils"
     cmp "./src/compress"
    cspace "./src/colorspace"
)

func main() {
  // The image path must be passed as an argument
  if len(os.Args) < 2 { log.Fatalln("Image path is required") }
  imgPath := os.Args[1]
  img := ut.DecodeImg(imgPath)

  // Create empty images
  size := img.Bounds().Size()
  rect := image.Rect(0, 0, ut.Resize(size.X), ut.Resize(size.Y))
  imgYcbcr := image.NewRGBA(rect)

  var channelsImg [3]*image.RGBA
  for i:=0; i<3; i++ {
    channelsImg[i] = image.NewRGBA(rect)
  }

  imgSubsample := image.NewRGBA(rect)
  //imgMergedChannels := image.NewRGBA(rect)

  // Convert to YCbCr
  cspace.ToYCbCr(img, imgYcbcr)
  ut.EncodeImg(imgYcbcr, ut.NewImgPath(imgPath, "ycbcr"))

  // Chroma Subsample
  cspace.ChromaSubsampling(imgYcbcr, imgSubsample)
  //ut.EncodeImg(imgSubsample, ut.NewImgPath(imgPath, "subsample"))

  // Split YCbCr channels
  cspace.SplitChannelsYCbCr(imgYcbcr, channelsImg)
  //cspace.MergeChannelsYCbCr(channelsImg, imgMergedChannels)

  //for i:=0; i<3; i++ {
  //  ut.EncodeImg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("ycbcr-%d", i+1)))
  //}
  //ut.EncodeImg(imgMergedChannels, ut.NewImgPath(imgPath, "merged-channels"))

   compressed := cmp.Compress(channelsImg[0], channelsImg[0].Bounds().Size())
   decompressed := cmp.Decompress(compressed)
   ut.EncodeImg(decompressed, ut.NewImgPath(imgPath, "decomp"))
}
