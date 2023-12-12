package resources

type CustomerResource struct {
	ID             int
	Customer_name  string
	Street         string
	City           string
	State          string
	Zip_code       string
	Account_number string
	Balance        string
	Agency         string
}

func (CustomerResource) TableName() string {
	return "customer"
}
