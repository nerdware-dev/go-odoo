package odoo

import (
	"fmt"
)

// SaleSubscriptionWizard represents sale.subscription.wizard model.
type SaleSubscriptionWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DateFrom       *Time     `xmlrpc:"date_from,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	OptionLines    *Relation `xmlrpc:"option_lines,omptempty"`
	SubscriptionId *Many2One `xmlrpc:"subscription_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionWizards represents array of sale.subscription.wizard model.
type SaleSubscriptionWizards []SaleSubscriptionWizard

// SaleSubscriptionWizardModel is the odoo model name.
const SaleSubscriptionWizardModel = "sale.subscription.wizard"

// Many2One convert SaleSubscriptionWizard to *Many2One.
func (ssw *SaleSubscriptionWizard) Many2One() *Many2One {
	return NewMany2One(ssw.Id.Get(), "")
}

// CreateSaleSubscriptionWizard creates a new sale.subscription.wizard model and returns its id.
func (c *Client) CreateSaleSubscriptionWizard(ssw *SaleSubscriptionWizard) (int64, error) {
	ids, err := c.CreateSaleSubscriptionWizards([]*SaleSubscriptionWizard{ssw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionWizard creates a new sale.subscription.wizard model and returns its id.
func (c *Client) CreateSaleSubscriptionWizards(ssws []*SaleSubscriptionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssws {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionWizardModel, vv)
}

// UpdateSaleSubscriptionWizard updates an existing sale.subscription.wizard record.
func (c *Client) UpdateSaleSubscriptionWizard(ssw *SaleSubscriptionWizard) error {
	return c.UpdateSaleSubscriptionWizards([]int64{ssw.Id.Get()}, ssw)
}

// UpdateSaleSubscriptionWizards updates existing sale.subscription.wizard records.
// All records (represented by ids) will be updated by ssw values.
func (c *Client) UpdateSaleSubscriptionWizards(ids []int64, ssw *SaleSubscriptionWizard) error {
	return c.Update(SaleSubscriptionWizardModel, ids, ssw)
}

// DeleteSaleSubscriptionWizard deletes an existing sale.subscription.wizard record.
func (c *Client) DeleteSaleSubscriptionWizard(id int64) error {
	return c.DeleteSaleSubscriptionWizards([]int64{id})
}

// DeleteSaleSubscriptionWizards deletes existing sale.subscription.wizard records.
func (c *Client) DeleteSaleSubscriptionWizards(ids []int64) error {
	return c.Delete(SaleSubscriptionWizardModel, ids)
}

// GetSaleSubscriptionWizard gets sale.subscription.wizard existing record.
func (c *Client) GetSaleSubscriptionWizard(id int64) (*SaleSubscriptionWizard, error) {
	ssws, err := c.GetSaleSubscriptionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssws != nil && len(*ssws) > 0 {
		return &((*ssws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.wizard not found", id)
}

// GetSaleSubscriptionWizards gets sale.subscription.wizard existing records.
func (c *Client) GetSaleSubscriptionWizards(ids []int64) (*SaleSubscriptionWizards, error) {
	ssws := &SaleSubscriptionWizards{}
	if err := c.Read(SaleSubscriptionWizardModel, ids, nil, ssws); err != nil {
		return nil, err
	}
	return ssws, nil
}

// FindSaleSubscriptionWizard finds sale.subscription.wizard record by querying it with criteria.
func (c *Client) FindSaleSubscriptionWizard(criteria *Criteria) (*SaleSubscriptionWizard, error) {
	ssws := &SaleSubscriptionWizards{}
	if err := c.SearchRead(SaleSubscriptionWizardModel, criteria, NewOptions().Limit(1), ssws); err != nil {
		return nil, err
	}
	if ssws != nil && len(*ssws) > 0 {
		return &((*ssws)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.wizard was not found with criteria %v", criteria)
}

// FindSaleSubscriptionWizards finds sale.subscription.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionWizards(criteria *Criteria, options *Options) (*SaleSubscriptionWizards, error) {
	ssws := &SaleSubscriptionWizards{}
	if err := c.SearchRead(SaleSubscriptionWizardModel, criteria, options, ssws); err != nil {
		return nil, err
	}
	return ssws, nil
}

// FindSaleSubscriptionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionWizardId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.wizard was not found with criteria %v and options %v", criteria, options)
}
