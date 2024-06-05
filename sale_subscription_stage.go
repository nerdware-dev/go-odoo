package odoo

import (
	"fmt"
)

// SaleSubscriptionStage represents sale.subscription.stage model.
type SaleSubscriptionStage struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	Description      *String   `xmlrpc:"description,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Fold             *Bool     `xmlrpc:"fold,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	InProgress       *Bool     `xmlrpc:"in_progress,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	RatingTemplateId *Many2One `xmlrpc:"rating_template_id,omptempty"`
	Sequence         *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionStages represents array of sale.subscription.stage model.
type SaleSubscriptionStages []SaleSubscriptionStage

// SaleSubscriptionStageModel is the odoo model name.
const SaleSubscriptionStageModel = "sale.subscription.stage"

// Many2One convert SaleSubscriptionStage to *Many2One.
func (sss *SaleSubscriptionStage) Many2One() *Many2One {
	return NewMany2One(sss.Id.Get(), "")
}

// CreateSaleSubscriptionStage creates a new sale.subscription.stage model and returns its id.
func (c *Client) CreateSaleSubscriptionStage(sss *SaleSubscriptionStage) (int64, error) {
	ids, err := c.CreateSaleSubscriptionStages([]*SaleSubscriptionStage{sss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionStage creates a new sale.subscription.stage model and returns its id.
func (c *Client) CreateSaleSubscriptionStages(ssss []*SaleSubscriptionStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssss {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionStageModel, vv)
}

// UpdateSaleSubscriptionStage updates an existing sale.subscription.stage record.
func (c *Client) UpdateSaleSubscriptionStage(sss *SaleSubscriptionStage) error {
	return c.UpdateSaleSubscriptionStages([]int64{sss.Id.Get()}, sss)
}

// UpdateSaleSubscriptionStages updates existing sale.subscription.stage records.
// All records (represented by ids) will be updated by sss values.
func (c *Client) UpdateSaleSubscriptionStages(ids []int64, sss *SaleSubscriptionStage) error {
	return c.Update(SaleSubscriptionStageModel, ids, sss)
}

// DeleteSaleSubscriptionStage deletes an existing sale.subscription.stage record.
func (c *Client) DeleteSaleSubscriptionStage(id int64) error {
	return c.DeleteSaleSubscriptionStages([]int64{id})
}

// DeleteSaleSubscriptionStages deletes existing sale.subscription.stage records.
func (c *Client) DeleteSaleSubscriptionStages(ids []int64) error {
	return c.Delete(SaleSubscriptionStageModel, ids)
}

// GetSaleSubscriptionStage gets sale.subscription.stage existing record.
func (c *Client) GetSaleSubscriptionStage(id int64) (*SaleSubscriptionStage, error) {
	ssss, err := c.GetSaleSubscriptionStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssss != nil && len(*ssss) > 0 {
		return &((*ssss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.stage not found", id)
}

// GetSaleSubscriptionStages gets sale.subscription.stage existing records.
func (c *Client) GetSaleSubscriptionStages(ids []int64) (*SaleSubscriptionStages, error) {
	ssss := &SaleSubscriptionStages{}
	if err := c.Read(SaleSubscriptionStageModel, ids, nil, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionStage finds sale.subscription.stage record by querying it with criteria.
func (c *Client) FindSaleSubscriptionStage(criteria *Criteria) (*SaleSubscriptionStage, error) {
	ssss := &SaleSubscriptionStages{}
	if err := c.SearchRead(SaleSubscriptionStageModel, criteria, NewOptions().Limit(1), ssss); err != nil {
		return nil, err
	}
	if ssss != nil && len(*ssss) > 0 {
		return &((*ssss)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.stage was not found with criteria %v", criteria)
}

// FindSaleSubscriptionStages finds sale.subscription.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionStages(criteria *Criteria, options *Options) (*SaleSubscriptionStages, error) {
	ssss := &SaleSubscriptionStages{}
	if err := c.SearchRead(SaleSubscriptionStageModel, criteria, options, ssss); err != nil {
		return nil, err
	}
	return ssss, nil
}

// FindSaleSubscriptionStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionStageId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.stage was not found with criteria %v and options %v", criteria, options)
}
