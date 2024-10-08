package odoo

// AccountFinancialHtmlReport represents account.financial.html.report model.
type AccountFinancialHtmlReport struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omitempty"`
	Analytic             *Bool     `xmlrpc:"analytic,omitempty"`
	ApplicableFiltersIds *Relation `xmlrpc:"applicable_filters_ids,omitempty"`
	CompanyId            *Many2One `xmlrpc:"company_id,omitempty"`
	Comparison           *Bool     `xmlrpc:"comparison,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DateRange            *Bool     `xmlrpc:"date_range,omitempty"`
	DebitCredit          *Bool     `xmlrpc:"debit_credit,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	GeneratedMenuId      *Many2One `xmlrpc:"generated_menu_id,omitempty"`
	HierarchyOption      *Bool     `xmlrpc:"hierarchy_option,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	LineIds              *Relation `xmlrpc:"line_ids,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	ParentId             *Many2One `xmlrpc:"parent_id,omitempty"`
	ShowJournalFilter    *Bool     `xmlrpc:"show_journal_filter,omitempty"`
	TaxReport            *Bool     `xmlrpc:"tax_report,omitempty"`
	UnfoldAllFilter      *Bool     `xmlrpc:"unfold_all_filter,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFinancialHtmlReports represents array of account.financial.html.report model.
type AccountFinancialHtmlReports []AccountFinancialHtmlReport

// AccountFinancialHtmlReportModel is the odoo model name.
const AccountFinancialHtmlReportModel = "account.financial.html.report"

// Many2One convert AccountFinancialHtmlReport to *Many2One.
func (afhr *AccountFinancialHtmlReport) Many2One() *Many2One {
	return NewMany2One(afhr.Id.Get(), "")
}

// CreateAccountFinancialHtmlReport creates a new account.financial.html.report model and returns its id.
func (c *Client) CreateAccountFinancialHtmlReport(afhr *AccountFinancialHtmlReport) (int64, error) {
	ids, err := c.CreateAccountFinancialHtmlReports([]*AccountFinancialHtmlReport{afhr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFinancialHtmlReport creates a new account.financial.html.report model and returns its id.
func (c *Client) CreateAccountFinancialHtmlReports(afhrs []*AccountFinancialHtmlReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range afhrs {
		vv = append(vv, v)
	}
	return c.Create(AccountFinancialHtmlReportModel, vv, nil)
}

// UpdateAccountFinancialHtmlReport updates an existing account.financial.html.report record.
func (c *Client) UpdateAccountFinancialHtmlReport(afhr *AccountFinancialHtmlReport) error {
	return c.UpdateAccountFinancialHtmlReports([]int64{afhr.Id.Get()}, afhr)
}

// UpdateAccountFinancialHtmlReports updates existing account.financial.html.report records.
// All records (represented by ids) will be updated by afhr values.
func (c *Client) UpdateAccountFinancialHtmlReports(ids []int64, afhr *AccountFinancialHtmlReport) error {
	return c.Update(AccountFinancialHtmlReportModel, ids, afhr, nil)
}

// DeleteAccountFinancialHtmlReport deletes an existing account.financial.html.report record.
func (c *Client) DeleteAccountFinancialHtmlReport(id int64) error {
	return c.DeleteAccountFinancialHtmlReports([]int64{id})
}

// DeleteAccountFinancialHtmlReports deletes existing account.financial.html.report records.
func (c *Client) DeleteAccountFinancialHtmlReports(ids []int64) error {
	return c.Delete(AccountFinancialHtmlReportModel, ids)
}

// GetAccountFinancialHtmlReport gets account.financial.html.report existing record.
func (c *Client) GetAccountFinancialHtmlReport(id int64) (*AccountFinancialHtmlReport, error) {
	afhrs, err := c.GetAccountFinancialHtmlReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afhrs)[0]), nil
}

// GetAccountFinancialHtmlReports gets account.financial.html.report existing records.
func (c *Client) GetAccountFinancialHtmlReports(ids []int64) (*AccountFinancialHtmlReports, error) {
	afhrs := &AccountFinancialHtmlReports{}
	if err := c.Read(AccountFinancialHtmlReportModel, ids, nil, afhrs); err != nil {
		return nil, err
	}
	return afhrs, nil
}

// FindAccountFinancialHtmlReport finds account.financial.html.report record by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReport(criteria *Criteria) (*AccountFinancialHtmlReport, error) {
	afhrs := &AccountFinancialHtmlReports{}
	if err := c.SearchRead(AccountFinancialHtmlReportModel, criteria, NewOptions().Limit(1), afhrs); err != nil {
		return nil, err
	}
	return &((*afhrs)[0]), nil
}

// FindAccountFinancialHtmlReports finds account.financial.html.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReports(criteria *Criteria, options *Options) (*AccountFinancialHtmlReports, error) {
	afhrs := &AccountFinancialHtmlReports{}
	if err := c.SearchRead(AccountFinancialHtmlReportModel, criteria, options, afhrs); err != nil {
		return nil, err
	}
	return afhrs, nil
}

// FindAccountFinancialHtmlReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFinancialHtmlReportModel, criteria, options)
}

// FindAccountFinancialHtmlReportId finds record id by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFinancialHtmlReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
