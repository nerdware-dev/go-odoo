package odoo

import (
	"fmt"
)

// ForumForum represents forum.forum model.
type ForumForum struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	AllowBump                   *Bool      `xmlrpc:"allow_bump,omptempty"`
	AllowShare                  *Bool      `xmlrpc:"allow_share,omptempty"`
	CountFlaggedPosts           *Int       `xmlrpc:"count_flagged_posts,omptempty"`
	CountPostsWaitingValidation *Int       `xmlrpc:"count_posts_waiting_validation,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultOrder                *Selection `xmlrpc:"default_order,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Faq                         *String    `xmlrpc:"faq,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	Image1024                   *String    `xmlrpc:"image_1024,omptempty"`
	Image128                    *String    `xmlrpc:"image_128,omptempty"`
	Image1920                   *String    `xmlrpc:"image_1920,omptempty"`
	Image256                    *String    `xmlrpc:"image_256,omptempty"`
	Image512                    *String    `xmlrpc:"image_512,omptempty"`
	IsSeoOptimized              *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	KarmaAnswer                 *Int       `xmlrpc:"karma_answer,omptempty"`
	KarmaAnswerAcceptAll        *Int       `xmlrpc:"karma_answer_accept_all,omptempty"`
	KarmaAnswerAcceptOwn        *Int       `xmlrpc:"karma_answer_accept_own,omptempty"`
	KarmaAsk                    *Int       `xmlrpc:"karma_ask,omptempty"`
	KarmaCloseAll               *Int       `xmlrpc:"karma_close_all,omptempty"`
	KarmaCloseOwn               *Int       `xmlrpc:"karma_close_own,omptempty"`
	KarmaCommentAll             *Int       `xmlrpc:"karma_comment_all,omptempty"`
	KarmaCommentConvertAll      *Int       `xmlrpc:"karma_comment_convert_all,omptempty"`
	KarmaCommentConvertOwn      *Int       `xmlrpc:"karma_comment_convert_own,omptempty"`
	KarmaCommentOwn             *Int       `xmlrpc:"karma_comment_own,omptempty"`
	KarmaCommentUnlinkAll       *Int       `xmlrpc:"karma_comment_unlink_all,omptempty"`
	KarmaCommentUnlinkOwn       *Int       `xmlrpc:"karma_comment_unlink_own,omptempty"`
	KarmaDofollow               *Int       `xmlrpc:"karma_dofollow,omptempty"`
	KarmaDownvote               *Int       `xmlrpc:"karma_downvote,omptempty"`
	KarmaEditAll                *Int       `xmlrpc:"karma_edit_all,omptempty"`
	KarmaEditOwn                *Int       `xmlrpc:"karma_edit_own,omptempty"`
	KarmaEditRetag              *Int       `xmlrpc:"karma_edit_retag,omptempty"`
	KarmaEditor                 *Int       `xmlrpc:"karma_editor,omptempty"`
	KarmaFlag                   *Int       `xmlrpc:"karma_flag,omptempty"`
	KarmaGenAnswerAccept        *Int       `xmlrpc:"karma_gen_answer_accept,omptempty"`
	KarmaGenAnswerAccepted      *Int       `xmlrpc:"karma_gen_answer_accepted,omptempty"`
	KarmaGenAnswerDownvote      *Int       `xmlrpc:"karma_gen_answer_downvote,omptempty"`
	KarmaGenAnswerFlagged       *Int       `xmlrpc:"karma_gen_answer_flagged,omptempty"`
	KarmaGenAnswerUpvote        *Int       `xmlrpc:"karma_gen_answer_upvote,omptempty"`
	KarmaGenQuestionDownvote    *Int       `xmlrpc:"karma_gen_question_downvote,omptempty"`
	KarmaGenQuestionNew         *Int       `xmlrpc:"karma_gen_question_new,omptempty"`
	KarmaGenQuestionUpvote      *Int       `xmlrpc:"karma_gen_question_upvote,omptempty"`
	KarmaModerate               *Int       `xmlrpc:"karma_moderate,omptempty"`
	KarmaPost                   *Int       `xmlrpc:"karma_post,omptempty"`
	KarmaTagCreate              *Int       `xmlrpc:"karma_tag_create,omptempty"`
	KarmaUnlinkAll              *Int       `xmlrpc:"karma_unlink_all,omptempty"`
	KarmaUnlinkOwn              *Int       `xmlrpc:"karma_unlink_own,omptempty"`
	KarmaUpvote                 *Int       `xmlrpc:"karma_upvote,omptempty"`
	KarmaUserBio                *Int       `xmlrpc:"karma_user_bio,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mode                        *Selection `xmlrpc:"mode,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PostIds                     *Relation  `xmlrpc:"post_ids,omptempty"`
	RelevancyPostVote           *Float     `xmlrpc:"relevancy_post_vote,omptempty"`
	RelevancyTimeDecay          *Float     `xmlrpc:"relevancy_time_decay,omptempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omptempty"`
	TotalAnswers                *Int       `xmlrpc:"total_answers,omptempty"`
	TotalFavorites              *Int       `xmlrpc:"total_favorites,omptempty"`
	TotalPosts                  *Int       `xmlrpc:"total_posts,omptempty"`
	TotalViews                  *Int       `xmlrpc:"total_views,omptempty"`
	WebsiteId                   *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription      *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords         *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg            *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle            *String    `xmlrpc:"website_meta_title,omptempty"`
	WelcomeMessage              *String    `xmlrpc:"welcome_message,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ForumForums represents array of forum.forum model.
type ForumForums []ForumForum

// ForumForumModel is the odoo model name.
const ForumForumModel = "forum.forum"

// Many2One convert ForumForum to *Many2One.
func (ff *ForumForum) Many2One() *Many2One {
	return NewMany2One(ff.Id.Get(), "")
}

// CreateForumForum creates a new forum.forum model and returns its id.
func (c *Client) CreateForumForum(ff *ForumForum) (int64, error) {
	ids, err := c.CreateForumForums([]*ForumForum{ff})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateForumForum creates a new forum.forum model and returns its id.
func (c *Client) CreateForumForums(ffs []*ForumForum) ([]int64, error) {
	var vv []interface{}
	for _, v := range ffs {
		vv = append(vv, v)
	}
	return c.Create(ForumForumModel, vv)
}

// UpdateForumForum updates an existing forum.forum record.
func (c *Client) UpdateForumForum(ff *ForumForum) error {
	return c.UpdateForumForums([]int64{ff.Id.Get()}, ff)
}

// UpdateForumForums updates existing forum.forum records.
// All records (represented by ids) will be updated by ff values.
func (c *Client) UpdateForumForums(ids []int64, ff *ForumForum) error {
	return c.Update(ForumForumModel, ids, ff)
}

// DeleteForumForum deletes an existing forum.forum record.
func (c *Client) DeleteForumForum(id int64) error {
	return c.DeleteForumForums([]int64{id})
}

// DeleteForumForums deletes existing forum.forum records.
func (c *Client) DeleteForumForums(ids []int64) error {
	return c.Delete(ForumForumModel, ids)
}

// GetForumForum gets forum.forum existing record.
func (c *Client) GetForumForum(id int64) (*ForumForum, error) {
	ffs, err := c.GetForumForums([]int64{id})
	if err != nil {
		return nil, err
	}
	if ffs != nil && len(*ffs) > 0 {
		return &((*ffs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of forum.forum not found", id)
}

// GetForumForums gets forum.forum existing records.
func (c *Client) GetForumForums(ids []int64) (*ForumForums, error) {
	ffs := &ForumForums{}
	if err := c.Read(ForumForumModel, ids, nil, ffs); err != nil {
		return nil, err
	}
	return ffs, nil
}

// FindForumForum finds forum.forum record by querying it with criteria.
func (c *Client) FindForumForum(criteria *Criteria) (*ForumForum, error) {
	ffs := &ForumForums{}
	if err := c.SearchRead(ForumForumModel, criteria, NewOptions().Limit(1), ffs); err != nil {
		return nil, err
	}
	if ffs != nil && len(*ffs) > 0 {
		return &((*ffs)[0]), nil
	}
	return nil, fmt.Errorf("forum.forum was not found with criteria %v", criteria)
}

// FindForumForums finds forum.forum records by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumForums(criteria *Criteria, options *Options) (*ForumForums, error) {
	ffs := &ForumForums{}
	if err := c.SearchRead(ForumForumModel, criteria, options, ffs); err != nil {
		return nil, err
	}
	return ffs, nil
}

// FindForumForumIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindForumForumIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ForumForumModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindForumForumId finds record id by querying it with criteria.
func (c *Client) FindForumForumId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumForumModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("forum.forum was not found with criteria %v and options %v", criteria, options)
}
