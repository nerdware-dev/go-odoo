package odoo

// L10NEuServiceWizard represents l10n_eu_service.wizard model.
type L10NEuServiceWizard struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	AccountCollectedId *Many2One `xmlrpc:"account_collected_id,omitempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	DoneCountryIds     *Relation `xmlrpc:"done_country_ids,omitempty"`
	FiscalPositionId   *Many2One `xmlrpc:"fiscal_position_id,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	TaxId              *Many2One `xmlrpc:"tax_id,omitempty"`
	TodoCountryIds     *Relation `xmlrpc:"todo_country_ids,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// L10NEuServiceWizards represents array of l10n_eu_service.wizard model.
type L10NEuServiceWizards []L10NEuServiceWizard

// L10NEuServiceWizardModel is the odoo model name.
const L10NEuServiceWizardModel = "l10n_eu_service.wizard"

// Many2One convert L10NEuServiceWizard to *Many2One.
func (lw *L10NEuServiceWizard) Many2One() *Many2One {
	return NewMany2One(lw.Id.Get(), "")
}

// CreateL10NEuServiceWizard creates a new l10n_eu_service.wizard model and returns its id.
func (c *Client) CreateL10NEuServiceWizard(lw *L10NEuServiceWizard) (int64, error) {
	ids, err := c.CreateL10NEuServiceWizards([]*L10NEuServiceWizard{lw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NEuServiceWizard creates a new l10n_eu_service.wizard model and returns its id.
func (c *Client) CreateL10NEuServiceWizards(lws []*L10NEuServiceWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lws {
		vv = append(vv, v)
	}
	return c.Create(L10NEuServiceWizardModel, vv, nil)
}

// UpdateL10NEuServiceWizard updates an existing l10n_eu_service.wizard record.
func (c *Client) UpdateL10NEuServiceWizard(lw *L10NEuServiceWizard) error {
	return c.UpdateL10NEuServiceWizards([]int64{lw.Id.Get()}, lw)
}

// UpdateL10NEuServiceWizards updates existing l10n_eu_service.wizard records.
// All records (represented by ids) will be updated by lw values.
func (c *Client) UpdateL10NEuServiceWizards(ids []int64, lw *L10NEuServiceWizard) error {
	return c.Update(L10NEuServiceWizardModel, ids, lw, nil)
}

// DeleteL10NEuServiceWizard deletes an existing l10n_eu_service.wizard record.
func (c *Client) DeleteL10NEuServiceWizard(id int64) error {
	return c.DeleteL10NEuServiceWizards([]int64{id})
}

// DeleteL10NEuServiceWizards deletes existing l10n_eu_service.wizard records.
func (c *Client) DeleteL10NEuServiceWizards(ids []int64) error {
	return c.Delete(L10NEuServiceWizardModel, ids)
}

// GetL10NEuServiceWizard gets l10n_eu_service.wizard existing record.
func (c *Client) GetL10NEuServiceWizard(id int64) (*L10NEuServiceWizard, error) {
	lws, err := c.GetL10NEuServiceWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lws)[0]), nil
}

// GetL10NEuServiceWizards gets l10n_eu_service.wizard existing records.
func (c *Client) GetL10NEuServiceWizards(ids []int64) (*L10NEuServiceWizards, error) {
	lws := &L10NEuServiceWizards{}
	if err := c.Read(L10NEuServiceWizardModel, ids, nil, lws); err != nil {
		return nil, err
	}
	return lws, nil
}

// FindL10NEuServiceWizard finds l10n_eu_service.wizard record by querying it with criteria.
func (c *Client) FindL10NEuServiceWizard(criteria *Criteria) (*L10NEuServiceWizard, error) {
	lws := &L10NEuServiceWizards{}
	if err := c.SearchRead(L10NEuServiceWizardModel, criteria, NewOptions().Limit(1), lws); err != nil {
		return nil, err
	}
	return &((*lws)[0]), nil
}

// FindL10NEuServiceWizards finds l10n_eu_service.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEuServiceWizards(criteria *Criteria, options *Options) (*L10NEuServiceWizards, error) {
	lws := &L10NEuServiceWizards{}
	if err := c.SearchRead(L10NEuServiceWizardModel, criteria, options, lws); err != nil {
		return nil, err
	}
	return lws, nil
}

// FindL10NEuServiceWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NEuServiceWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NEuServiceWizardModel, criteria, options)
}

// FindL10NEuServiceWizardId finds record id by querying it with criteria.
func (c *Client) FindL10NEuServiceWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NEuServiceWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
