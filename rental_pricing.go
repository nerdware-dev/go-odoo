package odoo

import (
	"fmt"
)

// RentalPricing represents rental.pricing model.
type RentalPricing struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Duration          *Int       `xmlrpc:"duration,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	Price             *Float     `xmlrpc:"price,omptempty"`
	PricelistId       *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProductTemplateId *Many2One  `xmlrpc:"product_template_id,omptempty"`
	ProductVariantIds *Relation  `xmlrpc:"product_variant_ids,omptempty"`
	Unit              *Selection `xmlrpc:"unit,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// RentalPricings represents array of rental.pricing model.
type RentalPricings []RentalPricing

// RentalPricingModel is the odoo model name.
const RentalPricingModel = "rental.pricing"

// Many2One convert RentalPricing to *Many2One.
func (rp *RentalPricing) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateRentalPricing creates a new rental.pricing model and returns its id.
func (c *Client) CreateRentalPricing(rp *RentalPricing) (int64, error) {
	ids, err := c.CreateRentalPricings([]*RentalPricing{rp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRentalPricing creates a new rental.pricing model and returns its id.
func (c *Client) CreateRentalPricings(rps []*RentalPricing) ([]int64, error) {
	var vv []interface{}
	for _, v := range rps {
		vv = append(vv, v)
	}
	return c.Create(RentalPricingModel, vv)
}

// UpdateRentalPricing updates an existing rental.pricing record.
func (c *Client) UpdateRentalPricing(rp *RentalPricing) error {
	return c.UpdateRentalPricings([]int64{rp.Id.Get()}, rp)
}

// UpdateRentalPricings updates existing rental.pricing records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateRentalPricings(ids []int64, rp *RentalPricing) error {
	return c.Update(RentalPricingModel, ids, rp)
}

// DeleteRentalPricing deletes an existing rental.pricing record.
func (c *Client) DeleteRentalPricing(id int64) error {
	return c.DeleteRentalPricings([]int64{id})
}

// DeleteRentalPricings deletes existing rental.pricing records.
func (c *Client) DeleteRentalPricings(ids []int64) error {
	return c.Delete(RentalPricingModel, ids)
}

// GetRentalPricing gets rental.pricing existing record.
func (c *Client) GetRentalPricing(id int64) (*RentalPricing, error) {
	rps, err := c.GetRentalPricings([]int64{id})
	if err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of rental.pricing not found", id)
}

// GetRentalPricings gets rental.pricing existing records.
func (c *Client) GetRentalPricings(ids []int64) (*RentalPricings, error) {
	rps := &RentalPricings{}
	if err := c.Read(RentalPricingModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindRentalPricing finds rental.pricing record by querying it with criteria.
func (c *Client) FindRentalPricing(criteria *Criteria) (*RentalPricing, error) {
	rps := &RentalPricings{}
	if err := c.SearchRead(RentalPricingModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("rental.pricing was not found with criteria %v", criteria)
}

// FindRentalPricings finds rental.pricing records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalPricings(criteria *Criteria, options *Options) (*RentalPricings, error) {
	rps := &RentalPricings{}
	if err := c.SearchRead(RentalPricingModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindRentalPricingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRentalPricingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RentalPricingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRentalPricingId finds record id by querying it with criteria.
func (c *Client) FindRentalPricingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RentalPricingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("rental.pricing was not found with criteria %v and options %v", criteria, options)
}
