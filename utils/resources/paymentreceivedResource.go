package resources

type PaymentReceivedResource struct {
	ID                        int               `json:"id" xorm:"'id' pk autoincr"`
	Number                    string            `json:"number" xorm:"'number'"`
	Date                      string            `json:"date" xorm:"'date'"`
	Balance                   string            `json:"balance" xorm:"'balance'"`
	Amount                    string            `json:"amount" xorm:"'amount'"`
	Currency_rate             float64           `json:"currency_rate" xorm:"'currency_rate'"`
	Fop                       string            `json:"fop" xorm:"'fop'"`
	Reference                 string            `json:"reference" xorm:"'reference'"`
	Deducted_tax              bool              `json:"deducted_tax" xorm:"'deducted_tax'"`
	Note                      string            `json:"note" xorm:"'note'"`
	Used_amount               string            `json:"used_amount" xorm:"'amount'"`
	Status                    string            `json:"status" xorm:"'status'"`
	Base_amount               string            `json:"base_amount" xorm:"'base_amount'"`
	Is_reconciled             bool              `json:"is_reconciled" xorm:"'is_reonciled'"`
	Slug                      int64             `json:"slug" xorm:"'slug'"`
	Id_customer               int64             `json:"id_customer" xorm:"'id_customer'"`
	Id_chart_of_accounts_from int64             `json:"id_chart_of_accounts_from" xorm:"'id_chart_of_accounts_from'"`
	Type                      string            `json:"type" xorm:"'type'"`
	Id_supplier               int64             `json:"id_supplier" xorm:"'id_supplier'"`
	Id_consultant             int64             `json:"id_consultant" xorm:"'id_consultant'"`
	Id_chart_of_accounts      int64             `json:"id_chart_of_accounts" xorm:"'id_chart_of_accounts'"`
	Id_currency               int64             `json:"id_currency" xorm:"'id_currency'"`
	Hidden_field              string            `json:"hidden_field" xorm:"'hidden_field'"`
	Transfert_type            string            `json:"transfert_type" xorm:"'transfert_type'"`
	Already_used              int64             `json:"already_used" xorm:"'already_used'"`
	Receipiant_name           string            `json:"receipiant_name" xorm:"'receipiant_name'"`
	Tag                       string            `json:"tag" xorm:"default:1"`
	CustomerResource          *CustomerResource `xorm:"'-"`
}

func (PaymentReceivedResource) TableName() string {
	return "payment_received"
}

type PaymentReceivedLastAmountApplyResource struct {
	ID                        int              `json:"id" xorm:"'id' pk autoincr"`
	Number                    string           `json:"number" xorm:"'number'"`
	Date                      string           `json:"date" xorm:"'date'"`
	Balance                   string           `json:"balance" xorm:"'balance'"`
	Amount                    string           `json:"amount" xorm:"'amount'"`
	Currency_rate             float64          `json:"currency_rate" xorm:"'currency_rate'"`
	Fop                       string           `json:"fop" xorm:"'fop'"`
	Reference                 string           `json:"reference" xorm:"'reference'"`
	Deducted_tax              bool             `json:"deducted_tax" xorm:"'deducted_tax'"`
	Note                      string           `json:"note" xorm:"'note'"`
	Used_amount               string           `json:"used_amount" xorm:"'amount'"`
	Status                    string           `json:"status" xorm:"'status'"`
	Base_amount               string           `json:"base_amount" xorm:"'base_amount'"`
	Is_reconciled             bool             `json:"is_reconciled" xorm:"'is_reonciled'"`
	Slug                      int64            `json:"slug" xorm:"'slug'"`
	Id_customer               int64            `json:"id_customer" xorm:"'id_customer'"`
	Id_chart_of_accounts_from int64            `json:"id_chart_of_accounts_from" xorm:"'id_chart_of_accounts_from'"`
	Type                      string           `json:"type" xorm:"'type'"`
	Id_supplier               int64            `json:"id_supplier" xorm:"'id_supplier'"`
	Id_consultant             int64            `json:"id_consultant" xorm:"'id_consultant'"`
	Id_chart_of_accounts      int64            `json:"id_chart_of_accounts" xorm:"'id_chart_of_accounts'"`
	Id_currency               int64            `json:"id_currency" xorm:"'id_currency'"`
	Hidden_field              string           `json:"hidden_field" xorm:"'hidden_field'"`
	Transfert_type            string           `json:"transfert_type" xorm:"'transfert_type'"`
	Already_used              int64            `json:"already_used" xorm:"'already_used'"`
	Receipiant_name           string           `json:"receipiant_name" xorm:"'receipiant_name'"`
	LastAmountApply           string           `json:"lastAmountApply" xorm:"'lastAmountApply'"`
	Tag                       string           `json:"tag" xorm:"default:1"`
	CustomerResource          CustomerResource `xorm:"'customer' null references(id) on delete cascade"`
}

func (PaymentReceivedLastAmountApplyResource) TableName() string {
	return "payment_received"
}
