package odoo

import (
	"fmt"
)

// SaleSubscriptionSnapshot represents sale.subscription.snapshot model.
type SaleSubscriptionSnapshot struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	Date             *Time     `xmlrpc:"date,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	RecurringMonthly *Float    `xmlrpc:"recurring_monthly,omptempty"`
	SubscriptionId   *Many2One `xmlrpc:"subscription_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionSnapshots represents array of sale.subscription.snapshot model.
type SaleSubscriptionSnapshots []SaleSubscriptionSnapshot

// SaleSubscriptionSnapshotModel is the odoo model name.
const SaleSubscriptionSnapshotModel = "sale.subscription.snapshot"

// Many2One convert SaleSubscriptionSnapshot to *Many2One.
func (sss *SaleSubscriptionSnapshot) Many2One() *Many2One {
	return NewMany2One(sss.Id.Get(), "")
}

// CreateSaleSubscriptionSnapshot creates a new sale.subscription.snapshot model and returns its id.
func (c *Client) CreateSaleSubscriptionSnapshot(sss *SaleSubscriptionSnapshot) (int64, error) {
	ids, err := c.CreateSaleSubscriptionSnapshots([]*SaleSubscriptionSnapshot{sss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionSnapshot creates a new sale.subscription.snapshot model and returns its id.
func (c *Client) CreateSaleSubscriptionSnapshots(ssss []*SaleSubscriptionSnapshot) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssss {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionSnapshotModel, vv)
}

// UpdateSaleSubscriptionSnapshot updates an existing sale.subscription.snapshot record.
func (c *Client) UpdateSaleSubscriptionSnapshot(sss *SaleSubscriptionSnapshot) error {
	return c.UpdateSaleSubscriptionSnapshots([]int64{sss.Id.Get()}, sss)
}

// UpdateSaleSubscriptionSnapshots updates existing sale.subscription.snapshot records.
// All records (represented by ids) will be updated by sss values.
func (c *Client) UpdateSaleSubscriptionSnapshots(ids []int64, sss *SaleSubscriptionSnapshot) error {
	return c.Update(SaleSubscriptionSnapshotModel, ids, sss)
}

// DeleteSaleSubscriptionSnapshot deletes an existing sale.subscription.snapshot record.
func (c *Client) DeleteSaleSubscriptionSnapshot(id int64) error {
	return c.DeleteSaleSubscriptionSnapshots([]int64{id})
}

// DeleteSaleSubscriptionSnapshots deletes existing sale.subscription.snapshot records.
func (c *Client) DeleteSaleSubscriptionSnapshots(ids []int64) error {
	return c.Delete(SaleSubscriptionSnapshotModel, ids)
}

// GetSaleSubscriptionSnapshot gets sale.subscription.snapshot existing record.
func (c *Client) GetSaleSubscriptionSnapshot(id int64) (*SaleSubscriptionSnapshot, error) {
	ssss, err := c.GetSaleSubscriptionSnapshots([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssss != nil && len(*ssss) > 0 {
		return &((*ssss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.snapshot not found", id)
}

// GetSaleSubscriptionSnapshots gets sale.subscription.snapshot existing records.
func (c *Client) GetSaleSubscriptionSnapshots(ids []int64) (*SaleSubscriptionSnapshots, error) {
	ssss := &SaleSubscriptionSnapshots{}
	if err := c.Read(SaleSubscriptionSnapshotModel, ids, nil, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionSnapshot finds sale.subscription.snapshot record by querying it with criteria.
func (c *Client) FindSaleSubscriptionSnapshot(criteria *Criteria) (*SaleSubscriptionSnapshot, error) {
	ssss := &SaleSubscriptionSnapshots{}
	if err := c.SearchRead(SaleSubscriptionSnapshotModel, criteria, NewOptions().Limit(1), ssss); err != nil {
		return nil, err
	}
	if ssss != nil && len(*ssss) > 0 {
		return &((*ssss)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.snapshot was not found with criteria %v", criteria)
}

// FindSaleSubscriptionSnapshots finds sale.subscription.snapshot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionSnapshots(criteria *Criteria, options *Options) (*SaleSubscriptionSnapshots, error) {
	ssss := &SaleSubscriptionSnapshots{}
	if err := c.SearchRead(SaleSubscriptionSnapshotModel, criteria, options, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionSnapshotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionSnapshotIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionSnapshotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionSnapshotId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionSnapshotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionSnapshotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.snapshot was not found with criteria %v and options %v", criteria, options)
}
