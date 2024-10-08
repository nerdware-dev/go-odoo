package odoo

// AccountAnalyticReport represents account.analytic.report model.
type AccountAnalyticReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountAnalyticReports represents array of account.analytic.report model.
type AccountAnalyticReports []AccountAnalyticReport

// AccountAnalyticReportModel is the odoo model name.
const AccountAnalyticReportModel = "account.analytic.report"

// Many2One convert AccountAnalyticReport to *Many2One.
func (aar *AccountAnalyticReport) Many2One() *Many2One {
	return NewMany2One(aar.Id.Get(), "")
}

// CreateAccountAnalyticReport creates a new account.analytic.report model and returns its id.
func (c *Client) CreateAccountAnalyticReport(aar *AccountAnalyticReport) (int64, error) {
	ids, err := c.CreateAccountAnalyticReports([]*AccountAnalyticReport{aar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticReport creates a new account.analytic.report model and returns its id.
func (c *Client) CreateAccountAnalyticReports(aars []*AccountAnalyticReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range aars {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticReportModel, vv, nil)
}

// UpdateAccountAnalyticReport updates an existing account.analytic.report record.
func (c *Client) UpdateAccountAnalyticReport(aar *AccountAnalyticReport) error {
	return c.UpdateAccountAnalyticReports([]int64{aar.Id.Get()}, aar)
}

// UpdateAccountAnalyticReports updates existing account.analytic.report records.
// All records (represented by ids) will be updated by aar values.
func (c *Client) UpdateAccountAnalyticReports(ids []int64, aar *AccountAnalyticReport) error {
	return c.Update(AccountAnalyticReportModel, ids, aar, nil)
}

// DeleteAccountAnalyticReport deletes an existing account.analytic.report record.
func (c *Client) DeleteAccountAnalyticReport(id int64) error {
	return c.DeleteAccountAnalyticReports([]int64{id})
}

// DeleteAccountAnalyticReports deletes existing account.analytic.report records.
func (c *Client) DeleteAccountAnalyticReports(ids []int64) error {
	return c.Delete(AccountAnalyticReportModel, ids)
}

// GetAccountAnalyticReport gets account.analytic.report existing record.
func (c *Client) GetAccountAnalyticReport(id int64) (*AccountAnalyticReport, error) {
	aars, err := c.GetAccountAnalyticReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aars)[0]), nil
}

// GetAccountAnalyticReports gets account.analytic.report existing records.
func (c *Client) GetAccountAnalyticReports(ids []int64) (*AccountAnalyticReports, error) {
	aars := &AccountAnalyticReports{}
	if err := c.Read(AccountAnalyticReportModel, ids, nil, aars); err != nil {
		return nil, err
	}
	return aars, nil
}

// FindAccountAnalyticReport finds account.analytic.report record by querying it with criteria.
func (c *Client) FindAccountAnalyticReport(criteria *Criteria) (*AccountAnalyticReport, error) {
	aars := &AccountAnalyticReports{}
	if err := c.SearchRead(AccountAnalyticReportModel, criteria, NewOptions().Limit(1), aars); err != nil {
		return nil, err
	}
	return &((*aars)[0]), nil
}

// FindAccountAnalyticReports finds account.analytic.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticReports(criteria *Criteria, options *Options) (*AccountAnalyticReports, error) {
	aars := &AccountAnalyticReports{}
	if err := c.SearchRead(AccountAnalyticReportModel, criteria, options, aars); err != nil {
		return nil, err
	}
	return aars, nil
}

// FindAccountAnalyticReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticReportModel, criteria, options)
}

// FindAccountAnalyticReportId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
