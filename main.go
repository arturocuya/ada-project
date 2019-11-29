package main

import (
    "os"
    "image"
    "log"
    "fmt"
    ut "./src/utils"
    // cmp "./src/compress"
    cspace "./src/colorspace"
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
  imgMergedChannels := image.NewRGBA(rect)
  imgFromYcbcr := image.NewRGBA(rect)

  // Convert to YCbCr
  cspace.ToYCbCr(img, imgYcbcr)
  ut.EncodeJpeg(imgYcbcr, ut.NewImgPath(imgPath, "ycbcr"))

  // Chroma Subsample
  cspace.ChromaSubsampling(imgYcbcr, imgSubsample)
  //ut.EncodeJpeg(imgSubsample, ut.NewImgPath(imgPath, "subsample"))

  // Split YCbCr channels
  cspace.SplitChannelsYCbCr(imgYcbcr, channelsImg)
  cspace.MergeChannelsYCbCr(channelsImg, imgMergedChannels)

  for i:=0; i<3; i++ {
    ut.EncodeJpeg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("ycbcr-%d", i+1)))
  }
  ut.EncodeJpeg(imgMergedChannels, ut.NewImgPath(imgPath, "merged-channels"))

	cspace.ToRGB(imgMergedChannels, imgFromYcbcr)
  ut.EncodeJpeg(imgFromYcbcr, ut.NewImgPath(imgPath, "fromycbcr"))

  // compressed := cmp.Compress(channelsImg[0], channelsImg[0].Bounds().Size())
  // decompressed := cmp.Decompress(compressed)
  // ut.EncodeJpeg(decompressed, ut.NewImgPath(imgPath, "decomp"))
}
