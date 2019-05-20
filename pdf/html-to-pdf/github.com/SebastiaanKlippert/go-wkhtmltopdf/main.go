package main

import (
	"fmt"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {

	ExampleNewPDFGenerator()
}

func ExampleNewPDFGenerator() {

	// Create new PDF generator
	wkhtmltopdf.SetPath("/Users/apple/workspace/go/src/bitbucket.org/gedw99/afira/client/flutter/pdf/print/github.com/SebastiaanKlippert/go-wkhtmltopdf/bin/wkhtmltox-0.12.5-1.macos-cocoa.pkg")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
