package odoo

// KnowsystemArticle represents knowsystem.article model.
type KnowsystemArticle struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omitempty"`
	AccessUserIds               *Relation  `xmlrpc:"access_user_ids,omitempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omitempty"`
	CanPublish                  *Bool      `xmlrpc:"can_publish,omitempty"`
	ContributorIds              *Relation  `xmlrpc:"contributor_ids,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DescriptionArch             *String    `xmlrpc:"description_arch,omitempty"`
	DislikeUserIds              *Relation  `xmlrpc:"dislike_user_ids,omitempty"`
	DislikesNumber              *Int       `xmlrpc:"dislikes_number,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	FavouriteUserIds            *Relation  `xmlrpc:"favourite_user_ids,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IndexedDescription          *String    `xmlrpc:"indexed_description,omitempty"`
	InternalUrl                 *String    `xmlrpc:"internal_url,omitempty"`
	IsPublished                 *Bool      `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized              *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	KanbanDescription           *String    `xmlrpc:"kanban_description,omitempty"`
	KanbanManualDescription     *String    `xmlrpc:"kanban_manual_description,omitempty"`
	LikeUserIds                 *Relation  `xmlrpc:"like_user_ids,omitempty"`
	LikesNumber                 *Int       `xmlrpc:"likes_number,omitempty"`
	LikesScore                  *Int       `xmlrpc:"likes_score,omitempty"`
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
	Name                        *String    `xmlrpc:"name,omitempty"`
	RevisionIds                 *Relation  `xmlrpc:"revision_ids,omitempty"`
	SectionId                   *Many2One  `xmlrpc:"section_id,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	ThisUserLikeState           *Selection `xmlrpc:"this_user_like_state,omitempty"`
	UsedInEmailCompose          *Int       `xmlrpc:"used_in_email_compose,omitempty"`
	UserGroupIds                *Relation  `xmlrpc:"user_group_ids,omitempty"`
	ViewStatIds                 *Relation  `xmlrpc:"view_stat_ids,omitempty"`
	ViewsNumberInternal         *Int       `xmlrpc:"views_number_internal,omitempty"`
	WebsiteId                   *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription      *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords         *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg            *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle            *String    `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished            *Bool      `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                  *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteRevisionDate           *Time      `xmlrpc:"write_revision_date,omitempty"`
	WriteRevisionUid            *Many2One  `xmlrpc:"write_revision_uid,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// KnowsystemArticles represents array of knowsystem.article model.
type KnowsystemArticles []KnowsystemArticle

// KnowsystemArticleModel is the odoo model name.
const KnowsystemArticleModel = "knowsystem.article"

// Many2One convert KnowsystemArticle to *Many2One.
func (ka *KnowsystemArticle) Many2One() *Many2One {
	return NewMany2One(ka.Id.Get(), "")
}

// CreateKnowsystemArticle creates a new knowsystem.article model and returns its id.
func (c *Client) CreateKnowsystemArticle(ka *KnowsystemArticle) (int64, error) {
	ids, err := c.CreateKnowsystemArticles([]*KnowsystemArticle{ka})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemArticle creates a new knowsystem.article model and returns its id.
func (c *Client) CreateKnowsystemArticles(kas []*KnowsystemArticle) ([]int64, error) {
	var vv []interface{}
	for _, v := range kas {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemArticleModel, vv, nil)
}

// UpdateKnowsystemArticle updates an existing knowsystem.article record.
func (c *Client) UpdateKnowsystemArticle(ka *KnowsystemArticle) error {
	return c.UpdateKnowsystemArticles([]int64{ka.Id.Get()}, ka)
}

// UpdateKnowsystemArticles updates existing knowsystem.article records.
// All records (represented by ids) will be updated by ka values.
func (c *Client) UpdateKnowsystemArticles(ids []int64, ka *KnowsystemArticle) error {
	return c.Update(KnowsystemArticleModel, ids, ka, nil)
}

// DeleteKnowsystemArticle deletes an existing knowsystem.article record.
func (c *Client) DeleteKnowsystemArticle(id int64) error {
	return c.DeleteKnowsystemArticles([]int64{id})
}

// DeleteKnowsystemArticles deletes existing knowsystem.article records.
func (c *Client) DeleteKnowsystemArticles(ids []int64) error {
	return c.Delete(KnowsystemArticleModel, ids)
}

// GetKnowsystemArticle gets knowsystem.article existing record.
func (c *Client) GetKnowsystemArticle(id int64) (*KnowsystemArticle, error) {
	kas, err := c.GetKnowsystemArticles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kas)[0]), nil
}

// GetKnowsystemArticles gets knowsystem.article existing records.
func (c *Client) GetKnowsystemArticles(ids []int64) (*KnowsystemArticles, error) {
	kas := &KnowsystemArticles{}
	if err := c.Read(KnowsystemArticleModel, ids, nil, kas); err != nil {
		return nil, err
	}
	return kas, nil
}

// FindKnowsystemArticle finds knowsystem.article record by querying it with criteria.
func (c *Client) FindKnowsystemArticle(criteria *Criteria) (*KnowsystemArticle, error) {
	kas := &KnowsystemArticles{}
	if err := c.SearchRead(KnowsystemArticleModel, criteria, NewOptions().Limit(1), kas); err != nil {
		return nil, err
	}
	return &((*kas)[0]), nil
}

// FindKnowsystemArticles finds knowsystem.article records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticles(criteria *Criteria, options *Options) (*KnowsystemArticles, error) {
	kas := &KnowsystemArticles{}
	if err := c.SearchRead(KnowsystemArticleModel, criteria, options, kas); err != nil {
		return nil, err
	}
	return kas, nil
}

// FindKnowsystemArticleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowsystemArticleModel, criteria, options)
}

// FindKnowsystemArticleId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
