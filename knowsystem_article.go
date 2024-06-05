package odoo

import (
	"fmt"
)

// KnowsystemArticle represents knowsystem.article model.
type KnowsystemArticle struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omptempty"`
	AccessUserIds               *Relation  `xmlrpc:"access_user_ids,omptempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omptempty"`
	CanPublish                  *Bool      `xmlrpc:"can_publish,omptempty"`
	ContributorIds              *Relation  `xmlrpc:"contributor_ids,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DescriptionArch             *String    `xmlrpc:"description_arch,omptempty"`
	DislikeUserIds              *Relation  `xmlrpc:"dislike_user_ids,omptempty"`
	DislikesNumber              *Int       `xmlrpc:"dislikes_number,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	FavouriteUserIds            *Relation  `xmlrpc:"favourite_user_ids,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IndexedDescription          *String    `xmlrpc:"indexed_description,omptempty"`
	InternalUrl                 *String    `xmlrpc:"internal_url,omptempty"`
	IsPublished                 *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized              *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	KanbanDescription           *String    `xmlrpc:"kanban_description,omptempty"`
	KanbanManualDescription     *String    `xmlrpc:"kanban_manual_description,omptempty"`
	LikeUserIds                 *Relation  `xmlrpc:"like_user_ids,omptempty"`
	LikesNumber                 *Int       `xmlrpc:"likes_number,omptempty"`
	LikesScore                  *Int       `xmlrpc:"likes_score,omptempty"`
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
	Name                        *String    `xmlrpc:"name,omptempty"`
	RevisionIds                 *Relation  `xmlrpc:"revision_ids,omptempty"`
	SectionId                   *Many2One  `xmlrpc:"section_id,omptempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omptempty"`
	ThisUserLikeState           *Selection `xmlrpc:"this_user_like_state,omptempty"`
	UsedInEmailCompose          *Int       `xmlrpc:"used_in_email_compose,omptempty"`
	UserGroupIds                *Relation  `xmlrpc:"user_group_ids,omptempty"`
	ViewStatIds                 *Relation  `xmlrpc:"view_stat_ids,omptempty"`
	ViewsNumberInternal         *Int       `xmlrpc:"views_number_internal,omptempty"`
	WebsiteId                   *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription      *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords         *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg            *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle            *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished            *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                  *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteRevisionDate           *Time      `xmlrpc:"write_revision_date,omptempty"`
	WriteRevisionUid            *Many2One  `xmlrpc:"write_revision_uid,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(KnowsystemArticleModel, vv)
}

// UpdateKnowsystemArticle updates an existing knowsystem.article record.
func (c *Client) UpdateKnowsystemArticle(ka *KnowsystemArticle) error {
	return c.UpdateKnowsystemArticles([]int64{ka.Id.Get()}, ka)
}

// UpdateKnowsystemArticles updates existing knowsystem.article records.
// All records (represented by ids) will be updated by ka values.
func (c *Client) UpdateKnowsystemArticles(ids []int64, ka *KnowsystemArticle) error {
	return c.Update(KnowsystemArticleModel, ids, ka)
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
	if kas != nil && len(*kas) > 0 {
		return &((*kas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.article not found", id)
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
	if kas != nil && len(*kas) > 0 {
		return &((*kas)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.article was not found with criteria %v", criteria)
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
	ids, err := c.Search(KnowsystemArticleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemArticleId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.article was not found with criteria %v and options %v", criteria, options)
}
