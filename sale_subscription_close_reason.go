package odoo

import (
	"fmt"
)

// SaleSubscriptionCloseReason represents sale.subscription.close.reason model.
type SaleSubscriptionCloseReason struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionCloseReasons represents array of sale.subscription.close.reason model.
type SaleSubscriptionCloseReasons []SaleSubscriptionCloseReason

// SaleSubscriptionCloseReasonModel is the odoo model name.
const SaleSubscriptionCloseReasonModel = "sale.subscription.close.reason"

// Many2One convert SaleSubscriptionCloseReason to *Many2One.
func (sscr *SaleSubscriptionCloseReason) Many2One() *Many2One {
	return NewMany2One(sscr.Id.Get(), "")
}

// CreateSaleSubscriptionCloseReason creates a new sale.subscription.close.reason model and returns its id.
func (c *Client) CreateSaleSubscriptionCloseReason(sscr *SaleSubscriptionCloseReason) (int64, error) {
	ids, err := c.CreateSaleSubscriptionCloseReasons([]*SaleSubscriptionCloseReason{sscr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionCloseReason creates a new sale.subscription.close.reason model and returns its id.
func (c *Client) CreateSaleSubscriptionCloseReasons(sscrs []*SaleSubscriptionCloseReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range sscrs {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionCloseReasonModel, vv)
}

// UpdateSaleSubscriptionCloseReason updates an existing sale.subscription.close.reason record.
func (c *Client) UpdateSaleSubscriptionCloseReason(sscr *SaleSubscriptionCloseReason) error {
	return c.UpdateSaleSubscriptionCloseReasons([]int64{sscr.Id.Get()}, sscr)
}

// UpdateSaleSubscriptionCloseReasons updates existing sale.subscription.close.reason records.
// All records (represented by ids) will be updated by sscr values.
func (c *Client) UpdateSaleSubscriptionCloseReasons(ids []int64, sscr *SaleSubscriptionCloseReason) error {
	return c.Update(SaleSubscriptionCloseReasonModel, ids, sscr)
}

// DeleteSaleSubscriptionCloseReason deletes an existing sale.subscription.close.reason record.
func (c *Client) DeleteSaleSubscriptionCloseReason(id int64) error {
	return c.DeleteSaleSubscriptionCloseReasons([]int64{id})
}

// DeleteSaleSubscriptionCloseReasons deletes existing sale.subscription.close.reason records.
func (c *Client) DeleteSaleSubscriptionCloseReasons(ids []int64) error {
	return c.Delete(SaleSubscriptionCloseReasonModel, ids)
}

// GetSaleSubscriptionCloseReason gets sale.subscription.close.reason existing record.
func (c *Client) GetSaleSubscriptionCloseReason(id int64) (*SaleSubscriptionCloseReason, error) {
	sscrs, err := c.GetSaleSubscriptionCloseReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	if sscrs != nil && len(*sscrs) > 0 {
		return &((*sscrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.close.reason not found", id)
}

// GetSaleSubscriptionCloseReasons gets sale.subscription.close.reason existing records.
func (c *Client) GetSaleSubscriptionCloseReasons(ids []int64) (*SaleSubscriptionCloseReasons, error) {
	sscrs := &SaleSubscriptionCloseReasons{}
	if err := c.Read(SaleSubscriptionCloseReasonModel, ids, nil, sscrs); err != nil {
		return nil, err
	}
	return sscrs, nil
}

// FindSaleSubscriptionCloseReason finds sale.subscription.close.reason record by querying it with criteria.
func (c *Client) FindSaleSubscriptionCloseReason(criteria *Criteria) (*SaleSubscriptionCloseReason, error) {
	sscrs := &SaleSubscriptionCloseReasons{}
	if err := c.SearchRead(SaleSubscriptionCloseReasonModel, criteria, NewOptions().Limit(1), sscrs); err != nil {
		return nil, err
	}
	if sscrs != nil && len(*sscrs) > 0 {
		return &((*sscrs)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.close.reason was not found with criteria %v", criteria)
}

// FindSaleSubscriptionCloseReasons finds sale.subscription.close.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionCloseReasons(criteria *Criteria, options *Options) (*SaleSubscriptionCloseReasons, error) {
	sscrs := &SaleSubscriptionCloseReasons{}
	if err := c.SearchRead(SaleSubscriptionCloseReasonModel, criteria, options, sscrs); err != nil {
		return nil, err
	}
	return sscrs, nil
}

// FindSaleSubscriptionCloseReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionCloseReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionCloseReasonModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionCloseReasonId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionCloseReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionCloseReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.close.reason was not found with criteria %v and options %v", criteria, options)
}
