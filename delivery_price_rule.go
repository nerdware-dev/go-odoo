package odoo

import (
	"fmt"
)

// DeliveryPriceRule represents delivery.price.rule model.
type DeliveryPriceRule struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CarrierId      *Many2One  `xmlrpc:"carrier_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	ListBasePrice  *Float     `xmlrpc:"list_base_price,omptempty"`
	ListPrice      *Float     `xmlrpc:"list_price,omptempty"`
	MaxValue       *Float     `xmlrpc:"max_value,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	Operator       *Selection `xmlrpc:"operator,omptempty"`
	Sequence       *Int       `xmlrpc:"sequence,omptempty"`
	Variable       *Selection `xmlrpc:"variable,omptempty"`
	VariableFactor *Selection `xmlrpc:"variable_factor,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DeliveryPriceRules represents array of delivery.price.rule model.
type DeliveryPriceRules []DeliveryPriceRule

// DeliveryPriceRuleModel is the odoo model name.
const DeliveryPriceRuleModel = "delivery.price.rule"

// Many2One convert DeliveryPriceRule to *Many2One.
func (dpr *DeliveryPriceRule) Many2One() *Many2One {
	return NewMany2One(dpr.Id.Get(), "")
}

// CreateDeliveryPriceRule creates a new delivery.price.rule model and returns its id.
func (c *Client) CreateDeliveryPriceRule(dpr *DeliveryPriceRule) (int64, error) {
	ids, err := c.CreateDeliveryPriceRules([]*DeliveryPriceRule{dpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDeliveryPriceRule creates a new delivery.price.rule model and returns its id.
func (c *Client) CreateDeliveryPriceRules(dprs []*DeliveryPriceRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range dprs {
		vv = append(vv, v)
	}
	return c.Create(DeliveryPriceRuleModel, vv)
}

// UpdateDeliveryPriceRule updates an existing delivery.price.rule record.
func (c *Client) UpdateDeliveryPriceRule(dpr *DeliveryPriceRule) error {
	return c.UpdateDeliveryPriceRules([]int64{dpr.Id.Get()}, dpr)
}

// UpdateDeliveryPriceRules updates existing delivery.price.rule records.
// All records (represented by ids) will be updated by dpr values.
func (c *Client) UpdateDeliveryPriceRules(ids []int64, dpr *DeliveryPriceRule) error {
	return c.Update(DeliveryPriceRuleModel, ids, dpr)
}

// DeleteDeliveryPriceRule deletes an existing delivery.price.rule record.
func (c *Client) DeleteDeliveryPriceRule(id int64) error {
	return c.DeleteDeliveryPriceRules([]int64{id})
}

// DeleteDeliveryPriceRules deletes existing delivery.price.rule records.
func (c *Client) DeleteDeliveryPriceRules(ids []int64) error {
	return c.Delete(DeliveryPriceRuleModel, ids)
}

// GetDeliveryPriceRule gets delivery.price.rule existing record.
func (c *Client) GetDeliveryPriceRule(id int64) (*DeliveryPriceRule, error) {
	dprs, err := c.GetDeliveryPriceRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if dprs != nil && len(*dprs) > 0 {
		return &((*dprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of delivery.price.rule not found", id)
}

// GetDeliveryPriceRules gets delivery.price.rule existing records.
func (c *Client) GetDeliveryPriceRules(ids []int64) (*DeliveryPriceRules, error) {
	dprs := &DeliveryPriceRules{}
	if err := c.Read(DeliveryPriceRuleModel, ids, nil, dprs); err != nil {
		return nil, err
	}
	return dprs, nil
}

// FindDeliveryPriceRule finds delivery.price.rule record by querying it with criteria.
func (c *Client) FindDeliveryPriceRule(criteria *Criteria) (*DeliveryPriceRule, error) {
	dprs := &DeliveryPriceRules{}
	if err := c.SearchRead(DeliveryPriceRuleModel, criteria, NewOptions().Limit(1), dprs); err != nil {
		return nil, err
	}
	if dprs != nil && len(*dprs) > 0 {
		return &((*dprs)[0]), nil
	}
	return nil, fmt.Errorf("delivery.price.rule was not found with criteria %v", criteria)
}

// FindDeliveryPriceRules finds delivery.price.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryPriceRules(criteria *Criteria, options *Options) (*DeliveryPriceRules, error) {
	dprs := &DeliveryPriceRules{}
	if err := c.SearchRead(DeliveryPriceRuleModel, criteria, options, dprs); err != nil {
		return nil, err
	}
	return dprs, nil
}

// FindDeliveryPriceRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryPriceRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DeliveryPriceRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDeliveryPriceRuleId finds record id by querying it with criteria.
func (c *Client) FindDeliveryPriceRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DeliveryPriceRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("delivery.price.rule was not found with criteria %v and options %v", criteria, options)
}
