package odoo

import (
	"fmt"
)

// KnowsystemSection represents knowsystem.section model.
type KnowsystemSection struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	Active       *Bool     `xmlrpc:"active,omptempty"`
	ArticleIds   *Relation `xmlrpc:"article_ids,omptempty"`
	ChildIds     *Relation `xmlrpc:"child_ids,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	Description  *String   `xmlrpc:"description,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	HasRightTo   *Relation `xmlrpc:"has_right_to,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	ParentId     *Many2One `xmlrpc:"parent_id,omptempty"`
	Sequence     *Int      `xmlrpc:"sequence,omptempty"`
	UserGroupIds *Relation `xmlrpc:"user_group_ids,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemSections represents array of knowsystem.section model.
type KnowsystemSections []KnowsystemSection

// KnowsystemSectionModel is the odoo model name.
const KnowsystemSectionModel = "knowsystem.section"

// Many2One convert KnowsystemSection to *Many2One.
func (ks *KnowsystemSection) Many2One() *Many2One {
	return NewMany2One(ks.Id.Get(), "")
}

// CreateKnowsystemSection creates a new knowsystem.section model and returns its id.
func (c *Client) CreateKnowsystemSection(ks *KnowsystemSection) (int64, error) {
	ids, err := c.CreateKnowsystemSections([]*KnowsystemSection{ks})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemSection creates a new knowsystem.section model and returns its id.
func (c *Client) CreateKnowsystemSections(kss []*KnowsystemSection) ([]int64, error) {
	var vv []interface{}
	for _, v := range kss {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemSectionModel, vv)
}

// UpdateKnowsystemSection updates an existing knowsystem.section record.
func (c *Client) UpdateKnowsystemSection(ks *KnowsystemSection) error {
	return c.UpdateKnowsystemSections([]int64{ks.Id.Get()}, ks)
}

// UpdateKnowsystemSections updates existing knowsystem.section records.
// All records (represented by ids) will be updated by ks values.
func (c *Client) UpdateKnowsystemSections(ids []int64, ks *KnowsystemSection) error {
	return c.Update(KnowsystemSectionModel, ids, ks)
}

// DeleteKnowsystemSection deletes an existing knowsystem.section record.
func (c *Client) DeleteKnowsystemSection(id int64) error {
	return c.DeleteKnowsystemSections([]int64{id})
}

// DeleteKnowsystemSections deletes existing knowsystem.section records.
func (c *Client) DeleteKnowsystemSections(ids []int64) error {
	return c.Delete(KnowsystemSectionModel, ids)
}

// GetKnowsystemSection gets knowsystem.section existing record.
func (c *Client) GetKnowsystemSection(id int64) (*KnowsystemSection, error) {
	kss, err := c.GetKnowsystemSections([]int64{id})
	if err != nil {
		return nil, err
	}
	if kss != nil && len(*kss) > 0 {
		return &((*kss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.section not found", id)
}

// GetKnowsystemSections gets knowsystem.section existing records.
func (c *Client) GetKnowsystemSections(ids []int64) (*KnowsystemSections, error) {
	kss := &KnowsystemSections{}
	if err := c.Read(KnowsystemSectionModel, ids, nil, kss); err != nil {
		return nil, err
	}
	return kss, nil
}

// FindKnowsystemSection finds knowsystem.section record by querying it with criteria.
func (c *Client) FindKnowsystemSection(criteria *Criteria) (*KnowsystemSection, error) {
	kss := &KnowsystemSections{}
	if err := c.SearchRead(KnowsystemSectionModel, criteria, NewOptions().Limit(1), kss); err != nil {
		return nil, err
	}
	if kss != nil && len(*kss) > 0 {
		return &((*kss)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.section was not found with criteria %v", criteria)
}

// FindKnowsystemSections finds knowsystem.section records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemSections(criteria *Criteria, options *Options) (*KnowsystemSections, error) {
	kss := &KnowsystemSections{}
	if err := c.SearchRead(KnowsystemSectionModel, criteria, options, kss); err != nil {
		return nil, err
	}
	return kss, nil
}

// FindKnowsystemSectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemSectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemSectionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemSectionId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemSectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemSectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.section was not found with criteria %v and options %v", criteria, options)
}
