package odoo

import (
	"fmt"
)

// L10NEuServiceServiceTaxRate represents l10n_eu_service.service_tax_rate model.
type L10NEuServiceServiceTaxRate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Rate        *Float    `xmlrpc:"rate,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// L10NEuServiceServiceTaxRates represents array of l10n_eu_service.service_tax_rate model.
type L10NEuServiceServiceTaxRates []L10NEuServiceServiceTaxRate

// L10NEuServiceServiceTaxRateModel is the odoo model name.
const L10NEuServiceServiceTaxRateModel = "l10n_eu_service.service_tax_rate"

// Many2One convert L10NEuServiceServiceTaxRate to *Many2One.
func (ls *L10NEuServiceServiceTaxRate) Many2One() *Many2One {
	return NewMany2One(ls.Id.Get(), "")
}

// CreateL10NEuServiceServiceTaxRate creates a new l10n_eu_service.service_tax_rate model and returns its id.
func (c *Client) CreateL10NEuServiceServiceTaxRate(ls *L10NEuServiceServiceTaxRate) (int64, error) {
	ids, err := c.CreateL10NEuServiceServiceTaxRates([]*L10NEuServiceServiceTaxRate{ls})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NEuServiceServiceTaxRate creates a new l10n_eu_service.service_tax_rate model and returns its id.
func (c *Client) CreateL10NEuServiceServiceTaxRates(lss []*L10NEuServiceServiceTaxRate) ([]int64, error) {
	var vv []interface{}
	for _, v := range lss {
		vv = append(vv, v)
	}
	return c.Create(L10NEuServiceServiceTaxRateModel, vv)
}

// UpdateL10NEuServiceServiceTaxRate updates an existing l10n_eu_service.service_tax_rate record.
func (c *Client) UpdateL10NEuServiceServiceTaxRate(ls *L10NEuServiceServiceTaxRate) error {
	return c.UpdateL10NEuServiceServiceTaxRates([]int64{ls.Id.Get()}, ls)
}

// UpdateL10NEuServiceServiceTaxRates updates existing l10n_eu_service.service_tax_rate records.
// All records (represented by ids) will be updated by ls values.
func (c *Client) UpdateL10NEuServiceServiceTaxRates(ids []int64, ls *L10NEuServiceServiceTaxRate) error {
	return c.Update(L10NEuServiceServiceTaxRateModel, ids, ls)
}

// DeleteL10NEuServiceServiceTaxRate deletes an existing l10n_eu_service.service_tax_rate record.
func (c *Client) DeleteL10NEuServiceServiceTaxRate(id int64) error {
	return c.DeleteL10NEuServiceServiceTaxRates([]int64{id})
}

// DeleteL10NEuServiceServiceTaxRates deletes existing l10n_eu_service.service_tax_rate records.
func (c *Client) DeleteL10NEuServiceServiceTaxRates(ids []int64) error {
	return c.Delete(L10NEuServiceServiceTaxRateModel, ids)
}

// GetL10NEuServiceServiceTaxRate gets l10n_eu_service.service_tax_rate existing record.
func (c *Client) GetL10NEuServiceServiceTaxRate(id int64) (*L10NEuServiceServiceTaxRate, error) {
	lss, err := c.GetL10NEuServiceServiceTaxRates([]int64{id})
	if err != nil {
		return nil, err
	}
	if lss != nil && len(*lss) > 0 {
		return &((*lss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of l10n_eu_service.service_tax_rate not found", id)
}

// GetL10NEuServiceServiceTaxRates gets l10n_eu_service.service_tax_rate existing records.
func (c *Client) GetL10NEuServiceServiceTaxRates(ids []int64) (*L10NEuServiceServiceTaxRates, error) {
	lss := &L10NEuServiceServiceTaxRates{}
	if err := c.Read(L10NEuServiceServiceTaxRateModel, ids, nil, lss); err != nil {
		return nil, err
	}
	return lss, nil
}

// FindL10NEuServiceServiceTaxRate finds l10n_eu_service.service_tax_rate record by querying it with criteria.
func (c *Client) FindL10NEuServiceServiceTaxRate(criteria *Criteria) (*L10NEuServiceServiceTaxRate, error) {
	lss := &L10NEuServiceServiceTaxRates{}
	if err := c.SearchRead(L10NEuServiceServiceTaxRateModel, criteria, NewOptions().Limit(1), lss); err != nil {
		return nil, err
	}
	if lss != nil && len(*lss) > 0 {
		return &((*lss)[0]), nil
	}
	return nil, fmt.Errorf("l10n_eu_service.service_tax_rate was not found with criteria %v", criteria)
}

// FindL10NEuServiceServiceTaxRates finds l10n_eu_service.service_tax_rate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEuServiceServiceTaxRates(criteria *Criteria, options *Options) (*L10NEuServiceServiceTaxRates, error) {
	lss := &L10NEuServiceServiceTaxRates{}
	if err := c.SearchRead(L10NEuServiceServiceTaxRateModel, criteria, options, lss); err != nil {
		return nil, err
	}
	return lss, nil
}

// FindL10NEuServiceServiceTaxRateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEuServiceServiceTaxRateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(L10NEuServiceServiceTaxRateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindL10NEuServiceServiceTaxRateId finds record id by querying it with criteria.
func (c *Client) FindL10NEuServiceServiceTaxRateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NEuServiceServiceTaxRateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("l10n_eu_service.service_tax_rate was not found with criteria %v and options %v", criteria, options)
}
