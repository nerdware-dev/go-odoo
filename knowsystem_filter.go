package odoo

import (
	"fmt"
)

// KnowsystemFilter represents knowsystem.filter model.
type KnowsystemFilter struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Domain      *String    `xmlrpc:"domain,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Model       *Selection `xmlrpc:"model,omptempty"`
	TagId       *Many2One  `xmlrpc:"tag_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(KnowsystemFilterModel, vv)
}

// UpdateKnowsystemFilter updates an existing knowsystem.filter record.
func (c *Client) UpdateKnowsystemFilter(kf *KnowsystemFilter) error {
	return c.UpdateKnowsystemFilters([]int64{kf.Id.Get()}, kf)
}

// UpdateKnowsystemFilters updates existing knowsystem.filter records.
// All records (represented by ids) will be updated by kf values.
func (c *Client) UpdateKnowsystemFilters(ids []int64, kf *KnowsystemFilter) error {
	return c.Update(KnowsystemFilterModel, ids, kf)
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
	if kfs != nil && len(*kfs) > 0 {
		return &((*kfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.filter not found", id)
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
	if kfs != nil && len(*kfs) > 0 {
		return &((*kfs)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.filter was not found with criteria %v", criteria)
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
	ids, err := c.Search(KnowsystemFilterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemFilterId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.filter was not found with criteria %v and options %v", criteria, options)
}
