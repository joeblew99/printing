package main

import (
	"log"

	"github.com/sibosendteam/pdfmarker"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	pdfmarker.EnableLog = true
	imagick.Initialize()
	defer imagick.Terminate()
	src := "ref.pdf"
	dst := "ref-marked.pdf"
	xRes := float64(297) // unimplemented
	yRes := float64(210) // unimplemented
	outScale := 1.0
	opacity := .33
	size := pdfmarker.ImageIOSize{InResolution: &pdfmarker.Coordinate{X: xRes, Y: yRes}, OutScale: outScale}
	watermarkStyle := pdfmarker.WatermarkStyle{Opacity: opacity, Degrees: -20, Autofit: true}
	// use text as watermark
	textStyle := pdfmarker.TextStyle{Size: &pdfmarker.Coordinate{X: 640, Y: 104}, Font: "Arial", PointSize: 96, Weight: 1, Color: "grey", Opacity: opacity}
	watermark := pdfmarker.TextWatermark{Style: &watermarkStyle, TextStyle: &textStyle, Text: "example.com"}
	watermarkMW, _ := watermark.NewWatermark()
	// use image as watermark
	// watermark := &ImageWatermark{watermarkStyle, "example.com.png"}
	if err := pdfmarker.AddWatermark(src, dst, &size, watermarkMW, &watermarkStyle); err != nil {
		log.Fatal(err)
	}
}
