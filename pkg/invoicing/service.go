package invoicing

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
)

type InvoiceService struct{}

func (s InvoiceService) MakePDF(w io.Writer, invoice Invoice) error {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitMillimeter, gofpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, invoice.InvoiceNumber)

	err := pdf.Output(w)
	if err != nil {
		return fmt.Errorf("%T.Output(): %v", pdf, err)
	}
	return nil
}
