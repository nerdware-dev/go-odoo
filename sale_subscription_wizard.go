package odoo

// SaleSubscriptionWizard represents sale.subscription.wizard model.
type SaleSubscriptionWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DateFrom       *Time     `xmlrpc:"date_from,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	OptionLines    *Relation `xmlrpc:"option_lines,omitempty"`
	SubscriptionId *Many2One `xmlrpc:"subscription_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleSubscriptionWizards represents array of sale.subscription.wizard model.
type SaleSubscriptionWizards []SaleSubscriptionWizard

// SaleSubscriptionWizardModel is the odoo model name.
const SaleSubscriptionWizardModel = "sale.subscription.wizard"

// Many2One convert SaleSubscriptionWizard to *Many2One.
func (ssw *SaleSubscriptionWizard) Many2One() *Many2One {
	return NewMany2One(ssw.Id.Get(), "")
}

// CreateSaleSubscriptionWizard creates a new sale.subscription.wizard model and returns its id.
func (c *Client) CreateSaleSubscriptionWizard(ssw *SaleSubscriptionWizard) (int64, error) {
	ids, err := c.CreateSaleSubscriptionWizards([]*SaleSubscriptionWizard{ssw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionWizard creates a new sale.subscription.wizard model and returns its id.
func (c *Client) CreateSaleSubscriptionWizards(ssws []*SaleSubscriptionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssws {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionWizardModel, vv, nil)
}

// UpdateSaleSubscriptionWizard updates an existing sale.subscription.wizard record.
func (c *Client) UpdateSaleSubscriptionWizard(ssw *SaleSubscriptionWizard) error {
	return c.UpdateSaleSubscriptionWizards([]int64{ssw.Id.Get()}, ssw)
}

// UpdateSaleSubscriptionWizards updates existing sale.subscription.wizard records.
// All records (represented by ids) will be updated by ssw values.
func (c *Client) UpdateSaleSubscriptionWizards(ids []int64, ssw *SaleSubscriptionWizard) error {
	return c.Update(SaleSubscriptionWizardModel, ids, ssw, nil)
}

// DeleteSaleSubscriptionWizard deletes an existing sale.subscription.wizard record.
func (c *Client) DeleteSaleSubscriptionWizard(id int64) error {
	return c.DeleteSaleSubscriptionWizards([]int64{id})
}

// DeleteSaleSubscriptionWizards deletes existing sale.subscription.wizard records.
func (c *Client) DeleteSaleSubscriptionWizards(ids []int64) error {
	return c.Delete(SaleSubscriptionWizardModel, ids)
}

// GetSaleSubscriptionWizard gets sale.subscription.wizard existing record.
func (c *Client) GetSaleSubscriptionWizard(id int64) (*SaleSubscriptionWizard, error) {
	ssws, err := c.GetSaleSubscriptionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssws)[0]), nil
}

// GetSaleSubscriptionWizards gets sale.subscription.wizard existing records.
func (c *Client) GetSaleSubscriptionWizards(ids []int64) (*SaleSubscriptionWizards, error) {
	ssws := &SaleSubscriptionWizards{}
	if err := c.Read(SaleSubscriptionWizardModel, ids, nil, ssws); err != nil {
		return nil, err
	}
	return ssws, nil
}

// FindSaleSubscriptionWizard finds sale.subscription.wizard record by querying it with criteria.
func (c *Client) FindSaleSubscriptionWizard(criteria *Criteria) (*SaleSubscriptionWizard, error) {
	ssws := &SaleSubscriptionWizards{}
	if err := c.SearchRead(SaleSubscriptionWizardModel, criteria, NewOptions().Limit(1), ssws); err != nil {
		return nil, err
	}
	return &((*ssws)[0]), nil
}

// FindSaleSubscriptionWizards finds sale.subscription.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionWizards(criteria *Criteria, options *Options) (*SaleSubscriptionWizards, error) {
	ssws := &SaleSubscriptionWizards{}
	if err := c.SearchRead(SaleSubscriptionWizardModel, criteria, options, ssws); err != nil {
		return nil, err
	}
	return ssws, nil
}

// FindSaleSubscriptionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleSubscriptionWizardModel, criteria, options)
}

// FindSaleSubscriptionWizardId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
