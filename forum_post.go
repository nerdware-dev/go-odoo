package odoo

// ForumPost represents forum.post model.
type ForumPost struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	BumpDate                 *Time      `xmlrpc:"bump_date,omitempty"`
	CanAccept                *Bool      `xmlrpc:"can_accept,omitempty"`
	CanAnswer                *Bool      `xmlrpc:"can_answer,omitempty"`
	CanAsk                   *Bool      `xmlrpc:"can_ask,omitempty"`
	CanClose                 *Bool      `xmlrpc:"can_close,omitempty"`
	CanComment               *Bool      `xmlrpc:"can_comment,omitempty"`
	CanCommentConvert        *Bool      `xmlrpc:"can_comment_convert,omitempty"`
	CanDisplayBiography      *Bool      `xmlrpc:"can_display_biography,omitempty"`
	CanDownvote              *Bool      `xmlrpc:"can_downvote,omitempty"`
	CanEdit                  *Bool      `xmlrpc:"can_edit,omitempty"`
	CanFlag                  *Bool      `xmlrpc:"can_flag,omitempty"`
	CanModerate              *Bool      `xmlrpc:"can_moderate,omitempty"`
	CanPost                  *Bool      `xmlrpc:"can_post,omitempty"`
	CanUnlink                *Bool      `xmlrpc:"can_unlink,omitempty"`
	CanUpvote                *Bool      `xmlrpc:"can_upvote,omitempty"`
	CanView                  *Bool      `xmlrpc:"can_view,omitempty"`
	ChildCount               *Int       `xmlrpc:"child_count,omitempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omitempty"`
	ClosedDate               *Time      `xmlrpc:"closed_date,omitempty"`
	ClosedReasonId           *Many2One  `xmlrpc:"closed_reason_id,omitempty"`
	ClosedUid                *Many2One  `xmlrpc:"closed_uid,omitempty"`
	Content                  *String    `xmlrpc:"content,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	FavouriteCount           *Int       `xmlrpc:"favourite_count,omitempty"`
	FavouriteIds             *Relation  `xmlrpc:"favourite_ids,omitempty"`
	FlagUserId               *Many2One  `xmlrpc:"flag_user_id,omitempty"`
	ForumId                  *Many2One  `xmlrpc:"forum_id,omitempty"`
	HasValidatedAnswer       *Bool      `xmlrpc:"has_validated_answer,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsCorrect                *Bool      `xmlrpc:"is_correct,omitempty"`
	IsSeoOptimized           *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	KarmaAccept              *Int       `xmlrpc:"karma_accept,omitempty"`
	KarmaClose               *Int       `xmlrpc:"karma_close,omitempty"`
	KarmaComment             *Int       `xmlrpc:"karma_comment,omitempty"`
	KarmaCommentConvert      *Int       `xmlrpc:"karma_comment_convert,omitempty"`
	KarmaEdit                *Int       `xmlrpc:"karma_edit,omitempty"`
	KarmaFlag                *Int       `xmlrpc:"karma_flag,omitempty"`
	KarmaUnlink              *Int       `xmlrpc:"karma_unlink,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	ModeratorId              *Many2One  `xmlrpc:"moderator_id,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omitempty"`
	PlainContent             *String    `xmlrpc:"plain_content,omitempty"`
	Relevancy                *Float     `xmlrpc:"relevancy,omitempty"`
	SelfReply                *Bool      `xmlrpc:"self_reply,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omitempty"`
	UidHasAnswered           *Bool      `xmlrpc:"uid_has_answered,omitempty"`
	UserFavourite            *Bool      `xmlrpc:"user_favourite,omitempty"`
	UserVote                 *Int       `xmlrpc:"user_vote,omitempty"`
	Views                    *Int       `xmlrpc:"views,omitempty"`
	VoteCount                *Int       `xmlrpc:"vote_count,omitempty"`
	VoteIds                  *Relation  `xmlrpc:"vote_ids,omitempty"`
	WebsiteId                *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription   *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords      *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg         *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle         *String    `xmlrpc:"website_meta_title,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ForumPosts represents array of forum.post model.
type ForumPosts []ForumPost

// ForumPostModel is the odoo model name.
const ForumPostModel = "forum.post"

// Many2One convert ForumPost to *Many2One.
func (fp *ForumPost) Many2One() *Many2One {
	return NewMany2One(fp.Id.Get(), "")
}

// CreateForumPost creates a new forum.post model and returns its id.
func (c *Client) CreateForumPost(fp *ForumPost) (int64, error) {
	ids, err := c.CreateForumPosts([]*ForumPost{fp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateForumPost creates a new forum.post model and returns its id.
func (c *Client) CreateForumPosts(fps []*ForumPost) ([]int64, error) {
	var vv []interface{}
	for _, v := range fps {
		vv = append(vv, v)
	}
	return c.Create(ForumPostModel, vv, nil)
}

// UpdateForumPost updates an existing forum.post record.
func (c *Client) UpdateForumPost(fp *ForumPost) error {
	return c.UpdateForumPosts([]int64{fp.Id.Get()}, fp)
}

// UpdateForumPosts updates existing forum.post records.
// All records (represented by ids) will be updated by fp values.
func (c *Client) UpdateForumPosts(ids []int64, fp *ForumPost) error {
	return c.Update(ForumPostModel, ids, fp, nil)
}

// DeleteForumPost deletes an existing forum.post record.
func (c *Client) DeleteForumPost(id int64) error {
	return c.DeleteForumPosts([]int64{id})
}

// DeleteForumPosts deletes existing forum.post records.
func (c *Client) DeleteForumPosts(ids []int64) error {
	return c.Delete(ForumPostModel, ids)
}

// GetForumPost gets forum.post existing record.
func (c *Client) GetForumPost(id int64) (*ForumPost, error) {
	fps, err := c.GetForumPosts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fps)[0]), nil
}

// GetForumPosts gets forum.post existing records.
func (c *Client) GetForumPosts(ids []int64) (*ForumPosts, error) {
	fps := &ForumPosts{}
	if err := c.Read(ForumPostModel, ids, nil, fps); err != nil {
		return nil, err
	}
	return fps, nil
}

// FindForumPost finds forum.post record by querying it with criteria.
func (c *Client) FindForumPost(criteria *Criteria) (*ForumPost, error) {
	fps := &ForumPosts{}
	if err := c.SearchRead(ForumPostModel, criteria, NewOptions().Limit(1), fps); err != nil {
		return nil, err
	}
	return &((*fps)[0]), nil
}

// FindForumPosts finds forum.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumPosts(criteria *Criteria, options *Options) (*ForumPosts, error) {
	fps := &ForumPosts{}
	if err := c.SearchRead(ForumPostModel, criteria, options, fps); err != nil {
		return nil, err
	}
	return fps, nil
}

// FindForumPostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumPostIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ForumPostModel, criteria, options)
}

// FindForumPostId finds record id by querying it with criteria.
func (c *Client) FindForumPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
