package odoo

// L10NEuServiceServiceTaxRate represents l10n_eu_service.service_tax_rate model.
type L10NEuServiceServiceTaxRate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Rate        *Float    `xmlrpc:"rate,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(L10NEuServiceServiceTaxRateModel, vv, nil)
}

// UpdateL10NEuServiceServiceTaxRate updates an existing l10n_eu_service.service_tax_rate record.
func (c *Client) UpdateL10NEuServiceServiceTaxRate(ls *L10NEuServiceServiceTaxRate) error {
	return c.UpdateL10NEuServiceServiceTaxRates([]int64{ls.Id.Get()}, ls)
}

// UpdateL10NEuServiceServiceTaxRates updates existing l10n_eu_service.service_tax_rate records.
// All records (represented by ids) will be updated by ls values.
func (c *Client) UpdateL10NEuServiceServiceTaxRates(ids []int64, ls *L10NEuServiceServiceTaxRate) error {
	return c.Update(L10NEuServiceServiceTaxRateModel, ids, ls, nil)
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
	return &((*lss)[0]), nil
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
	return &((*lss)[0]), nil
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
	return c.Search(L10NEuServiceServiceTaxRateModel, criteria, options)
}

// FindL10NEuServiceServiceTaxRateId finds record id by querying it with criteria.
func (c *Client) FindL10NEuServiceServiceTaxRateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NEuServiceServiceTaxRateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
