package odoo

import (
	"fmt"
)

// ForumPostReason represents forum.post.reason model.
type ForumPostReason struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ReasonType  *String   `xmlrpc:"reason_type,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ForumPostReasons represents array of forum.post.reason model.
type ForumPostReasons []ForumPostReason

// ForumPostReasonModel is the odoo model name.
const ForumPostReasonModel = "forum.post.reason"

// Many2One convert ForumPostReason to *Many2One.
func (fpr *ForumPostReason) Many2One() *Many2One {
	return NewMany2One(fpr.Id.Get(), "")
}

// CreateForumPostReason creates a new forum.post.reason model and returns its id.
func (c *Client) CreateForumPostReason(fpr *ForumPostReason) (int64, error) {
	ids, err := c.CreateForumPostReasons([]*ForumPostReason{fpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateForumPostReason creates a new forum.post.reason model and returns its id.
func (c *Client) CreateForumPostReasons(fprs []*ForumPostReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range fprs {
		vv = append(vv, v)
	}
	return c.Create(ForumPostReasonModel, vv)
}

// UpdateForumPostReason updates an existing forum.post.reason record.
func (c *Client) UpdateForumPostReason(fpr *ForumPostReason) error {
	return c.UpdateForumPostReasons([]int64{fpr.Id.Get()}, fpr)
}

// UpdateForumPostReasons updates existing forum.post.reason records.
// All records (represented by ids) will be updated by fpr values.
func (c *Client) UpdateForumPostReasons(ids []int64, fpr *ForumPostReason) error {
	return c.Update(ForumPostReasonModel, ids, fpr)
}

// DeleteForumPostReason deletes an existing forum.post.reason record.
func (c *Client) DeleteForumPostReason(id int64) error {
	return c.DeleteForumPostReasons([]int64{id})
}

// DeleteForumPostReasons deletes existing forum.post.reason records.
func (c *Client) DeleteForumPostReasons(ids []int64) error {
	return c.Delete(ForumPostReasonModel, ids)
}

// GetForumPostReason gets forum.post.reason existing record.
func (c *Client) GetForumPostReason(id int64) (*ForumPostReason, error) {
	fprs, err := c.GetForumPostReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	if fprs != nil && len(*fprs) > 0 {
		return &((*fprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of forum.post.reason not found", id)
}

// GetForumPostReasons gets forum.post.reason existing records.
func (c *Client) GetForumPostReasons(ids []int64) (*ForumPostReasons, error) {
	fprs := &ForumPostReasons{}
	if err := c.Read(ForumPostReasonModel, ids, nil, fprs); err != nil {
		return nil, err
	}
	return fprs, nil
}

// FindForumPostReason finds forum.post.reason record by querying it with criteria.
func (c *Client) FindForumPostReason(criteria *Criteria) (*ForumPostReason, error) {
	fprs := &ForumPostReasons{}
	if err := c.SearchRead(ForumPostReasonModel, criteria, NewOptions().Limit(1), fprs); err != nil {
		return nil, err
	}
	if fprs != nil && len(*fprs) > 0 {
		return &((*fprs)[0]), nil
	}
	return nil, fmt.Errorf("forum.post.reason was not found with criteria %v", criteria)
}

// FindForumPostReasons finds forum.post.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumPostReasons(criteria *Criteria, options *Options) (*ForumPostReasons, error) {
	fprs := &ForumPostReasons{}
	if err := c.SearchRead(ForumPostReasonModel, criteria, options, fprs); err != nil {
		return nil, err
	}
	return fprs, nil
}

// FindForumPostReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumPostReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ForumPostReasonModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindForumPostReasonId finds record id by querying it with criteria.
func (c *Client) FindForumPostReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumPostReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("forum.post.reason was not found with criteria %v and options %v", criteria, options)
}
