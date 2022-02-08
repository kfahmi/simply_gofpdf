package sgofpdf

import (
	"fmt"

	"gitlab.com/kfahmi/simple-gofpdf/logics"

	"github.com/jung-kurt/gofpdf"
)

type Init struct {
	Width          float64
	Heigth         float64
	OrientationStr string
	UnitStr        string
	FontDirStr     string
	PdfObj         *gofpdf.Fpdf
	FontPath       string
	sFPDFText      logics.ISPDFText
}

func main() {
	fmt.Println("Hello, World!")
}

//start docs init function sgofpdf
func (init *Init) InitSPDFDoc() error {
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
