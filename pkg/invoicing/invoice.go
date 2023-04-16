package invoicing

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

const (
	invoiceDateFmt  = "2006-01-02"
	fromPlaceholder = `From Goldenhand Software
1 Electric Wharf, Generator Hall
Coventry, England, CV1 4JL
info@goldenhandosoftware.co.uk`
	toPlaceholder = `To Willy Wonka
Candy Factory, 1445 Norwood Ave
Itasca, IL 60143
willy@wonka.com`
)

type InvoiceRow struct {
	ID          int
	Description string
	Rate        string
	Quantity    string
	Amount      string
}

type Invoice struct {
	InvoiceNumber string
	Logo          []byte
	IssueDate     time.Time
	DueDate       time.Time
	From          string
	To            string
	Rows          []*InvoiceRow
	TotalAmount   string
	Symbol        string
	Notes         string
}

func FormToInvoice(form *multipart.Form) (*Invoice, error) {
	if form == nil {
		return nil, errors.New("multipart form can not be blank")
	}

	var i Invoice
	for name, values := range form.Value {
		if len(values) > 0 {
			val := values[0]
			if strings.TrimSpace(val) == "" && name != "notes" && name != "to" && name != "from" && name != "logo" {
				return nil, fmt.Errorf("please fill the necessary form field(%q)", name)
			}

			switch name {
			case "invoiceNumber":
				i.InvoiceNumber = val
			case "issueDate":
				issueDate, err := time.Parse(invoiceDateFmt, val)
				if err != nil {
					return nil, fmt.Errorf("issueDate(%q) has to be in format of %s", val, invoiceDateFmt)
				}
				i.IssueDate = issueDate
			case "dueDate":
				dueDate, err := time.Parse(invoiceDateFmt, val)
				if err != nil {
					return nil, fmt.Errorf("dueDate(%q) has to be in format of %s", val, invoiceDateFmt)
				}
				i.DueDate = dueDate
			case "from":
				if strings.TrimSpace(val) != "" {
					i.From = val
				} else {
					i.From = fromPlaceholder
				}
			case "to":
				if strings.TrimSpace(val) != "" {
					i.To = val
				} else {
					i.To = toPlaceholder
				}
			case "totalAmount":
				i.TotalAmount = val
			case "symbol":
				i.Symbol = val
			case "notes":
				i.Notes = val
			}

			if strings.HasPrefix(name, "description") {
				parts := strings.Split(name, "-")
				if len(parts) != 2 {
					return nil, fmt.Errorf("%q is not valid invoice row description", name)
				}
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("invoice row description(%q) doesn't contain ID", name)
				}
				row := findInvoiceRow(i.Rows, id)
				if row == nil {
					i.Rows = append(i.Rows, &InvoiceRow{ID: id, Description: val})
					continue
				}
				row.Description = val
				continue
			}
			if strings.HasPrefix(name, "rate") {
				parts := strings.Split(name, "-")
				if len(parts) != 2 {
					return nil, fmt.Errorf("%q is not valid invoice row rate", name)
				}
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("invoice row rate(%q) doesn't contain ID", name)
				}
				row := findInvoiceRow(i.Rows, id)
				if row == nil {
					i.Rows = append(i.Rows, &InvoiceRow{ID: id, Rate: val})
					continue
				}
				row.Rate = val
				continue
			}
			if strings.HasPrefix(name, "quantity") {
				parts := strings.Split(name, "-")
				if len(parts) != 2 {
					return nil, fmt.Errorf("%q is not valid invoice row quantity", name)
				}
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("invoice row quantity(%q) doesn't contain ID", name)
				}
				row := findInvoiceRow(i.Rows, id)
				if row == nil {
					i.Rows = append(i.Rows, &InvoiceRow{ID: id, Quantity: val})
					continue
				}
				row.Quantity = val
				continue
			}
			if strings.HasPrefix(name, "amount") {
				parts := strings.Split(name, "-")
				if len(parts) != 2 {
					return nil, fmt.Errorf("%q is not valid invoice row amount", name)
				}
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("invoice row amount(%q) doesn't contain ID", name)
				}
				row := findInvoiceRow(i.Rows, id)
				if row == nil {
					i.Rows = append(i.Rows, &InvoiceRow{ID: id, Amount: val})
					continue
				}
				row.Amount = val
			}
		}
	}

	for name, values := range form.File {
		if name == "logo" && len(values) > 0 {
			val := values[0]
			// filename should have an image extension here as validation, png, jpg, etc...
			// so I don't get a porn or something
			f, err := val.Open()
			if err != nil {
				return nil, fmt.Errorf("%T.Open(): %v", val, err)
			}

			raw, err := io.ReadAll(f)
			if err != nil {
				return nil, fmt.Errorf("io.ReadAll(): %v", err)
			}
			i.Logo = raw
		}
	}
	return &i, nil
}

func findInvoiceRow(rows []*InvoiceRow, id int) *InvoiceRow {
	for _, row := range rows {
		if row.ID == id {
			return row
		}
	}
	return nil
}
