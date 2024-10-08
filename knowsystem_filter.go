package odoo

// KnowsystemFilter represents knowsystem.filter model.
type KnowsystemFilter struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Domain      *String    `xmlrpc:"domain,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Model       *Selection `xmlrpc:"model,omitempty"`
	TagId       *Many2One  `xmlrpc:"tag_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// KnowsystemFilters represents array of knowsystem.filter model.
type KnowsystemFilters []KnowsystemFilter

// KnowsystemFilterModel is the odoo model name.
const KnowsystemFilterModel = "knowsystem.filter"

// Many2One convert KnowsystemFilter to *Many2One.
func (kf *KnowsystemFilter) Many2One() *Many2One {
	return NewMany2One(kf.Id.Get(), "")
}

// CreateKnowsystemFilter creates a new knowsystem.filter model and returns its id.
func (c *Client) CreateKnowsystemFilter(kf *KnowsystemFilter) (int64, error) {
	ids, err := c.CreateKnowsystemFilters([]*KnowsystemFilter{kf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemFilter creates a new knowsystem.filter model and returns its id.
func (c *Client) CreateKnowsystemFilters(kfs []*KnowsystemFilter) ([]int64, error) {
	var vv []interface{}
	for _, v := range kfs {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemFilterModel, vv, nil)
}

// UpdateKnowsystemFilter updates an existing knowsystem.filter record.
func (c *Client) UpdateKnowsystemFilter(kf *KnowsystemFilter) error {
	return c.UpdateKnowsystemFilters([]int64{kf.Id.Get()}, kf)
}

// UpdateKnowsystemFilters updates existing knowsystem.filter records.
// All records (represented by ids) will be updated by kf values.
func (c *Client) UpdateKnowsystemFilters(ids []int64, kf *KnowsystemFilter) error {
	return c.Update(KnowsystemFilterModel, ids, kf, nil)
}

// DeleteKnowsystemFilter deletes an existing knowsystem.filter record.
func (c *Client) DeleteKnowsystemFilter(id int64) error {
	return c.DeleteKnowsystemFilters([]int64{id})
}

// DeleteKnowsystemFilters deletes existing knowsystem.filter records.
func (c *Client) DeleteKnowsystemFilters(ids []int64) error {
	return c.Delete(KnowsystemFilterModel, ids)
}

// GetKnowsystemFilter gets knowsystem.filter existing record.
func (c *Client) GetKnowsystemFilter(id int64) (*KnowsystemFilter, error) {
	kfs, err := c.GetKnowsystemFilters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kfs)[0]), nil
}

// GetKnowsystemFilters gets knowsystem.filter existing records.
func (c *Client) GetKnowsystemFilters(ids []int64) (*KnowsystemFilters, error) {
	kfs := &KnowsystemFilters{}
	if err := c.Read(KnowsystemFilterModel, ids, nil, kfs); err != nil {
		return nil, err
	}
	return kfs, nil
}

// FindKnowsystemFilter finds knowsystem.filter record by querying it with criteria.
func (c *Client) FindKnowsystemFilter(criteria *Criteria) (*KnowsystemFilter, error) {
	kfs := &KnowsystemFilters{}
	if err := c.SearchRead(KnowsystemFilterModel, criteria, NewOptions().Limit(1), kfs); err != nil {
		return nil, err
	}
	return &((*kfs)[0]), nil
}

// FindKnowsystemFilters finds knowsystem.filter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemFilters(criteria *Criteria, options *Options) (*KnowsystemFilters, error) {
	kfs := &KnowsystemFilters{}
	if err := c.SearchRead(KnowsystemFilterModel, criteria, options, kfs); err != nil {
		return nil, err
	}
	return kfs, nil
}

// FindKnowsystemFilterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemFilterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowsystemFilterModel, criteria, options)
}

// FindKnowsystemFilterId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
