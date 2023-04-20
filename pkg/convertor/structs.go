package convertor

import (
	"fmt"
	"strings"

	t "github.com/mehanizm/iuliia-go"
)

type DataCSV struct {
	ClientName        string `csv:"customer"`
	ClientEmail       string `csv:"email"`
	ClientAddress     string
	ClientCountry     string `csv:"country"`
	ClientStreetLine1 string `csv:"street_line_1"`
	ClientStreetLine2 string `csv:"street_line_2"`
	ClientCity        string `csv:"city"`
	ClientRegion      string `csv:"region"`
	ClientCode        string `csv:"postal_code"`

	Number      string `csv:"number"`
	Date        string `csv:"date"`
	PaymentDate string `csv:"payment_date"`
	Currency    string `csv:"accounting_currency"`
	TotalAmount string `csv:"total_amount"`
	NetAmount   string `csv:"net_amount"`
	VATAmount   string `csv:"tax_1_amount"`
	VAT         string `csv:"tax_1_rate"`
	VATFull     string
	VATShort    string
}

func (d *DataCSV) preRender() {
	d.ClientName = t.Wikipedia.Translate(d.ClientName)

	d.Date = strings.ReplaceAll(d.Date, "/", "-")
	d.PaymentDate = strings.ReplaceAll(d.PaymentDate, "/", "-")

	var address []string
	if d.ClientStreetLine1 != "" {
		address = append(address, d.ClientStreetLine1)
	}
	if d.ClientStreetLine2 != "" {
		address = append(address, d.ClientStreetLine2)
	}
	if d.ClientCity != "" {
		address = append(address, d.ClientCity)
	}
	if d.ClientRegion != "" {
		address = append(address, d.ClientRegion)
	}
	if d.ClientCode != "" {
		address = append(address, d.ClientCode)
	}

	d.ClientAddress = strings.Join(address, ", ")

	d.VATShort = strings.TrimSuffix(d.VAT, ".00%")
	d.VATFull = fmt.Sprintf("VAT (%s%%)", d.VATShort)
}

//
//document_type,
//date,
//number,
//description,
//original_currency,
//original_amount,
//accounting_currency,
//total_amount,
//gross_amount,
//discounts,
//net_amount,
//tax_1_amount,
//tax_1_rate,
//tax_1_country,
//tax_2_amount,
//tax_2_rate,
//tax_2_country,
//paid_amount,
//payment_method,
//payment_date,
//customer,
//country,
//region,
//postal_code,
//tax_id,
//processor,
//processor_id,
//processor_fee,
//due_date,
//po_number,
//hashtags,
//related_document_type,
//related_document_number,
//related_document_date,
//street_line_1,
//street_line_2,
//city,
//email,
//status,
//payment_processor,
//payment_processor_id
