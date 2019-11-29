package main

import (
    "os"
    "image"
    "log"
    "fmt"
    ut "./src/utils"
    cmp "./src/compress"
    cspace "./src/colorspace"
    huffman "./src/compress/huffman"
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
  var compressedChannels [3]huffman.HfEncodedBlocks
  var imgDecompressedChannels [3]*image.RGBA

  // Convert to YCbCr
  cspace.ToYCbCr(img, imgYcbcr)
  ut.EncodeJpeg(imgYcbcr, ut.NewImgPath(imgPath, "1-ycbcr"))

  // Chroma Subsample
  cspace.ChromaSubsampling(imgYcbcr, imgSubsample)
  ut.EncodeJpeg(imgSubsample, ut.NewImgPath(imgPath, "2-subsample"))

  // Split YCbCr channels
  cspace.SplitChannelsYCbCr(imgYcbcr, channelsImg)
  
  for i:=0; i<3; i++ {
    ut.EncodeJpeg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("3-ycbcr-%d", i+1)))
  }

  // Compress each channel
  for i:=0; i<3; i++ {
    compressedChannels[i] = cmp.Compress(channelsImg[i], channelsImg[i].Bounds().Size())
  }

  // Decompress each channel
  for i:=0; i<3; i++ {
    imgDecompressedChannels[i] = cmp.Decompress(compressedChannels[i])
    ut.EncodeJpeg(imgDecompressedChannels[i], ut.NewImgPath(imgPath, fmt.Sprintf("4-decompressed-%d", i+1)))
  }

  // Merge channels
	cspace.MergeChannelsYCbCr(imgDecompressedChannels, imgMergedChannels)
  ut.EncodeJpeg(imgMergedChannels, ut.NewImgPath(imgPath, "5-merged-channels"))

  // Convert to RGB
	cspace.ToRGB(imgMergedChannels, imgFromYcbcr)
  ut.EncodeJpeg(imgFromYcbcr, ut.NewImgPath(imgPath, "6-fromycbcr"))
}
