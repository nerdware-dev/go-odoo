package odoo

import (
	"fmt"
)

// SaleSubscriptionLine represents sale.subscription.line model.
type SaleSubscriptionLine struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	AnalyticAccountId    *Many2One `xmlrpc:"analytic_account_id,omptempty"`
	CompanyId            *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	Discount             *Float    `xmlrpc:"discount,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	Name                 *String   `xmlrpc:"name,omptempty"`
	PriceSubtotal        *Float    `xmlrpc:"price_subtotal,omptempty"`
	PriceUnit            *Float    `xmlrpc:"price_unit,omptempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omptempty"`
	Quantity             *Float    `xmlrpc:"quantity,omptempty"`
	UomId                *Many2One `xmlrpc:"uom_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionLines represents array of sale.subscription.line model.
type SaleSubscriptionLines []SaleSubscriptionLine

// SaleSubscriptionLineModel is the odoo model name.
const SaleSubscriptionLineModel = "sale.subscription.line"

// Many2One convert SaleSubscriptionLine to *Many2One.
func (ssl *SaleSubscriptionLine) Many2One() *Many2One {
	return NewMany2One(ssl.Id.Get(), "")
}

// CreateSaleSubscriptionLine creates a new sale.subscription.line model and returns its id.
func (c *Client) CreateSaleSubscriptionLine(ssl *SaleSubscriptionLine) (int64, error) {
	ids, err := c.CreateSaleSubscriptionLines([]*SaleSubscriptionLine{ssl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionLine creates a new sale.subscription.line model and returns its id.
func (c *Client) CreateSaleSubscriptionLines(ssls []*SaleSubscriptionLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssls {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionLineModel, vv)
}

// UpdateSaleSubscriptionLine updates an existing sale.subscription.line record.
func (c *Client) UpdateSaleSubscriptionLine(ssl *SaleSubscriptionLine) error {
	return c.UpdateSaleSubscriptionLines([]int64{ssl.Id.Get()}, ssl)
}

// UpdateSaleSubscriptionLines updates existing sale.subscription.line records.
// All records (represented by ids) will be updated by ssl values.
func (c *Client) UpdateSaleSubscriptionLines(ids []int64, ssl *SaleSubscriptionLine) error {
	return c.Update(SaleSubscriptionLineModel, ids, ssl)
}

// DeleteSaleSubscriptionLine deletes an existing sale.subscription.line record.
func (c *Client) DeleteSaleSubscriptionLine(id int64) error {
	return c.DeleteSaleSubscriptionLines([]int64{id})
}

// DeleteSaleSubscriptionLines deletes existing sale.subscription.line records.
func (c *Client) DeleteSaleSubscriptionLines(ids []int64) error {
	return c.Delete(SaleSubscriptionLineModel, ids)
}

// GetSaleSubscriptionLine gets sale.subscription.line existing record.
func (c *Client) GetSaleSubscriptionLine(id int64) (*SaleSubscriptionLine, error) {
	ssls, err := c.GetSaleSubscriptionLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssls != nil && len(*ssls) > 0 {
		return &((*ssls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.line not found", id)
}

// GetSaleSubscriptionLines gets sale.subscription.line existing records.
func (c *Client) GetSaleSubscriptionLines(ids []int64) (*SaleSubscriptionLines, error) {
	ssls := &SaleSubscriptionLines{}
	if err := c.Read(SaleSubscriptionLineModel, ids, nil, ssls); err != nil {
		return nil, err
	}
	return ssls, nil
}

// FindSaleSubscriptionLine finds sale.subscription.line record by querying it with criteria.
func (c *Client) FindSaleSubscriptionLine(criteria *Criteria) (*SaleSubscriptionLine, error) {
	ssls := &SaleSubscriptionLines{}
	if err := c.SearchRead(SaleSubscriptionLineModel, criteria, NewOptions().Limit(1), ssls); err != nil {
		return nil, err
	}
	if ssls != nil && len(*ssls) > 0 {
		return &((*ssls)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.line was not found with criteria %v", criteria)
}

// FindSaleSubscriptionLines finds sale.subscription.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionLines(criteria *Criteria, options *Options) (*SaleSubscriptionLines, error) {
	ssls := &SaleSubscriptionLines{}
	if err := c.SearchRead(SaleSubscriptionLineModel, criteria, options, ssls); err != nil {
		return nil, err
	}
	return ssls, nil
}

// FindSaleSubscriptionLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionLineId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.line was not found with criteria %v and options %v", criteria, options)
}
