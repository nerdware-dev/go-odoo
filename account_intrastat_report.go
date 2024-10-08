package odoo

// AccountIntrastatReport represents account.intrastat.report model.
type AccountIntrastatReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountIntrastatReports represents array of account.intrastat.report model.
type AccountIntrastatReports []AccountIntrastatReport

// AccountIntrastatReportModel is the odoo model name.
const AccountIntrastatReportModel = "account.intrastat.report"

// Many2One convert AccountIntrastatReport to *Many2One.
func (air *AccountIntrastatReport) Many2One() *Many2One {
	return NewMany2One(air.Id.Get(), "")
}

// CreateAccountIntrastatReport creates a new account.intrastat.report model and returns its id.
func (c *Client) CreateAccountIntrastatReport(air *AccountIntrastatReport) (int64, error) {
	ids, err := c.CreateAccountIntrastatReports([]*AccountIntrastatReport{air})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountIntrastatReport creates a new account.intrastat.report model and returns its id.
func (c *Client) CreateAccountIntrastatReports(airs []*AccountIntrastatReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range airs {
		vv = append(vv, v)
	}
	return c.Create(AccountIntrastatReportModel, vv, nil)
}

// UpdateAccountIntrastatReport updates an existing account.intrastat.report record.
func (c *Client) UpdateAccountIntrastatReport(air *AccountIntrastatReport) error {
	return c.UpdateAccountIntrastatReports([]int64{air.Id.Get()}, air)
}

// UpdateAccountIntrastatReports updates existing account.intrastat.report records.
// All records (represented by ids) will be updated by air values.
func (c *Client) UpdateAccountIntrastatReports(ids []int64, air *AccountIntrastatReport) error {
	return c.Update(AccountIntrastatReportModel, ids, air, nil)
}

// DeleteAccountIntrastatReport deletes an existing account.intrastat.report record.
func (c *Client) DeleteAccountIntrastatReport(id int64) error {
	return c.DeleteAccountIntrastatReports([]int64{id})
}

// DeleteAccountIntrastatReports deletes existing account.intrastat.report records.
func (c *Client) DeleteAccountIntrastatReports(ids []int64) error {
	return c.Delete(AccountIntrastatReportModel, ids)
}

// GetAccountIntrastatReport gets account.intrastat.report existing record.
func (c *Client) GetAccountIntrastatReport(id int64) (*AccountIntrastatReport, error) {
	airs, err := c.GetAccountIntrastatReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*airs)[0]), nil
}

// GetAccountIntrastatReports gets account.intrastat.report existing records.
func (c *Client) GetAccountIntrastatReports(ids []int64) (*AccountIntrastatReports, error) {
	airs := &AccountIntrastatReports{}
	if err := c.Read(AccountIntrastatReportModel, ids, nil, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountIntrastatReport finds account.intrastat.report record by querying it with criteria.
func (c *Client) FindAccountIntrastatReport(criteria *Criteria) (*AccountIntrastatReport, error) {
	airs := &AccountIntrastatReports{}
	if err := c.SearchRead(AccountIntrastatReportModel, criteria, NewOptions().Limit(1), airs); err != nil {
		return nil, err
	}
	return &((*airs)[0]), nil
}

// FindAccountIntrastatReports finds account.intrastat.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountIntrastatReports(criteria *Criteria, options *Options) (*AccountIntrastatReports, error) {
	airs := &AccountIntrastatReports{}
	if err := c.SearchRead(AccountIntrastatReportModel, criteria, options, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountIntrastatReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountIntrastatReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountIntrastatReportModel, criteria, options)
}

// FindAccountIntrastatReportId finds record id by querying it with criteria.
func (c *Client) FindAccountIntrastatReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountIntrastatReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
