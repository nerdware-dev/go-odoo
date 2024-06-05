package odoo

import (
	"fmt"
)

// AccountCashFlowReport represents account.cash.flow.report model.
type AccountCashFlowReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountCashFlowReports represents array of account.cash.flow.report model.
type AccountCashFlowReports []AccountCashFlowReport

// AccountCashFlowReportModel is the odoo model name.
const AccountCashFlowReportModel = "account.cash.flow.report"

// Many2One convert AccountCashFlowReport to *Many2One.
func (acfr *AccountCashFlowReport) Many2One() *Many2One {
	return NewMany2One(acfr.Id.Get(), "")
}

// CreateAccountCashFlowReport creates a new account.cash.flow.report model and returns its id.
func (c *Client) CreateAccountCashFlowReport(acfr *AccountCashFlowReport) (int64, error) {
	ids, err := c.CreateAccountCashFlowReports([]*AccountCashFlowReport{acfr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCashFlowReport creates a new account.cash.flow.report model and returns its id.
func (c *Client) CreateAccountCashFlowReports(acfrs []*AccountCashFlowReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range acfrs {
		vv = append(vv, v)
	}
	return c.Create(AccountCashFlowReportModel, vv)
}

// UpdateAccountCashFlowReport updates an existing account.cash.flow.report record.
func (c *Client) UpdateAccountCashFlowReport(acfr *AccountCashFlowReport) error {
	return c.UpdateAccountCashFlowReports([]int64{acfr.Id.Get()}, acfr)
}

// UpdateAccountCashFlowReports updates existing account.cash.flow.report records.
// All records (represented by ids) will be updated by acfr values.
func (c *Client) UpdateAccountCashFlowReports(ids []int64, acfr *AccountCashFlowReport) error {
	return c.Update(AccountCashFlowReportModel, ids, acfr)
}

// DeleteAccountCashFlowReport deletes an existing account.cash.flow.report record.
func (c *Client) DeleteAccountCashFlowReport(id int64) error {
	return c.DeleteAccountCashFlowReports([]int64{id})
}

// DeleteAccountCashFlowReports deletes existing account.cash.flow.report records.
func (c *Client) DeleteAccountCashFlowReports(ids []int64) error {
	return c.Delete(AccountCashFlowReportModel, ids)
}

// GetAccountCashFlowReport gets account.cash.flow.report existing record.
func (c *Client) GetAccountCashFlowReport(id int64) (*AccountCashFlowReport, error) {
	acfrs, err := c.GetAccountCashFlowReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if acfrs != nil && len(*acfrs) > 0 {
		return &((*acfrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.cash.flow.report not found", id)
}

// GetAccountCashFlowReports gets account.cash.flow.report existing records.
func (c *Client) GetAccountCashFlowReports(ids []int64) (*AccountCashFlowReports, error) {
	acfrs := &AccountCashFlowReports{}
	if err := c.Read(AccountCashFlowReportModel, ids, nil, acfrs); err != nil {
		return nil, err
	}
	return acfrs, nil
}

// FindAccountCashFlowReport finds account.cash.flow.report record by querying it with criteria.
func (c *Client) FindAccountCashFlowReport(criteria *Criteria) (*AccountCashFlowReport, error) {
	acfrs := &AccountCashFlowReports{}
	if err := c.SearchRead(AccountCashFlowReportModel, criteria, NewOptions().Limit(1), acfrs); err != nil {
		return nil, err
	}
	if acfrs != nil && len(*acfrs) > 0 {
		return &((*acfrs)[0]), nil
	}
	return nil, fmt.Errorf("account.cash.flow.report was not found with criteria %v", criteria)
}

// FindAccountCashFlowReports finds account.cash.flow.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashFlowReports(criteria *Criteria, options *Options) (*AccountCashFlowReports, error) {
	acfrs := &AccountCashFlowReports{}
	if err := c.SearchRead(AccountCashFlowReportModel, criteria, options, acfrs); err != nil {
		return nil, err
	}
	return acfrs, nil
}

// FindAccountCashFlowReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashFlowReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountCashFlowReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountCashFlowReportId finds record id by querying it with criteria.
func (c *Client) FindAccountCashFlowReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCashFlowReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.cash.flow.report was not found with criteria %v and options %v", criteria, options)
}
