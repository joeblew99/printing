# pdf

SO pipeline is:

HTML --> FL Widgets --> FL PDF
- If not using HTML its all fine.


If there is a way to use zephyr to write any doc then great, but getting that to PDF using Flutter is not so easy.
There many be a need for using golang to make PDF's from PDF Templates
https://github.com/unidoc/unidoc-examples/tree/v3
- Looks very powerful now.
- zephyr is for generate Front end RTE, but maybe UniDoc can help with pre-setup templates and merging.

https://github.com/memspace/zefyr/pull/59
- Converts the Quill --> HTML. Ok this is great
- Now just need to write a QUILL --> PDF using the PDF lib. It MIGHT make sense to take the HTML and then convert that to PDF !!
	- Why because HTML will drive the whole thing later and so then we are sorted...
	- SO just write a HTML to PDF ( https://github.com/DavBfr/dart_pdf)


We want a unidirectional Rendering

HTML is per widget. Like Web Custom Elements
- Flutter ONYL gets the data and then turns it into pure Flutter

Printing is to take a screen shot and so its raster
- From the pure Flutter, we take a screen shot and embed it inside a PDF
- From there the Raster is injected inside a PDF, and ten send off to the Native PDF Preview for Printing

The other way is to write a PDF Rendered for each widget, instead of just taking a screen shot

-  We DONT have to compose many Widgets together if we use this approahc because for business apps you never want to do that.
	- Data View is a list, and you dont want any other shit on the screen. Same with Forms
	- Reporting View with lots of Charts and Tables on it is a single Widget. In fact you would be better off Rendering to PDF; and then showing that as the Flutter Widget
	- RTE, if a tough one. But the best way is to Write a PDF renderer that understands the Quill format. Or wimp out and take a screen shot.
	- Cal, i think its reasonable to write a PDF rendrer. Its likely that the PDF version looks different anyway.


The only issue is that the Print Format is not known until you get to the PDF Printing stage
- Instead you want to ask the Print size before hand, and render the Screenshot using:
	- The required DPI for the Paper size
	- The orientation ( Landscape or Portrait)
The Solution is to have a Flutter PDF Viewer, where the user selects the Paper format and based on that we re-render the PDF.
- It should be a matter of just making each Widgets PDF Rendered parametric for Landscape or portrait, and Paper size.
- We wont be able to know the potential Printer formats they have because we cant access that info easily, but thats ok, because:
	- the main thing is if its Portrait or Landscape
	- The Paper size chosen later wont matter because its Vector anyway and so will scale up fine.



---

golang libs

https://github.com/jung-kurt/gofpdf
- zero dependencies

what i need:
- convert to raster
- protect ( pass phrase) for sending to someone
- anoate and save
- templating ( for printing)
	- i will maybe have to parse FLUTTER OR HTML and then pump to the PDF:
	- or a screen shot

## Printing
Well thats the next challenge.

IF i have HTML.
1. 
Use Chrome embedded to do HTML to PDF
https://medium.com/compass-true-north/go-service-to-convert-web-pages-to-pdf-using-headless-chrome-5fd9ffbae1af
https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
- that will work !!

2. Use wkhtmltopdf
https://github.com/wkhtmltopdf/wkhtmltopdf/releases

golang wrapper
https://github.com/SebastiaanKlippert/go-wkhtmltopdf



## Process Manager
https://github.com/Unitech/pm2
- its nodeJS but seems to work !


