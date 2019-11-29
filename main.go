package main

import (
    "os"
    "image"
    "log"
    "fmt"
    "strconv"
    ut "./src/utils"
    cmp "./src/compress"
    cspace "./src/colorspace"
    huffman "./src/compress/huffman"
)

func main() {
  // The image path must be passed as an argument
  if len(os.Args) == 1 {
		log.Fatalln("Image path and smooth value is required")
  } else if len(os.Args) == 2 {
  	log.Fatalln("Smooth value is required")
  }

  imgPath := os.Args[1]
  smooth, err := strconv.ParseFloat(os.Args[2], 64)
  ut.Check(err)

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
  imgMergedChannels := image.NewRGBA(rect)
  imgFromYcbcr := image.NewRGBA(rect)
  var compressedChannels [3]huffman.HfEncodedBlocks
  var imgDecompressedChannels [3]*image.RGBA

  // Convert to YCbCr
  cspace.ToYCbCr(img, imgYcbcr)
  ut.EncodeImg(imgYcbcr, ut.NewImgPath(imgPath, "1-ycbcr"))

  // Chroma Subsample
  cspace.ChromaSubsampling(imgYcbcr, imgSubsample)
  ut.EncodeImg(imgSubsample, ut.NewImgPath(imgPath, "2-subsample"))

  // Split YCbCr channels
  cspace.SplitChannelsYCbCr(imgYcbcr, channelsImg)
  
  for i:=0; i<3; i++ {
    ut.EncodeImg(channelsImg[i], ut.NewImgPath(imgPath, fmt.Sprintf("3-ycbcr-%d", i+1)))
  }

  // Compress each channel
  for i:=0; i<3; i++ {
    compressedChannels[i] = cmp.Compress(channelsImg[i], channelsImg[i].Bounds().Size(), smooth)
  }

  // Decompress each channel
  for i:=0; i<3; i++ {
    imgDecompressedChannels[i] = cmp.Decompress(compressedChannels[i])
    ut.EncodeImg(imgDecompressedChannels[i], ut.NewImgPath(imgPath, fmt.Sprintf("4-decompressed-%d", i+1)))
  }

  // Merge channels
	cspace.MergeChannelsYCbCr(imgDecompressedChannels, imgMergedChannels)
  ut.EncodeImg(imgMergedChannels, ut.NewImgPath(imgPath, "5-merged-channels"))

  // Convert to RGB
	cspace.ToRGB(imgMergedChannels, imgFromYcbcr)
  ut.EncodeImg(imgFromYcbcr, ut.NewImgPath(imgPath, "6-fromycbcr"))
}
