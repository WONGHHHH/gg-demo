package main

import (
	"log"

	"github.com/fogleman/gg"
)

/**
转圆角
*/
func main() {
	im, err := gg.LoadImage("./baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG("out.png")
}
