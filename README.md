# Flutter PDF & Printing for Mobile and Desktop



## Status

Planning and getting approach together with code.

Please feel free to add any tips, info you can see or critique the approach. 

NOTE: repo is a mess right now :) Doing experiments.

## Pipeline

HTML --> FL Widgets --> PDF Widgets --> PDF --> Printer


# Notes 

Whats the big deal ?

For Mobile and Desktop.

This is the lib: https://github.com/DavBfr/dart_pdf

Its good but just needs a bit more:

Flutter
- Choose Printer before PDF output.
	- Is the wrong way around currently it seems.
- Parametrics for Flutter Widgets
	- When you make a flutter widget you have to provide the PDF Rendering for that Widget. The Standard Flutter Rendering tree is then used for Screen and Printing - its nice.
	- But Layout needs work. Orientation and Sizes, just like we have different Screen sizes, we have different Paper sizes.
- Make examples
	- Test it out.

Desktop Integration

- Use Google Cloud Connector lib. 
	- https://github.com/google/cloud-print-connector
	- It has the code to get a Desktops and RaspberryPI’s printer list, options and print spoolers

- Make a Printer dialogue exactly how Google Chrome does it. They had the same issue actually.
	- https://i.stack.imgur.com/pHFNs.png
- Flow:
	- User selects Printer and then options (paper size, etc). 
	- Code asks Flutter PDF to rerender. We display Render as a PDF → Raster (GhostScript golang plug can do this) in the right pane. And round and round it goes.
	- When User clicks Print, We pass the PDF to the local print spooler. 

There is a Half way house too
- Spit out the PDF and use local PDF Previewer. This sucks though because its spewing medical data onto disk !! Also its breaks the UX experience.


