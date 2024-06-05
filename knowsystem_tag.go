package odoo

import (
	"fmt"
)

// KnowsystemTag represents knowsystem.tag model.
type KnowsystemTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	ApplyToAll  *Bool     `xmlrpc:"apply_to_all,omptempty"`
	ArticleIds  *Relation `xmlrpc:"article_ids,omptempty"`
	CanPublish  *Bool     `xmlrpc:"can_publish,omptempty"`
	ChildIds    *Relation `xmlrpc:"child_ids,omptempty"`
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Description *String   `xmlrpc:"description,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FilterIds   *Relation `xmlrpc:"filter_ids,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	IsPublished *Bool     `xmlrpc:"is_published,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WebsiteId   *Many2One `xmlrpc:"website_id,omptempty"`
	WebsiteUrl  *String   `xmlrpc:"website_url,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemTags represents array of knowsystem.tag model.
type KnowsystemTags []KnowsystemTag

// KnowsystemTagModel is the odoo model name.
const KnowsystemTagModel = "knowsystem.tag"

// Many2One convert KnowsystemTag to *Many2One.
func (kt *KnowsystemTag) Many2One() *Many2One {
	return NewMany2One(kt.Id.Get(), "")
}

// CreateKnowsystemTag creates a new knowsystem.tag model and returns its id.
func (c *Client) CreateKnowsystemTag(kt *KnowsystemTag) (int64, error) {
	ids, err := c.CreateKnowsystemTags([]*KnowsystemTag{kt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemTag creates a new knowsystem.tag model and returns its id.
func (c *Client) CreateKnowsystemTags(kts []*KnowsystemTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range kts {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemTagModel, vv)
}

// UpdateKnowsystemTag updates an existing knowsystem.tag record.
func (c *Client) UpdateKnowsystemTag(kt *KnowsystemTag) error {
	return c.UpdateKnowsystemTags([]int64{kt.Id.Get()}, kt)
}

// UpdateKnowsystemTags updates existing knowsystem.tag records.
// All records (represented by ids) will be updated by kt values.
func (c *Client) UpdateKnowsystemTags(ids []int64, kt *KnowsystemTag) error {
	return c.Update(KnowsystemTagModel, ids, kt)
}

// DeleteKnowsystemTag deletes an existing knowsystem.tag record.
func (c *Client) DeleteKnowsystemTag(id int64) error {
	return c.DeleteKnowsystemTags([]int64{id})
}

// DeleteKnowsystemTags deletes existing knowsystem.tag records.
func (c *Client) DeleteKnowsystemTags(ids []int64) error {
	return c.Delete(KnowsystemTagModel, ids)
}

// GetKnowsystemTag gets knowsystem.tag existing record.
func (c *Client) GetKnowsystemTag(id int64) (*KnowsystemTag, error) {
	kts, err := c.GetKnowsystemTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if kts != nil && len(*kts) > 0 {
		return &((*kts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.tag not found", id)
}

// GetKnowsystemTags gets knowsystem.tag existing records.
func (c *Client) GetKnowsystemTags(ids []int64) (*KnowsystemTags, error) {
	kts := &KnowsystemTags{}
	if err := c.Read(KnowsystemTagModel, ids, nil, kts); err != nil {
		return nil, err
	}
	return kts, nil
}

// FindKnowsystemTag finds knowsystem.tag record by querying it with criteria.
func (c *Client) FindKnowsystemTag(criteria *Criteria) (*KnowsystemTag, error) {
	kts := &KnowsystemTags{}
	if err := c.SearchRead(KnowsystemTagModel, criteria, NewOptions().Limit(1), kts); err != nil {
		return nil, err
	}
	if kts != nil && len(*kts) > 0 {
		return &((*kts)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.tag was not found with criteria %v", criteria)
}

// FindKnowsystemTags finds knowsystem.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTags(criteria *Criteria, options *Options) (*KnowsystemTags, error) {
	kts := &KnowsystemTags{}
	if err := c.SearchRead(KnowsystemTagModel, criteria, options, kts); err != nil {
		return nil, err
	}
	return kts, nil
}

// FindKnowsystemTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemTagId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.tag was not found with criteria %v and options %v", criteria, options)
}
