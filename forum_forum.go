package odoo

// ForumForum represents forum.forum model.
type ForumForum struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	AllowBump                   *Bool      `xmlrpc:"allow_bump,omitempty"`
	AllowShare                  *Bool      `xmlrpc:"allow_share,omitempty"`
	CountFlaggedPosts           *Int       `xmlrpc:"count_flagged_posts,omitempty"`
	CountPostsWaitingValidation *Int       `xmlrpc:"count_posts_waiting_validation,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultOrder                *Selection `xmlrpc:"default_order,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Faq                         *String    `xmlrpc:"faq,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	Image1024                   *String    `xmlrpc:"image_1024,omitempty"`
	Image128                    *String    `xmlrpc:"image_128,omitempty"`
	Image1920                   *String    `xmlrpc:"image_1920,omitempty"`
	Image256                    *String    `xmlrpc:"image_256,omitempty"`
	Image512                    *String    `xmlrpc:"image_512,omitempty"`
	IsSeoOptimized              *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	KarmaAnswer                 *Int       `xmlrpc:"karma_answer,omitempty"`
	KarmaAnswerAcceptAll        *Int       `xmlrpc:"karma_answer_accept_all,omitempty"`
	KarmaAnswerAcceptOwn        *Int       `xmlrpc:"karma_answer_accept_own,omitempty"`
	KarmaAsk                    *Int       `xmlrpc:"karma_ask,omitempty"`
	KarmaCloseAll               *Int       `xmlrpc:"karma_close_all,omitempty"`
	KarmaCloseOwn               *Int       `xmlrpc:"karma_close_own,omitempty"`
	KarmaCommentAll             *Int       `xmlrpc:"karma_comment_all,omitempty"`
	KarmaCommentConvertAll      *Int       `xmlrpc:"karma_comment_convert_all,omitempty"`
	KarmaCommentConvertOwn      *Int       `xmlrpc:"karma_comment_convert_own,omitempty"`
	KarmaCommentOwn             *Int       `xmlrpc:"karma_comment_own,omitempty"`
	KarmaCommentUnlinkAll       *Int       `xmlrpc:"karma_comment_unlink_all,omitempty"`
	KarmaCommentUnlinkOwn       *Int       `xmlrpc:"karma_comment_unlink_own,omitempty"`
	KarmaDofollow               *Int       `xmlrpc:"karma_dofollow,omitempty"`
	KarmaDownvote               *Int       `xmlrpc:"karma_downvote,omitempty"`
	KarmaEditAll                *Int       `xmlrpc:"karma_edit_all,omitempty"`
	KarmaEditOwn                *Int       `xmlrpc:"karma_edit_own,omitempty"`
	KarmaEditRetag              *Int       `xmlrpc:"karma_edit_retag,omitempty"`
	KarmaEditor                 *Int       `xmlrpc:"karma_editor,omitempty"`
	KarmaFlag                   *Int       `xmlrpc:"karma_flag,omitempty"`
	KarmaGenAnswerAccept        *Int       `xmlrpc:"karma_gen_answer_accept,omitempty"`
	KarmaGenAnswerAccepted      *Int       `xmlrpc:"karma_gen_answer_accepted,omitempty"`
	KarmaGenAnswerDownvote      *Int       `xmlrpc:"karma_gen_answer_downvote,omitempty"`
	KarmaGenAnswerFlagged       *Int       `xmlrpc:"karma_gen_answer_flagged,omitempty"`
	KarmaGenAnswerUpvote        *Int       `xmlrpc:"karma_gen_answer_upvote,omitempty"`
	KarmaGenQuestionDownvote    *Int       `xmlrpc:"karma_gen_question_downvote,omitempty"`
	KarmaGenQuestionNew         *Int       `xmlrpc:"karma_gen_question_new,omitempty"`
	KarmaGenQuestionUpvote      *Int       `xmlrpc:"karma_gen_question_upvote,omitempty"`
	KarmaModerate               *Int       `xmlrpc:"karma_moderate,omitempty"`
	KarmaPost                   *Int       `xmlrpc:"karma_post,omitempty"`
	KarmaTagCreate              *Int       `xmlrpc:"karma_tag_create,omitempty"`
	KarmaUnlinkAll              *Int       `xmlrpc:"karma_unlink_all,omitempty"`
	KarmaUnlinkOwn              *Int       `xmlrpc:"karma_unlink_own,omitempty"`
	KarmaUpvote                 *Int       `xmlrpc:"karma_upvote,omitempty"`
	KarmaUserBio                *Int       `xmlrpc:"karma_user_bio,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Mode                        *Selection `xmlrpc:"mode,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	PostIds                     *Relation  `xmlrpc:"post_ids,omitempty"`
	RelevancyPostVote           *Float     `xmlrpc:"relevancy_post_vote,omitempty"`
	RelevancyTimeDecay          *Float     `xmlrpc:"relevancy_time_decay,omitempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omitempty"`
	TotalAnswers                *Int       `xmlrpc:"total_answers,omitempty"`
	TotalFavorites              *Int       `xmlrpc:"total_favorites,omitempty"`
	TotalPosts                  *Int       `xmlrpc:"total_posts,omitempty"`
	TotalViews                  *Int       `xmlrpc:"total_views,omitempty"`
	WebsiteId                   *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription      *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords         *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg            *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle            *String    `xmlrpc:"website_meta_title,omitempty"`
	WelcomeMessage              *String    `xmlrpc:"welcome_message,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(ForumForumModel, vv, nil)
}

// UpdateForumForum updates an existing forum.forum record.
func (c *Client) UpdateForumForum(ff *ForumForum) error {
	return c.UpdateForumForums([]int64{ff.Id.Get()}, ff)
}

// UpdateForumForums updates existing forum.forum records.
// All records (represented by ids) will be updated by ff values.
func (c *Client) UpdateForumForums(ids []int64, ff *ForumForum) error {
	return c.Update(ForumForumModel, ids, ff, nil)
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
	return &((*ffs)[0]), nil
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
	return &((*ffs)[0]), nil
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
	return c.Search(ForumForumModel, criteria, options)
}

// FindForumForumId finds record id by querying it with criteria.
func (c *Client) FindForumForumId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ForumForumModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
