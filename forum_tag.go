package odoo

import (
	"fmt"
)

// ForumTag represents forum.tag model.
type ForumTag struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	ForumId                  *Many2One `xmlrpc:"forum_id,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsSeoOptimized           *Bool     `xmlrpc:"is_seo_optimized,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	PostIds                  *Relation `xmlrpc:"post_ids,omptempty"`
	PostsCount               *Int      `xmlrpc:"posts_count,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription   *String   `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords      *String   `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg         *String   `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle         *String   `xmlrpc:"website_meta_title,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ForumTags represents array of forum.tag model.
type ForumTags []ForumTag

// ForumTagModel is the odoo model name.
const ForumTagModel = "forum.tag"

// Many2One convert ForumTag to *Many2One.
func (ft *ForumTag) Many2One() *Many2One {
	return NewMany2One(ft.Id.Get(), "")
}

// CreateForumTag creates a new forum.tag model and returns its id.
func (c *Client) CreateForumTag(ft *ForumTag) (int64, error) {
	ids, err := c.CreateForumTags([]*ForumTag{ft})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateForumTag creates a new forum.tag model and returns its id.
func (c *Client) CreateForumTags(fts []*ForumTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range fts {
		vv = append(vv, v)
	}
	return c.Create(ForumTagModel, vv)
}

// UpdateForumTag updates an existing forum.tag record.
func (c *Client) UpdateForumTag(ft *ForumTag) error {
	return c.UpdateForumTags([]int64{ft.Id.Get()}, ft)
}

// UpdateForumTags updates existing forum.tag records.
// All records (represented by ids) will be updated by ft values.
func (c *Client) UpdateForumTags(ids []int64, ft *ForumTag) error {
	return c.Update(ForumTagModel, ids, ft)
}

// DeleteForumTag deletes an existing forum.tag record.
func (c *Client) DeleteForumTag(id int64) error {
	return c.DeleteForumTags([]int64{id})
}

// DeleteForumTags deletes existing forum.tag records.
func (c *Client) DeleteForumTags(ids []int64) error {
	return c.Delete(ForumTagModel, ids)
}

// GetForumTag gets forum.tag existing record.
func (c *Client) GetForumTag(id int64) (*ForumTag, error) {
	fts, err := c.GetForumTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if fts != nil && len(*fts) > 0 {
		return &((*fts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of forum.tag not found", id)
}

// GetForumTags gets forum.tag existing records.
func (c *Client) GetForumTags(ids []int64) (*ForumTags, error) {
	fts := &ForumTags{}
	if err := c.Read(ForumTagModel, ids, nil, fts); err != nil {
		return nil, err
	}
	return fts, nil
}

// FindForumTag finds forum.tag record by querying it with criteria.
func (c *Client) FindForumTag(criteria *Criteria) (*ForumTag, error) {
	fts := &ForumTags{}
	if err := c.SearchRead(ForumTagModel, criteria, NewOptions().Limit(1), fts); err != nil {
		return nil, err
	}
	if fts != nil && len(*fts) > 0 {
		return &((*fts)[0]), nil
	}
	return nil, fmt.Errorf("forum.tag was not found with criteria %v", criteria)
}

// FindForumTags finds forum.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumTags(criteria *Criteria, options *Options) (*ForumTags, error) {
	fts := &ForumTags{}
	if err := c.SearchRead(ForumTagModel, criteria, options, fts); err != nil {
		return nil, err
	}
	return fts, nil
}

// FindForumTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ForumTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindForumTagId finds record id by querying it with criteria.
func (c *Client) FindForumTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("forum.tag was not found with criteria %v and options %v", criteria, options)
}
