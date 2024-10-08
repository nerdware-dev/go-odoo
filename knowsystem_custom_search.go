package odoo

// KnowsystemCustomSearch represents knowsystem.custom.search model.
type KnowsystemCustomSearch struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	CustomFieldId *Many2One `xmlrpc:"custom_field_id,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowsystemCustomSearchs represents array of knowsystem.custom.search model.
type KnowsystemCustomSearchs []KnowsystemCustomSearch

// KnowsystemCustomSearchModel is the odoo model name.
const KnowsystemCustomSearchModel = "knowsystem.custom.search"

// Many2One convert KnowsystemCustomSearch to *Many2One.
func (kcs *KnowsystemCustomSearch) Many2One() *Many2One {
	return NewMany2One(kcs.Id.Get(), "")
}

// CreateKnowsystemCustomSearch creates a new knowsystem.custom.search model and returns its id.
func (c *Client) CreateKnowsystemCustomSearch(kcs *KnowsystemCustomSearch) (int64, error) {
	ids, err := c.CreateKnowsystemCustomSearchs([]*KnowsystemCustomSearch{kcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemCustomSearch creates a new knowsystem.custom.search model and returns its id.
func (c *Client) CreateKnowsystemCustomSearchs(kcss []*KnowsystemCustomSearch) ([]int64, error) {
	var vv []interface{}
	for _, v := range kcss {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemCustomSearchModel, vv, nil)
}

// UpdateKnowsystemCustomSearch updates an existing knowsystem.custom.search record.
func (c *Client) UpdateKnowsystemCustomSearch(kcs *KnowsystemCustomSearch) error {
	return c.UpdateKnowsystemCustomSearchs([]int64{kcs.Id.Get()}, kcs)
}

// UpdateKnowsystemCustomSearchs updates existing knowsystem.custom.search records.
// All records (represented by ids) will be updated by kcs values.
func (c *Client) UpdateKnowsystemCustomSearchs(ids []int64, kcs *KnowsystemCustomSearch) error {
	return c.Update(KnowsystemCustomSearchModel, ids, kcs, nil)
}

// DeleteKnowsystemCustomSearch deletes an existing knowsystem.custom.search record.
func (c *Client) DeleteKnowsystemCustomSearch(id int64) error {
	return c.DeleteKnowsystemCustomSearchs([]int64{id})
}

// DeleteKnowsystemCustomSearchs deletes existing knowsystem.custom.search records.
func (c *Client) DeleteKnowsystemCustomSearchs(ids []int64) error {
	return c.Delete(KnowsystemCustomSearchModel, ids)
}

// GetKnowsystemCustomSearch gets knowsystem.custom.search existing record.
func (c *Client) GetKnowsystemCustomSearch(id int64) (*KnowsystemCustomSearch, error) {
	kcss, err := c.GetKnowsystemCustomSearchs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kcss)[0]), nil
}

// GetKnowsystemCustomSearchs gets knowsystem.custom.search existing records.
func (c *Client) GetKnowsystemCustomSearchs(ids []int64) (*KnowsystemCustomSearchs, error) {
	kcss := &KnowsystemCustomSearchs{}
	if err := c.Read(KnowsystemCustomSearchModel, ids, nil, kcss); err != nil {
		return nil, err
	}
	return kcss, nil
}

// FindKnowsystemCustomSearch finds knowsystem.custom.search record by querying it with criteria.
func (c *Client) FindKnowsystemCustomSearch(criteria *Criteria) (*KnowsystemCustomSearch, error) {
	kcss := &KnowsystemCustomSearchs{}
	if err := c.SearchRead(KnowsystemCustomSearchModel, criteria, NewOptions().Limit(1), kcss); err != nil {
		return nil, err
	}
	return &((*kcss)[0]), nil
}

// FindKnowsystemCustomSearchs finds knowsystem.custom.search records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemCustomSearchs(criteria *Criteria, options *Options) (*KnowsystemCustomSearchs, error) {
	kcss := &KnowsystemCustomSearchs{}
	if err := c.SearchRead(KnowsystemCustomSearchModel, criteria, options, kcss); err != nil {
		return nil, err
	}
	return kcss, nil
}

// FindKnowsystemCustomSearchIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemCustomSearchIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowsystemCustomSearchModel, criteria, options)
}

// FindKnowsystemCustomSearchId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemCustomSearchId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemCustomSearchModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
