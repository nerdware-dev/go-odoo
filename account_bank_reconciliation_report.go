package odoo

// AccountBankReconciliationReport represents account.bank.reconciliation.report model.
type AccountBankReconciliationReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountBankReconciliationReports represents array of account.bank.reconciliation.report model.
type AccountBankReconciliationReports []AccountBankReconciliationReport

// AccountBankReconciliationReportModel is the odoo model name.
const AccountBankReconciliationReportModel = "account.bank.reconciliation.report"

// Many2One convert AccountBankReconciliationReport to *Many2One.
func (abrr *AccountBankReconciliationReport) Many2One() *Many2One {
	return NewMany2One(abrr.Id.Get(), "")
}

// CreateAccountBankReconciliationReport creates a new account.bank.reconciliation.report model and returns its id.
func (c *Client) CreateAccountBankReconciliationReport(abrr *AccountBankReconciliationReport) (int64, error) {
	ids, err := c.CreateAccountBankReconciliationReports([]*AccountBankReconciliationReport{abrr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankReconciliationReport creates a new account.bank.reconciliation.report model and returns its id.
func (c *Client) CreateAccountBankReconciliationReports(abrrs []*AccountBankReconciliationReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range abrrs {
		vv = append(vv, v)
	}
	return c.Create(AccountBankReconciliationReportModel, vv, nil)
}

// UpdateAccountBankReconciliationReport updates an existing account.bank.reconciliation.report record.
func (c *Client) UpdateAccountBankReconciliationReport(abrr *AccountBankReconciliationReport) error {
	return c.UpdateAccountBankReconciliationReports([]int64{abrr.Id.Get()}, abrr)
}

// UpdateAccountBankReconciliationReports updates existing account.bank.reconciliation.report records.
// All records (represented by ids) will be updated by abrr values.
func (c *Client) UpdateAccountBankReconciliationReports(ids []int64, abrr *AccountBankReconciliationReport) error {
	return c.Update(AccountBankReconciliationReportModel, ids, abrr, nil)
}

// DeleteAccountBankReconciliationReport deletes an existing account.bank.reconciliation.report record.
func (c *Client) DeleteAccountBankReconciliationReport(id int64) error {
	return c.DeleteAccountBankReconciliationReports([]int64{id})
}

// DeleteAccountBankReconciliationReports deletes existing account.bank.reconciliation.report records.
func (c *Client) DeleteAccountBankReconciliationReports(ids []int64) error {
	return c.Delete(AccountBankReconciliationReportModel, ids)
}

// GetAccountBankReconciliationReport gets account.bank.reconciliation.report existing record.
func (c *Client) GetAccountBankReconciliationReport(id int64) (*AccountBankReconciliationReport, error) {
	abrrs, err := c.GetAccountBankReconciliationReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abrrs)[0]), nil
}

// GetAccountBankReconciliationReports gets account.bank.reconciliation.report existing records.
func (c *Client) GetAccountBankReconciliationReports(ids []int64) (*AccountBankReconciliationReports, error) {
	abrrs := &AccountBankReconciliationReports{}
	if err := c.Read(AccountBankReconciliationReportModel, ids, nil, abrrs); err != nil {
		return nil, err
	}
	return abrrs, nil
}

// FindAccountBankReconciliationReport finds account.bank.reconciliation.report record by querying it with criteria.
func (c *Client) FindAccountBankReconciliationReport(criteria *Criteria) (*AccountBankReconciliationReport, error) {
	abrrs := &AccountBankReconciliationReports{}
	if err := c.SearchRead(AccountBankReconciliationReportModel, criteria, NewOptions().Limit(1), abrrs); err != nil {
		return nil, err
	}
	return &((*abrrs)[0]), nil
}

// FindAccountBankReconciliationReports finds account.bank.reconciliation.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankReconciliationReports(criteria *Criteria, options *Options) (*AccountBankReconciliationReports, error) {
	abrrs := &AccountBankReconciliationReports{}
	if err := c.SearchRead(AccountBankReconciliationReportModel, criteria, options, abrrs); err != nil {
		return nil, err
	}
	return abrrs, nil
}

// FindAccountBankReconciliationReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankReconciliationReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankReconciliationReportModel, criteria, options)
}

// FindAccountBankReconciliationReportId finds record id by querying it with criteria.
func (c *Client) FindAccountBankReconciliationReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankReconciliationReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
