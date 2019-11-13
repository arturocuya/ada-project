package main

import (
    "os"; "image"; "log"
    ut "./utils"
)

func main() {
  // The image path must be passed as an argument
  if len(os.Args) < 2 { log.Fatalln("Image path is required") }
  imgPath := os.Args[1]
  img := ut.DecodeJpeg(imgPath)

  size := img.Bounds().Size()
  rect := image.Rect(0, 0, size.X, size.Y)
  newImg := image.NewRGBA(rect)

  ut.ToBlackAndWhite(img, newImg, size)
  ut.EncodeJpeg(newImg, ut.NewImgPath(imgPath, "gray"))
}
