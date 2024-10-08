package odoo

// KnowViewStat represents know.view.stat model.
type KnowViewStat struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Counter     *Int      `xmlrpc:"counter,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowViewStats represents array of know.view.stat model.
type KnowViewStats []KnowViewStat

// KnowViewStatModel is the odoo model name.
const KnowViewStatModel = "know.view.stat"

// Many2One convert KnowViewStat to *Many2One.
func (kvs *KnowViewStat) Many2One() *Many2One {
	return NewMany2One(kvs.Id.Get(), "")
}

// CreateKnowViewStat creates a new know.view.stat model and returns its id.
func (c *Client) CreateKnowViewStat(kvs *KnowViewStat) (int64, error) {
	ids, err := c.CreateKnowViewStats([]*KnowViewStat{kvs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowViewStat creates a new know.view.stat model and returns its id.
func (c *Client) CreateKnowViewStats(kvss []*KnowViewStat) ([]int64, error) {
	var vv []interface{}
	for _, v := range kvss {
		vv = append(vv, v)
	}
	return c.Create(KnowViewStatModel, vv, nil)
}

// UpdateKnowViewStat updates an existing know.view.stat record.
func (c *Client) UpdateKnowViewStat(kvs *KnowViewStat) error {
	return c.UpdateKnowViewStats([]int64{kvs.Id.Get()}, kvs)
}

// UpdateKnowViewStats updates existing know.view.stat records.
// All records (represented by ids) will be updated by kvs values.
func (c *Client) UpdateKnowViewStats(ids []int64, kvs *KnowViewStat) error {
	return c.Update(KnowViewStatModel, ids, kvs, nil)
}

// DeleteKnowViewStat deletes an existing know.view.stat record.
func (c *Client) DeleteKnowViewStat(id int64) error {
	return c.DeleteKnowViewStats([]int64{id})
}

// DeleteKnowViewStats deletes existing know.view.stat records.
func (c *Client) DeleteKnowViewStats(ids []int64) error {
	return c.Delete(KnowViewStatModel, ids)
}

// GetKnowViewStat gets know.view.stat existing record.
func (c *Client) GetKnowViewStat(id int64) (*KnowViewStat, error) {
	kvss, err := c.GetKnowViewStats([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kvss)[0]), nil
}

// GetKnowViewStats gets know.view.stat existing records.
func (c *Client) GetKnowViewStats(ids []int64) (*KnowViewStats, error) {
	kvss := &KnowViewStats{}
	if err := c.Read(KnowViewStatModel, ids, nil, kvss); err != nil {
		return nil, err
	}
	return kvss, nil
}

// FindKnowViewStat finds know.view.stat record by querying it with criteria.
func (c *Client) FindKnowViewStat(criteria *Criteria) (*KnowViewStat, error) {
	kvss := &KnowViewStats{}
	if err := c.SearchRead(KnowViewStatModel, criteria, NewOptions().Limit(1), kvss); err != nil {
		return nil, err
	}
	return &((*kvss)[0]), nil
}

// FindKnowViewStats finds know.view.stat records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowViewStats(criteria *Criteria, options *Options) (*KnowViewStats, error) {
	kvss := &KnowViewStats{}
	if err := c.SearchRead(KnowViewStatModel, criteria, options, kvss); err != nil {
		return nil, err
	}
	return kvss, nil
}

// FindKnowViewStatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowViewStatIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowViewStatModel, criteria, options)
}

// FindKnowViewStatId finds record id by querying it with criteria.
func (c *Client) FindKnowViewStatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowViewStatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
