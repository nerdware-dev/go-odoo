package odoo

import (
	"fmt"
)

// KnowsystemTourArticle represents knowsystem.tour.article model.
type KnowsystemTourArticle struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	ArticleId     *Many2One `xmlrpc:"article_id,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	DoneByUserIds *Relation `xmlrpc:"done_by_user_ids,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	TourId        *Many2One `xmlrpc:"tour_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowsystemTourArticles represents array of knowsystem.tour.article model.
type KnowsystemTourArticles []KnowsystemTourArticle

// KnowsystemTourArticleModel is the odoo model name.
const KnowsystemTourArticleModel = "knowsystem.tour.article"

// Many2One convert KnowsystemTourArticle to *Many2One.
func (kta *KnowsystemTourArticle) Many2One() *Many2One {
	return NewMany2One(kta.Id.Get(), "")
}

// CreateKnowsystemTourArticle creates a new knowsystem.tour.article model and returns its id.
func (c *Client) CreateKnowsystemTourArticle(kta *KnowsystemTourArticle) (int64, error) {
	ids, err := c.CreateKnowsystemTourArticles([]*KnowsystemTourArticle{kta})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemTourArticle creates a new knowsystem.tour.article model and returns its id.
func (c *Client) CreateKnowsystemTourArticles(ktas []*KnowsystemTourArticle) ([]int64, error) {
	var vv []interface{}
	for _, v := range ktas {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemTourArticleModel, vv)
}

// UpdateKnowsystemTourArticle updates an existing knowsystem.tour.article record.
func (c *Client) UpdateKnowsystemTourArticle(kta *KnowsystemTourArticle) error {
	return c.UpdateKnowsystemTourArticles([]int64{kta.Id.Get()}, kta)
}

// UpdateKnowsystemTourArticles updates existing knowsystem.tour.article records.
// All records (represented by ids) will be updated by kta values.
func (c *Client) UpdateKnowsystemTourArticles(ids []int64, kta *KnowsystemTourArticle) error {
	return c.Update(KnowsystemTourArticleModel, ids, kta)
}

// DeleteKnowsystemTourArticle deletes an existing knowsystem.tour.article record.
func (c *Client) DeleteKnowsystemTourArticle(id int64) error {
	return c.DeleteKnowsystemTourArticles([]int64{id})
}

// DeleteKnowsystemTourArticles deletes existing knowsystem.tour.article records.
func (c *Client) DeleteKnowsystemTourArticles(ids []int64) error {
	return c.Delete(KnowsystemTourArticleModel, ids)
}

// GetKnowsystemTourArticle gets knowsystem.tour.article existing record.
func (c *Client) GetKnowsystemTourArticle(id int64) (*KnowsystemTourArticle, error) {
	ktas, err := c.GetKnowsystemTourArticles([]int64{id})
	if err != nil {
		return nil, err
	}
	if ktas != nil && len(*ktas) > 0 {
		return &((*ktas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowsystem.tour.article not found", id)
}

// GetKnowsystemTourArticles gets knowsystem.tour.article existing records.
func (c *Client) GetKnowsystemTourArticles(ids []int64) (*KnowsystemTourArticles, error) {
	ktas := &KnowsystemTourArticles{}
	if err := c.Read(KnowsystemTourArticleModel, ids, nil, ktas); err != nil {
		return nil, err
	}
	return ktas, nil
}

// FindKnowsystemTourArticle finds knowsystem.tour.article record by querying it with criteria.
func (c *Client) FindKnowsystemTourArticle(criteria *Criteria) (*KnowsystemTourArticle, error) {
	ktas := &KnowsystemTourArticles{}
	if err := c.SearchRead(KnowsystemTourArticleModel, criteria, NewOptions().Limit(1), ktas); err != nil {
		return nil, err
	}
	if ktas != nil && len(*ktas) > 0 {
		return &((*ktas)[0]), nil
	}
	return nil, fmt.Errorf("knowsystem.tour.article was not found with criteria %v", criteria)
}

// FindKnowsystemTourArticles finds knowsystem.tour.article records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTourArticles(criteria *Criteria, options *Options) (*KnowsystemTourArticles, error) {
	ktas := &KnowsystemTourArticles{}
	if err := c.SearchRead(KnowsystemTourArticleModel, criteria, options, ktas); err != nil {
		return nil, err
	}
	return ktas, nil
}

// FindKnowsystemTourArticleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemTourArticleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowsystemTourArticleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowsystemTourArticleId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemTourArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemTourArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowsystem.tour.article was not found with criteria %v and options %v", criteria, options)
}
