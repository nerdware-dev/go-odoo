package odoo

// AccountReportManager represents account.report.manager model.
type AccountReportManager struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	FinancialReportId *Many2One `xmlrpc:"financial_report_id,omitempty"`
	FootnotesIds      *Relation `xmlrpc:"footnotes_ids,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omitempty"`
	ReportName        *String   `xmlrpc:"report_name,omitempty"`
	Summary           *String   `xmlrpc:"summary,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReportManagers represents array of account.report.manager model.
type AccountReportManagers []AccountReportManager

// AccountReportManagerModel is the odoo model name.
const AccountReportManagerModel = "account.report.manager"

// Many2One convert AccountReportManager to *Many2One.
func (arm *AccountReportManager) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReportManager creates a new account.report.manager model and returns its id.
func (c *Client) CreateAccountReportManager(arm *AccountReportManager) (int64, error) {
	ids, err := c.CreateAccountReportManagers([]*AccountReportManager{arm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportManager creates a new account.report.manager model and returns its id.
func (c *Client) CreateAccountReportManagers(arms []*AccountReportManager) ([]int64, error) {
	var vv []interface{}
	for _, v := range arms {
		vv = append(vv, v)
	}
	return c.Create(AccountReportManagerModel, vv, nil)
}

// UpdateAccountReportManager updates an existing account.report.manager record.
func (c *Client) UpdateAccountReportManager(arm *AccountReportManager) error {
	return c.UpdateAccountReportManagers([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReportManagers updates existing account.report.manager records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReportManagers(ids []int64, arm *AccountReportManager) error {
	return c.Update(AccountReportManagerModel, ids, arm, nil)
}

// DeleteAccountReportManager deletes an existing account.report.manager record.
func (c *Client) DeleteAccountReportManager(id int64) error {
	return c.DeleteAccountReportManagers([]int64{id})
}

// DeleteAccountReportManagers deletes existing account.report.manager records.
func (c *Client) DeleteAccountReportManagers(ids []int64) error {
	return c.Delete(AccountReportManagerModel, ids)
}

// GetAccountReportManager gets account.report.manager existing record.
func (c *Client) GetAccountReportManager(id int64) (*AccountReportManager, error) {
	arms, err := c.GetAccountReportManagers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// GetAccountReportManagers gets account.report.manager existing records.
func (c *Client) GetAccountReportManagers(ids []int64) (*AccountReportManagers, error) {
	arms := &AccountReportManagers{}
	if err := c.Read(AccountReportManagerModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReportManager finds account.report.manager record by querying it with criteria.
func (c *Client) FindAccountReportManager(criteria *Criteria) (*AccountReportManager, error) {
	arms := &AccountReportManagers{}
	if err := c.SearchRead(AccountReportManagerModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// FindAccountReportManagers finds account.report.manager records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportManagers(criteria *Criteria, options *Options) (*AccountReportManagers, error) {
	arms := &AccountReportManagers{}
	if err := c.SearchRead(AccountReportManagerModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReportManagerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportManagerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportManagerModel, criteria, options)
}

// FindAccountReportManagerId finds record id by querying it with criteria.
func (c *Client) FindAccountReportManagerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportManagerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
