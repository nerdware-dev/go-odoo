package odoo

import (
	"fmt"
)

// DeliveryCarrier represents delivery.carrier model.
type DeliveryCarrier struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	Amount                   *Float     `xmlrpc:"amount,omptempty"`
	CanGenerateReturn        *Bool      `xmlrpc:"can_generate_return,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryIds               *Relation  `xmlrpc:"country_ids,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DebugLogging             *Bool      `xmlrpc:"debug_logging,omptempty"`
	DeliveryType             *Selection `xmlrpc:"delivery_type,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FixedPrice               *Float     `xmlrpc:"fixed_price,omptempty"`
	FreeOver                 *Bool      `xmlrpc:"free_over,omptempty"`
	GetReturnLabelFromPortal *Bool      `xmlrpc:"get_return_label_from_portal,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IntegrationLevel         *Selection `xmlrpc:"integration_level,omptempty"`
	InvoicePolicy            *Selection `xmlrpc:"invoice_policy,omptempty"`
	Margin                   *Float     `xmlrpc:"margin,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	PriceRuleIds             *Relation  `xmlrpc:"price_rule_ids,omptempty"`
	ProdEnvironment          *Bool      `xmlrpc:"prod_environment,omptempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omptempty"`
	ReturnLabelOnDelivery    *Bool      `xmlrpc:"return_label_on_delivery,omptempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omptempty"`
	StateIds                 *Relation  `xmlrpc:"state_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
	ZipFrom                  *String    `xmlrpc:"zip_from,omptempty"`
	ZipTo                    *String    `xmlrpc:"zip_to,omptempty"`
}

// DeliveryCarriers represents array of delivery.carrier model.
type DeliveryCarriers []DeliveryCarrier

// DeliveryCarrierModel is the odoo model name.
const DeliveryCarrierModel = "delivery.carrier"

// Many2One convert DeliveryCarrier to *Many2One.
func (dc *DeliveryCarrier) Many2One() *Many2One {
	return NewMany2One(dc.Id.Get(), "")
}

// CreateDeliveryCarrier creates a new delivery.carrier model and returns its id.
func (c *Client) CreateDeliveryCarrier(dc *DeliveryCarrier) (int64, error) {
	ids, err := c.CreateDeliveryCarriers([]*DeliveryCarrier{dc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDeliveryCarrier creates a new delivery.carrier model and returns its id.
func (c *Client) CreateDeliveryCarriers(dcs []*DeliveryCarrier) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcs {
		vv = append(vv, v)
	}
	return c.Create(DeliveryCarrierModel, vv)
}

// UpdateDeliveryCarrier updates an existing delivery.carrier record.
func (c *Client) UpdateDeliveryCarrier(dc *DeliveryCarrier) error {
	return c.UpdateDeliveryCarriers([]int64{dc.Id.Get()}, dc)
}

// UpdateDeliveryCarriers updates existing delivery.carrier records.
// All records (represented by ids) will be updated by dc values.
func (c *Client) UpdateDeliveryCarriers(ids []int64, dc *DeliveryCarrier) error {
	return c.Update(DeliveryCarrierModel, ids, dc)
}

// DeleteDeliveryCarrier deletes an existing delivery.carrier record.
func (c *Client) DeleteDeliveryCarrier(id int64) error {
	return c.DeleteDeliveryCarriers([]int64{id})
}

// DeleteDeliveryCarriers deletes existing delivery.carrier records.
func (c *Client) DeleteDeliveryCarriers(ids []int64) error {
	return c.Delete(DeliveryCarrierModel, ids)
}

// GetDeliveryCarrier gets delivery.carrier existing record.
func (c *Client) GetDeliveryCarrier(id int64) (*DeliveryCarrier, error) {
	dcs, err := c.GetDeliveryCarriers([]int64{id})
	if err != nil {
		return nil, err
	}
	if dcs != nil && len(*dcs) > 0 {
		return &((*dcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of delivery.carrier not found", id)
}

// GetDeliveryCarriers gets delivery.carrier existing records.
func (c *Client) GetDeliveryCarriers(ids []int64) (*DeliveryCarriers, error) {
	dcs := &DeliveryCarriers{}
	if err := c.Read(DeliveryCarrierModel, ids, nil, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDeliveryCarrier finds delivery.carrier record by querying it with criteria.
func (c *Client) FindDeliveryCarrier(criteria *Criteria) (*DeliveryCarrier, error) {
	dcs := &DeliveryCarriers{}
	if err := c.SearchRead(DeliveryCarrierModel, criteria, NewOptions().Limit(1), dcs); err != nil {
		return nil, err
	}
	if dcs != nil && len(*dcs) > 0 {
		return &((*dcs)[0]), nil
	}
	return nil, fmt.Errorf("delivery.carrier was not found with criteria %v", criteria)
}

// FindDeliveryCarriers finds delivery.carrier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryCarriers(criteria *Criteria, options *Options) (*DeliveryCarriers, error) {
	dcs := &DeliveryCarriers{}
	if err := c.SearchRead(DeliveryCarrierModel, criteria, options, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDeliveryCarrierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryCarrierIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DeliveryCarrierModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDeliveryCarrierId finds record id by querying it with criteria.
func (c *Client) FindDeliveryCarrierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DeliveryCarrierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("delivery.carrier was not found with criteria %v and options %v", criteria, options)
}
