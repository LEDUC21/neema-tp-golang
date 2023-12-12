package models

type PaymentReceived struct {
	ID                        int
	Number                    string
	Date                      string `xorm:"default:2023-12-12"`
	Balance                   string
	Amount                    string
	Currency_rate             float64 `xorm:"default:1"`
	Fop                       string  `xorm:"default:cash"`
	Reference                 string  `xorm:"default:HLJDA86D"`
	Deducted_tax              bool    `xorm:"default:false"`
	Note                      string  `xorm:"default:null"`
	Used_amount               string  `xorm:"default:$0.00"`
	Status                    string  `xorm:"default:open"`
	Base_amount               string  `xorm:"default:$358,000.00"`
	Is_reconciled             bool    `xorm:"default:false"`
	Slug                      int64
	Id_customer               int64  `xorm:"default:372"`
	Id_chart_of_accounts_from int64  `xorm:"default:55"`
	Type                      string `xorm:"default:customer_payment"`
	Id_supplier               int64  `xorm:"default:9"`
	Id_consultant             int64  `xorm:"default:9"`
	Id_chart_of_accounts      int64  `xorm:"default:39"`
	Id_currency               int64  `xorm:"default:550"`
	Hidden_field              string `xorm:"default:null"`
	Transfert_type            string `xorm:"default:sales_without_invoices"`
	Already_used              int64  `xorm:"default:0"`
	Receipiant_name           string `xorm:"default:null"`
	Tag                       string `xorm:"default:1"`
}

func (PaymentReceived) TableName() string {
	return "payment_received"
}
