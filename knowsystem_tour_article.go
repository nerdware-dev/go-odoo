package odoo

// KnowsystemTourArticle represents knowsystem.tour.article model.
type KnowsystemTourArticle struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	ArticleId     *Many2One `xmlrpc:"article_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	DoneByUserIds *Relation `xmlrpc:"done_by_user_ids,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Sequence      *Int      `xmlrpc:"sequence,omitempty"`
	TourId        *Many2One `xmlrpc:"tour_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(KnowsystemTourArticleModel, vv, nil)
}

// UpdateKnowsystemTourArticle updates an existing knowsystem.tour.article record.
func (c *Client) UpdateKnowsystemTourArticle(kta *KnowsystemTourArticle) error {
	return c.UpdateKnowsystemTourArticles([]int64{kta.Id.Get()}, kta)
}

// UpdateKnowsystemTourArticles updates existing knowsystem.tour.article records.
// All records (represented by ids) will be updated by kta values.
func (c *Client) UpdateKnowsystemTourArticles(ids []int64, kta *KnowsystemTourArticle) error {
	return c.Update(KnowsystemTourArticleModel, ids, kta, nil)
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
	return &((*ktas)[0]), nil
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
	return &((*ktas)[0]), nil
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
	return c.Search(KnowsystemTourArticleModel, criteria, options)
}

// FindKnowsystemTourArticleId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemTourArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemTourArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
