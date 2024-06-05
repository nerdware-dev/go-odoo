package odoo

import (
	"fmt"
)

// KnowsystemArticleTemplate represents knowsystem.article.template model.
type KnowsystemArticleTemplate struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Active          *Bool     `xmlrpc:"active,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	Description     *String   `xmlrpc:"description,omptempty"`
	DescriptionArch *String   `xmlrpc:"description_arch,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	Sequence        *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemArticleTemplates represents array of knowsystem.article.template model.
type KnowsystemArticleTemplates []KnowsystemArticleTemplate

// KnowsystemArticleTemplateModel is the odoo model name.
const KnowsystemArticleTemplateModel = "knowsystem.article.template"

// Many2One convert KnowsystemArticleTemplate to *Many2One.
func (kat *KnowsystemArticleTemplate) Many2One() *Many2One {
	return NewMany2One(kat.Id.Get(), "")
}

// CreateKnowsystemArticleTemplate creates a new knowsystem.article.template model and returns its id.
func (c *Client) CreateKnowsystemArticleTemplate(kat *KnowsystemArticleTemplate) (int64, error) {
	ids, err := c.CreateKnowsystemArticleTemplates([]*KnowsystemArticleTemplate{kat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemArticleTemplate creates a new knowsystem.article.template model and returns its id.
func (c *Client) CreateKnowsystemArticleTemplates(kats []*KnowsystemArticleTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range kats {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemArticleTemplateModel, vv)
}

// UpdateKnowsystemArticleTemplate updates an existing knowsystem.article.template record.
func (c *Client) UpdateKnowsystemArticleTemplate(kat *KnowsystemArticleTemplate) error {
	return c.UpdateKnowsystemArticleTemplates([]int64{kat.Id.Get()}, kat)
}

// UpdateKnowsystemArticleTemplates updates existing knowsystem.article.template records.
// All records (represented by ids) will be updated by kat values.
func (c *Client) UpdateKnowsystemArticleTemplates(ids []int64, kat *KnowsystemArticleTemplate) error {
	return c.Update(KnowsystemArticleTemplateModel, ids, kat)
}

// DeleteKnowsystemArticleTemplate deletes an existing knowsystem.article.template record.
func (c *Client) DeleteKnowsystemArticleTemplate(id int64) error {
	return c.DeleteKnowsystemArticleTemplates([]int64{id})
}

// DeleteKnowsystemArticleTemplates deletes existing knowsystem.article.template records.
func (c *Client) DeleteKnowsystemArticleTemplates(ids []int64) error {
	return c.Delete(KnowsystemArticleTemplateModel, ids)
}

// GetKnowsystemArticleTemplate gets knowsystem.article.template existing record.
func (c *Client) GetKnowsystemArticleTemplate(id int64) (*KnowsystemArticleTemplate, error) {
	kats, err := c.GetKnowsystemArticleTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if kats != nil && len(*kats) > 0 {
		return &((*kats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.article.template not found", id)
}

// GetKnowsystemArticleTemplates gets knowsystem.article.template existing records.
func (c *Client) GetKnowsystemArticleTemplates(ids []int64) (*KnowsystemArticleTemplates, error) {
	kats := &KnowsystemArticleTemplates{}
	if err := c.Read(KnowsystemArticleTemplateModel, ids, nil, kats); err != nil {
		return nil, err
	}
	return kats, nil
}

// FindKnowsystemArticleTemplate finds knowsystem.article.template record by querying it with criteria.
func (c *Client) FindKnowsystemArticleTemplate(criteria *Criteria) (*KnowsystemArticleTemplate, error) {
	kats := &KnowsystemArticleTemplates{}
	if err := c.SearchRead(KnowsystemArticleTemplateModel, criteria, NewOptions().Limit(1), kats); err != nil {
		return nil, err
	}
	if kats != nil && len(*kats) > 0 {
		return &((*kats)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.article.template was not found with criteria %v", criteria)
}

// FindKnowsystemArticleTemplates finds knowsystem.article.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleTemplates(criteria *Criteria, options *Options) (*KnowsystemArticleTemplates, error) {
	kats := &KnowsystemArticleTemplates{}
	if err := c.SearchRead(KnowsystemArticleTemplateModel, criteria, options, kats); err != nil {
		return nil, err
	}
	return kats, nil
}

// FindKnowsystemArticleTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemArticleTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemArticleTemplateId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemArticleTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemArticleTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.article.template was not found with criteria %v and options %v", criteria, options)
}
