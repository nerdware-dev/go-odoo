package odoo

import (
	"fmt"
)

// KnowsystemArticleRevision represents knowsystem.article.revision model.
type KnowsystemArticleRevision struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	ArticleId         *Many2One `xmlrpc:"article_id,omptempty"`
	AttachmentIds     *Relation `xmlrpc:"attachment_ids,omptempty"`
	AttachmentsChange *Bool     `xmlrpc:"attachments_change,omptempty"`
	AuthorId          *Many2One `xmlrpc:"author_id,omptempty"`
	ChangeDatetime    *Time     `xmlrpc:"change_datetime,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	Description       *String   `xmlrpc:"description,omptempty"`
	DescriptionChange *Int      `xmlrpc:"description_change,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	Name              *String   `xmlrpc:"name,omptempty"`
	NameChange        *Bool     `xmlrpc:"name_change,omptempty"`
	SectionChange     *Bool     `xmlrpc:"section_change,omptempty"`
	SectionId         *Many2One `xmlrpc:"section_id,omptempty"`
	TagIds            *Relation `xmlrpc:"tag_ids,omptempty"`
	TagsChange        *Bool     `xmlrpc:"tags_change,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemArticleRevisions represents array of knowsystem.article.revision model.
type KnowsystemArticleRevisions []KnowsystemArticleRevision

// KnowsystemArticleRevisionModel is the odoo model name.
const KnowsystemArticleRevisionModel = "knowsystem.article.revision"

// Many2One convert KnowsystemArticleRevision to *Many2One.
func (kar *KnowsystemArticleRevision) Many2One() *Many2One {
	return NewMany2One(kar.Id.Get(), "")
}

// CreateKnowsystemArticleRevision creates a new knowsystem.article.revision model and returns its id.
func (c *Client) CreateKnowsystemArticleRevision(kar *KnowsystemArticleRevision) (int64, error) {
	ids, err := c.CreateKnowsystemArticleRevisions([]*KnowsystemArticleRevision{kar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemArticleRevision creates a new knowsystem.article.revision model and returns its id.
func (c *Client) CreateKnowsystemArticleRevisions(kars []*KnowsystemArticleRevision) ([]int64, error) {
	var vv []interface{}
	for _, v := range kars {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemArticleRevisionModel, vv)
}

// UpdateKnowsystemArticleRevision updates an existing knowsystem.article.revision record.
func (c *Client) UpdateKnowsystemArticleRevision(kar *KnowsystemArticleRevision) error {
	return c.UpdateKnowsystemArticleRevisions([]int64{kar.Id.Get()}, kar)
}

// UpdateKnowsystemArticleRevisions updates existing knowsystem.article.revision records.
// All records (represented by ids) will be updated by kar values.
func (c *Client) UpdateKnowsystemArticleRevisions(ids []int64, kar *KnowsystemArticleRevision) error {
	return c.Update(KnowsystemArticleRevisionModel, ids, kar)
}

// DeleteKnowsystemArticleRevision deletes an existing knowsystem.article.revision record.
func (c *Client) DeleteKnowsystemArticleRevision(id int64) error {
	return c.DeleteKnowsystemArticleRevisions([]int64{id})
}

// DeleteKnowsystemArticleRevisions deletes existing knowsystem.article.revision records.
func (c *Client) DeleteKnowsystemArticleRevisions(ids []int64) error {
	return c.Delete(KnowsystemArticleRevisionModel, ids)
}

// GetKnowsystemArticleRevision gets knowsystem.article.revision existing record.
func (c *Client) GetKnowsystemArticleRevision(id int64) (*KnowsystemArticleRevision, error) {
	kars, err := c.GetKnowsystemArticleRevisions([]int64{id})
	if err != nil {
		return nil, err
	}
	if kars != nil && len(*kars) > 0 {
		return &((*kars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.article.revision not found", id)
}

// GetKnowsystemArticleRevisions gets knowsystem.article.revision existing records.
func (c *Client) GetKnowsystemArticleRevisions(ids []int64) (*KnowsystemArticleRevisions, error) {
	kars := &KnowsystemArticleRevisions{}
	if err := c.Read(KnowsystemArticleRevisionModel, ids, nil, kars); err != nil {
		return nil, err
	}
	return kars, nil
}

// FindKnowsystemArticleRevision finds knowsystem.article.revision record by querying it with criteria.
func (c *Client) FindKnowsystemArticleRevision(criteria *Criteria) (*KnowsystemArticleRevision, error) {
	kars := &KnowsystemArticleRevisions{}
	if err := c.SearchRead(KnowsystemArticleRevisionModel, criteria, NewOptions().Limit(1), kars); err != nil {
		return nil, err
	}
	if kars != nil && len(*kars) > 0 {
		return &((*kars)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.article.revision was not found with criteria %v", criteria)
}

// FindKnowsystemArticleRevisions finds knowsystem.article.revision records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleRevisions(criteria *Criteria, options *Options) (*KnowsystemArticleRevisions, error) {
	kars := &KnowsystemArticleRevisions{}
	if err := c.SearchRead(KnowsystemArticleRevisionModel, criteria, options, kars); err != nil {
		return nil, err
	}
	return kars, nil
}

// FindKnowsystemArticleRevisionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleRevisionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemArticleRevisionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemArticleRevisionId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemArticleRevisionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemArticleRevisionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.article.revision was not found with criteria %v and options %v", criteria, options)
}
