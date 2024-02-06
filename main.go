package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	a, _ := os.Open("./imgs/image1.png")
	img1, err := png.Decode(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := os.Open("./imgs/image2.png")
	img2, err := png.Decode(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	dst := image.NewRGBA(image.Rect(0, 0, img1.Bounds().Max.X, img1.Bounds().Max.Y))

	draw.Draw(dst, dst.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(dst, dst.Bounds(), img2, image.Point{0, 0}, draw.Src)

	buf := new(bytes.Buffer)
	png.Encode(buf, dst)

	fmt.Println("Ghép ảnh thành công!", buf.Bytes())
}
