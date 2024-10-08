package odoo

// SaleSubscriptionWizardOption represents sale.subscription.wizard.option model.
type SaleSubscriptionWizardOption struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omitempty"`
	Quantity             *Float    `xmlrpc:"quantity,omitempty"`
	UomId                *Many2One `xmlrpc:"uom_id,omitempty"`
	WizardId             *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleSubscriptionWizardOptions represents array of sale.subscription.wizard.option model.
type SaleSubscriptionWizardOptions []SaleSubscriptionWizardOption

// SaleSubscriptionWizardOptionModel is the odoo model name.
const SaleSubscriptionWizardOptionModel = "sale.subscription.wizard.option"

// Many2One convert SaleSubscriptionWizardOption to *Many2One.
func (sswo *SaleSubscriptionWizardOption) Many2One() *Many2One {
	return NewMany2One(sswo.Id.Get(), "")
}

// CreateSaleSubscriptionWizardOption creates a new sale.subscription.wizard.option model and returns its id.
func (c *Client) CreateSaleSubscriptionWizardOption(sswo *SaleSubscriptionWizardOption) (int64, error) {
	ids, err := c.CreateSaleSubscriptionWizardOptions([]*SaleSubscriptionWizardOption{sswo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionWizardOption creates a new sale.subscription.wizard.option model and returns its id.
func (c *Client) CreateSaleSubscriptionWizardOptions(sswos []*SaleSubscriptionWizardOption) ([]int64, error) {
	var vv []interface{}
	for _, v := range sswos {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionWizardOptionModel, vv, nil)
}

// UpdateSaleSubscriptionWizardOption updates an existing sale.subscription.wizard.option record.
func (c *Client) UpdateSaleSubscriptionWizardOption(sswo *SaleSubscriptionWizardOption) error {
	return c.UpdateSaleSubscriptionWizardOptions([]int64{sswo.Id.Get()}, sswo)
}

// UpdateSaleSubscriptionWizardOptions updates existing sale.subscription.wizard.option records.
// All records (represented by ids) will be updated by sswo values.
func (c *Client) UpdateSaleSubscriptionWizardOptions(ids []int64, sswo *SaleSubscriptionWizardOption) error {
	return c.Update(SaleSubscriptionWizardOptionModel, ids, sswo, nil)
}

// DeleteSaleSubscriptionWizardOption deletes an existing sale.subscription.wizard.option record.
func (c *Client) DeleteSaleSubscriptionWizardOption(id int64) error {
	return c.DeleteSaleSubscriptionWizardOptions([]int64{id})
}

// DeleteSaleSubscriptionWizardOptions deletes existing sale.subscription.wizard.option records.
func (c *Client) DeleteSaleSubscriptionWizardOptions(ids []int64) error {
	return c.Delete(SaleSubscriptionWizardOptionModel, ids)
}

// GetSaleSubscriptionWizardOption gets sale.subscription.wizard.option existing record.
func (c *Client) GetSaleSubscriptionWizardOption(id int64) (*SaleSubscriptionWizardOption, error) {
	sswos, err := c.GetSaleSubscriptionWizardOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sswos)[0]), nil
}

// GetSaleSubscriptionWizardOptions gets sale.subscription.wizard.option existing records.
func (c *Client) GetSaleSubscriptionWizardOptions(ids []int64) (*SaleSubscriptionWizardOptions, error) {
	sswos := &SaleSubscriptionWizardOptions{}
	if err := c.Read(SaleSubscriptionWizardOptionModel, ids, nil, sswos); err != nil {
		return nil, err
	}
	return sswos, nil
}

// FindSaleSubscriptionWizardOption finds sale.subscription.wizard.option record by querying it with criteria.
func (c *Client) FindSaleSubscriptionWizardOption(criteria *Criteria) (*SaleSubscriptionWizardOption, error) {
	sswos := &SaleSubscriptionWizardOptions{}
	if err := c.SearchRead(SaleSubscriptionWizardOptionModel, criteria, NewOptions().Limit(1), sswos); err != nil {
		return nil, err
	}
	return &((*sswos)[0]), nil
}

// FindSaleSubscriptionWizardOptions finds sale.subscription.wizard.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionWizardOptions(criteria *Criteria, options *Options) (*SaleSubscriptionWizardOptions, error) {
	sswos := &SaleSubscriptionWizardOptions{}
	if err := c.SearchRead(SaleSubscriptionWizardOptionModel, criteria, options, sswos); err != nil {
		return nil, err
	}
	return sswos, nil
}

// FindSaleSubscriptionWizardOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionWizardOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleSubscriptionWizardOptionModel, criteria, options)
}

// FindSaleSubscriptionWizardOptionId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionWizardOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionWizardOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
