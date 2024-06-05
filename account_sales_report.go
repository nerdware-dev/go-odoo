package odoo

import (
	"fmt"
)

// AccountSalesReport represents account.sales.report model.
type AccountSalesReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountSalesReports represents array of account.sales.report model.
type AccountSalesReports []AccountSalesReport

// AccountSalesReportModel is the odoo model name.
const AccountSalesReportModel = "account.sales.report"

// Many2One convert AccountSalesReport to *Many2One.
func (asr *AccountSalesReport) Many2One() *Many2One {
	return NewMany2One(asr.Id.Get(), "")
}

// CreateAccountSalesReport creates a new account.sales.report model and returns its id.
func (c *Client) CreateAccountSalesReport(asr *AccountSalesReport) (int64, error) {
	ids, err := c.CreateAccountSalesReports([]*AccountSalesReport{asr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountSalesReport creates a new account.sales.report model and returns its id.
func (c *Client) CreateAccountSalesReports(asrs []*AccountSalesReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range asrs {
		vv = append(vv, v)
	}
	return c.Create(AccountSalesReportModel, vv)
}

// UpdateAccountSalesReport updates an existing account.sales.report record.
func (c *Client) UpdateAccountSalesReport(asr *AccountSalesReport) error {
	return c.UpdateAccountSalesReports([]int64{asr.Id.Get()}, asr)
}

// UpdateAccountSalesReports updates existing account.sales.report records.
// All records (represented by ids) will be updated by asr values.
func (c *Client) UpdateAccountSalesReports(ids []int64, asr *AccountSalesReport) error {
	return c.Update(AccountSalesReportModel, ids, asr)
}

// DeleteAccountSalesReport deletes an existing account.sales.report record.
func (c *Client) DeleteAccountSalesReport(id int64) error {
	return c.DeleteAccountSalesReports([]int64{id})
}

// DeleteAccountSalesReports deletes existing account.sales.report records.
func (c *Client) DeleteAccountSalesReports(ids []int64) error {
	return c.Delete(AccountSalesReportModel, ids)
}

// GetAccountSalesReport gets account.sales.report existing record.
func (c *Client) GetAccountSalesReport(id int64) (*AccountSalesReport, error) {
	asrs, err := c.GetAccountSalesReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if asrs != nil && len(*asrs) > 0 {
		return &((*asrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.sales.report not found", id)
}

// GetAccountSalesReports gets account.sales.report existing records.
func (c *Client) GetAccountSalesReports(ids []int64) (*AccountSalesReports, error) {
	asrs := &AccountSalesReports{}
	if err := c.Read(AccountSalesReportModel, ids, nil, asrs); err != nil {
		return nil, err
	}
	return asrs, nil
}

// FindAccountSalesReport finds account.sales.report record by querying it with criteria.
func (c *Client) FindAccountSalesReport(criteria *Criteria) (*AccountSalesReport, error) {
	asrs := &AccountSalesReports{}
	if err := c.SearchRead(AccountSalesReportModel, criteria, NewOptions().Limit(1), asrs); err != nil {
		return nil, err
	}
	if asrs != nil && len(*asrs) > 0 {
		return &((*asrs)[0]), nil
	}
	return nil, fmt.Errorf("account.sales.report was not found with criteria %v", criteria)
}

// FindAccountSalesReports finds account.sales.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSalesReports(criteria *Criteria, options *Options) (*AccountSalesReports, error) {
	asrs := &AccountSalesReports{}
	if err := c.SearchRead(AccountSalesReportModel, criteria, options, asrs); err != nil {
		return nil, err
	}
	return asrs, nil
}

// FindAccountSalesReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSalesReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountSalesReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountSalesReportId finds record id by querying it with criteria.
func (c *Client) FindAccountSalesReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSalesReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.sales.report was not found with criteria %v and options %v", criteria, options)
}
