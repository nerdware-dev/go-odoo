package odoo

import (
	"fmt"
)

// ForumPost represents forum.post model.
type ForumPost struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	BumpDate                 *Time      `xmlrpc:"bump_date,omptempty"`
	CanAccept                *Bool      `xmlrpc:"can_accept,omptempty"`
	CanAnswer                *Bool      `xmlrpc:"can_answer,omptempty"`
	CanAsk                   *Bool      `xmlrpc:"can_ask,omptempty"`
	CanClose                 *Bool      `xmlrpc:"can_close,omptempty"`
	CanComment               *Bool      `xmlrpc:"can_comment,omptempty"`
	CanCommentConvert        *Bool      `xmlrpc:"can_comment_convert,omptempty"`
	CanDisplayBiography      *Bool      `xmlrpc:"can_display_biography,omptempty"`
	CanDownvote              *Bool      `xmlrpc:"can_downvote,omptempty"`
	CanEdit                  *Bool      `xmlrpc:"can_edit,omptempty"`
	CanFlag                  *Bool      `xmlrpc:"can_flag,omptempty"`
	CanModerate              *Bool      `xmlrpc:"can_moderate,omptempty"`
	CanPost                  *Bool      `xmlrpc:"can_post,omptempty"`
	CanUnlink                *Bool      `xmlrpc:"can_unlink,omptempty"`
	CanUpvote                *Bool      `xmlrpc:"can_upvote,omptempty"`
	CanView                  *Bool      `xmlrpc:"can_view,omptempty"`
	ChildCount               *Int       `xmlrpc:"child_count,omptempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omptempty"`
	ClosedDate               *Time      `xmlrpc:"closed_date,omptempty"`
	ClosedReasonId           *Many2One  `xmlrpc:"closed_reason_id,omptempty"`
	ClosedUid                *Many2One  `xmlrpc:"closed_uid,omptempty"`
	Content                  *String    `xmlrpc:"content,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FavouriteCount           *Int       `xmlrpc:"favourite_count,omptempty"`
	FavouriteIds             *Relation  `xmlrpc:"favourite_ids,omptempty"`
	FlagUserId               *Many2One  `xmlrpc:"flag_user_id,omptempty"`
	ForumId                  *Many2One  `xmlrpc:"forum_id,omptempty"`
	HasValidatedAnswer       *Bool      `xmlrpc:"has_validated_answer,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IsCorrect                *Bool      `xmlrpc:"is_correct,omptempty"`
	IsSeoOptimized           *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	KarmaAccept              *Int       `xmlrpc:"karma_accept,omptempty"`
	KarmaClose               *Int       `xmlrpc:"karma_close,omptempty"`
	KarmaComment             *Int       `xmlrpc:"karma_comment,omptempty"`
	KarmaCommentConvert      *Int       `xmlrpc:"karma_comment_convert,omptempty"`
	KarmaEdit                *Int       `xmlrpc:"karma_edit,omptempty"`
	KarmaFlag                *Int       `xmlrpc:"karma_flag,omptempty"`
	KarmaUnlink              *Int       `xmlrpc:"karma_unlink,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	ModeratorId              *Many2One  `xmlrpc:"moderator_id,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omptempty"`
	PlainContent             *String    `xmlrpc:"plain_content,omptempty"`
	Relevancy                *Float     `xmlrpc:"relevancy,omptempty"`
	SelfReply                *Bool      `xmlrpc:"self_reply,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omptempty"`
	UidHasAnswered           *Bool      `xmlrpc:"uid_has_answered,omptempty"`
	UserFavourite            *Bool      `xmlrpc:"user_favourite,omptempty"`
	UserVote                 *Int       `xmlrpc:"user_vote,omptempty"`
	Views                    *Int       `xmlrpc:"views,omptempty"`
	VoteCount                *Int       `xmlrpc:"vote_count,omptempty"`
	VoteIds                  *Relation  `xmlrpc:"vote_ids,omptempty"`
	WebsiteId                *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription   *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords      *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg         *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle         *String    `xmlrpc:"website_meta_title,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(ForumPostModel, vv)
}

// UpdateForumPost updates an existing forum.post record.
func (c *Client) UpdateForumPost(fp *ForumPost) error {
	return c.UpdateForumPosts([]int64{fp.Id.Get()}, fp)
}

// UpdateForumPosts updates existing forum.post records.
// All records (represented by ids) will be updated by fp values.
func (c *Client) UpdateForumPosts(ids []int64, fp *ForumPost) error {
	return c.Update(ForumPostModel, ids, fp)
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
	if fps != nil && len(*fps) > 0 {
		return &((*fps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of forum.post not found", id)
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
	if fps != nil && len(*fps) > 0 {
		return &((*fps)[0]), nil
	}
	return nil, fmt.Errorf("forum.post was not found with criteria %v", criteria)
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
	ids, err := c.Search(ForumPostModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindForumPostId finds record id by querying it with criteria.
func (c *Client) FindForumPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("forum.post was not found with criteria %v and options %v", criteria, options)
}
