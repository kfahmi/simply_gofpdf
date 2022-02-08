package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/kfahmi/simply_gofpdf/logics"
)

type Init struct {
	Width          float64
	Heigth         float64
	OrientationStr string
	UnitStr        string
	PdfObj         *gofpdf.Fpdf
	FontPath       string
	sFPDFText      logics.ISPDFText
}

// func main() {
// 	rootdir, _ := os.Getwd()
// 	fontPath := path.Join(rootdir, "font")
// 	pdf := &Init{
// 		Width:          800,
// 		Heigth:         800,
// 		OrientationStr: "P",
// 		UnitStr:        "mm",
// 		FontPath:       fontPath,
// 	}
// 	pdf.InitSPDFDoc()
// }

//start docs init function sgofpdf
func (init *Init) InitSPDFDoc() error {
	fmt.Println("HELLOW")
	init.PdfObj = gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: init.OrientationStr,
		UnitStr:        init.UnitStr,
		Size:           gofpdf.SizeType{Wd: init.Width, Ht: init.Heigth},
		FontDirStr:     init.FontPath,
	})
	init.sFPDFText = logics.InitSFPDFText(init.PdfObj)
	return nil
}

// Other Function
func (init *Init) RegisterFont(fontList []string) error {
	for _, fontType := range fontList {
		init.PdfObj.AddUTF8Font(fontType, "", fontType)
	}
	return nil
}

func (init *Init) WriteText(data interface{}) error {
	init.sFPDFText.CreateSFPDFText(data)
	return nil
}
