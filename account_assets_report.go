package odoo

import (
	"fmt"
)

// AccountAssetsReport represents account.assets.report model.
type AccountAssetsReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountAssetsReports represents array of account.assets.report model.
type AccountAssetsReports []AccountAssetsReport

// AccountAssetsReportModel is the odoo model name.
const AccountAssetsReportModel = "account.assets.report"

// Many2One convert AccountAssetsReport to *Many2One.
func (aar *AccountAssetsReport) Many2One() *Many2One {
	return NewMany2One(aar.Id.Get(), "")
}

// CreateAccountAssetsReport creates a new account.assets.report model and returns its id.
func (c *Client) CreateAccountAssetsReport(aar *AccountAssetsReport) (int64, error) {
	ids, err := c.CreateAccountAssetsReports([]*AccountAssetsReport{aar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAssetsReport creates a new account.assets.report model and returns its id.
func (c *Client) CreateAccountAssetsReports(aars []*AccountAssetsReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range aars {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetsReportModel, vv)
}

// UpdateAccountAssetsReport updates an existing account.assets.report record.
func (c *Client) UpdateAccountAssetsReport(aar *AccountAssetsReport) error {
	return c.UpdateAccountAssetsReports([]int64{aar.Id.Get()}, aar)
}

// UpdateAccountAssetsReports updates existing account.assets.report records.
// All records (represented by ids) will be updated by aar values.
func (c *Client) UpdateAccountAssetsReports(ids []int64, aar *AccountAssetsReport) error {
	return c.Update(AccountAssetsReportModel, ids, aar)
}

// DeleteAccountAssetsReport deletes an existing account.assets.report record.
func (c *Client) DeleteAccountAssetsReport(id int64) error {
	return c.DeleteAccountAssetsReports([]int64{id})
}

// DeleteAccountAssetsReports deletes existing account.assets.report records.
func (c *Client) DeleteAccountAssetsReports(ids []int64) error {
	return c.Delete(AccountAssetsReportModel, ids)
}

// GetAccountAssetsReport gets account.assets.report existing record.
func (c *Client) GetAccountAssetsReport(id int64) (*AccountAssetsReport, error) {
	aars, err := c.GetAccountAssetsReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if aars != nil && len(*aars) > 0 {
		return &((*aars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.assets.report not found", id)
}

// GetAccountAssetsReports gets account.assets.report existing records.
func (c *Client) GetAccountAssetsReports(ids []int64) (*AccountAssetsReports, error) {
	aars := &AccountAssetsReports{}
	if err := c.Read(AccountAssetsReportModel, ids, nil, aars); err != nil {
		return nil, err
	}
	return aars, nil
}

// FindAccountAssetsReport finds account.assets.report record by querying it with criteria.
func (c *Client) FindAccountAssetsReport(criteria *Criteria) (*AccountAssetsReport, error) {
	aars := &AccountAssetsReports{}
	if err := c.SearchRead(AccountAssetsReportModel, criteria, NewOptions().Limit(1), aars); err != nil {
		return nil, err
	}
	if aars != nil && len(*aars) > 0 {
		return &((*aars)[0]), nil
	}
	return nil, fmt.Errorf("account.assets.report was not found with criteria %v", criteria)
}

// FindAccountAssetsReports finds account.assets.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetsReports(criteria *Criteria, options *Options) (*AccountAssetsReports, error) {
	aars := &AccountAssetsReports{}
	if err := c.SearchRead(AccountAssetsReportModel, criteria, options, aars); err != nil {
		return nil, err
	}
	return aars, nil
}

// FindAccountAssetsReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetsReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAssetsReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAssetsReportId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetsReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetsReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.assets.report was not found with criteria %v and options %v", criteria, options)
}
