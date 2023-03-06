package main

import "github.com/fogleman/gg"

func main() {
	dc := gg.NewContext(512, 512)
	dc.SetRGB255(255, 222, 173)
	dc.Clear()
	dc.SavePNG("out.png")
}
