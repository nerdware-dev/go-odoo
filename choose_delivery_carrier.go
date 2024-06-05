package odoo

import (
	"fmt"
)

// ChooseDeliveryCarrier represents choose.delivery.carrier model.
type ChooseDeliveryCarrier struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	AvailableCarrierIds *Relation  `xmlrpc:"available_carrier_ids,omptempty"`
	CarrierId           *Many2One  `xmlrpc:"carrier_id,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omptempty"`
	DeliveryMessage     *String    `xmlrpc:"delivery_message,omptempty"`
	DeliveryPrice       *Float     `xmlrpc:"delivery_price,omptempty"`
	DeliveryType        *Selection `xmlrpc:"delivery_type,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	DisplayPrice        *Float     `xmlrpc:"display_price,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InvoicingMessage    *String    `xmlrpc:"invoicing_message,omptempty"`
	OrderId             *Many2One  `xmlrpc:"order_id,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ChooseDeliveryCarriers represents array of choose.delivery.carrier model.
type ChooseDeliveryCarriers []ChooseDeliveryCarrier

// ChooseDeliveryCarrierModel is the odoo model name.
const ChooseDeliveryCarrierModel = "choose.delivery.carrier"

// Many2One convert ChooseDeliveryCarrier to *Many2One.
func (cdc *ChooseDeliveryCarrier) Many2One() *Many2One {
	return NewMany2One(cdc.Id.Get(), "")
}

// CreateChooseDeliveryCarrier creates a new choose.delivery.carrier model and returns its id.
func (c *Client) CreateChooseDeliveryCarrier(cdc *ChooseDeliveryCarrier) (int64, error) {
	ids, err := c.CreateChooseDeliveryCarriers([]*ChooseDeliveryCarrier{cdc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChooseDeliveryCarrier creates a new choose.delivery.carrier model and returns its id.
func (c *Client) CreateChooseDeliveryCarriers(cdcs []*ChooseDeliveryCarrier) ([]int64, error) {
	var vv []interface{}
	for _, v := range cdcs {
		vv = append(vv, v)
	}
	return c.Create(ChooseDeliveryCarrierModel, vv)
}

// UpdateChooseDeliveryCarrier updates an existing choose.delivery.carrier record.
func (c *Client) UpdateChooseDeliveryCarrier(cdc *ChooseDeliveryCarrier) error {
	return c.UpdateChooseDeliveryCarriers([]int64{cdc.Id.Get()}, cdc)
}

// UpdateChooseDeliveryCarriers updates existing choose.delivery.carrier records.
// All records (represented by ids) will be updated by cdc values.
func (c *Client) UpdateChooseDeliveryCarriers(ids []int64, cdc *ChooseDeliveryCarrier) error {
	return c.Update(ChooseDeliveryCarrierModel, ids, cdc)
}

// DeleteChooseDeliveryCarrier deletes an existing choose.delivery.carrier record.
func (c *Client) DeleteChooseDeliveryCarrier(id int64) error {
	return c.DeleteChooseDeliveryCarriers([]int64{id})
}

// DeleteChooseDeliveryCarriers deletes existing choose.delivery.carrier records.
func (c *Client) DeleteChooseDeliveryCarriers(ids []int64) error {
	return c.Delete(ChooseDeliveryCarrierModel, ids)
}

// GetChooseDeliveryCarrier gets choose.delivery.carrier existing record.
func (c *Client) GetChooseDeliveryCarrier(id int64) (*ChooseDeliveryCarrier, error) {
	cdcs, err := c.GetChooseDeliveryCarriers([]int64{id})
	if err != nil {
		return nil, err
	}
	if cdcs != nil && len(*cdcs) > 0 {
		return &((*cdcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of choose.delivery.carrier not found", id)
}

// GetChooseDeliveryCarriers gets choose.delivery.carrier existing records.
func (c *Client) GetChooseDeliveryCarriers(ids []int64) (*ChooseDeliveryCarriers, error) {
	cdcs := &ChooseDeliveryCarriers{}
	if err := c.Read(ChooseDeliveryCarrierModel, ids, nil, cdcs); err != nil {
		return nil, err
	}
	return cdcs, nil
}

// FindChooseDeliveryCarrier finds choose.delivery.carrier record by querying it with criteria.
func (c *Client) FindChooseDeliveryCarrier(criteria *Criteria) (*ChooseDeliveryCarrier, error) {
	cdcs := &ChooseDeliveryCarriers{}
	if err := c.SearchRead(ChooseDeliveryCarrierModel, criteria, NewOptions().Limit(1), cdcs); err != nil {
		return nil, err
	}
	if cdcs != nil && len(*cdcs) > 0 {
		return &((*cdcs)[0]), nil
	}
	return nil, fmt.Errorf("choose.delivery.carrier was not found with criteria %v", criteria)
}

// FindChooseDeliveryCarriers finds choose.delivery.carrier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChooseDeliveryCarriers(criteria *Criteria, options *Options) (*ChooseDeliveryCarriers, error) {
	cdcs := &ChooseDeliveryCarriers{}
	if err := c.SearchRead(ChooseDeliveryCarrierModel, criteria, options, cdcs); err != nil {
		return nil, err
	}
	return cdcs, nil
}

// FindChooseDeliveryCarrierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChooseDeliveryCarrierIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ChooseDeliveryCarrierModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindChooseDeliveryCarrierId finds record id by querying it with criteria.
func (c *Client) FindChooseDeliveryCarrierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChooseDeliveryCarrierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("choose.delivery.carrier was not found with criteria %v and options %v", criteria, options)
}
