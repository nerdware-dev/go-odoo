package odoo

// ForumTag represents forum.tag model.
type ForumTag struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	ForumId                  *Many2One `xmlrpc:"forum_id,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	IsSeoOptimized           *Bool     `xmlrpc:"is_seo_optimized,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	PostIds                  *Relation `xmlrpc:"post_ids,omitempty"`
	PostsCount               *Int      `xmlrpc:"posts_count,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription   *String   `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords      *String   `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg         *String   `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle         *String   `xmlrpc:"website_meta_title,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(ForumTagModel, vv, nil)
}

// UpdateForumTag updates an existing forum.tag record.
func (c *Client) UpdateForumTag(ft *ForumTag) error {
	return c.UpdateForumTags([]int64{ft.Id.Get()}, ft)
}

// UpdateForumTags updates existing forum.tag records.
// All records (represented by ids) will be updated by ft values.
func (c *Client) UpdateForumTags(ids []int64, ft *ForumTag) error {
	return c.Update(ForumTagModel, ids, ft, nil)
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
	return &((*fts)[0]), nil
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
	return &((*fts)[0]), nil
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
	return c.Search(ForumTagModel, criteria, options)
}

// FindForumTagId finds record id by querying it with criteria.
func (c *Client) FindForumTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
