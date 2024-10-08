package odoo

// ArticleSearch represents article.search model.
type ArticleSearch struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	ArticleIds         *Relation `xmlrpc:"article_ids,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	NoSelection        *Bool     `xmlrpc:"no_selection,omitempty"`
	Search             *String   `xmlrpc:"search,omitempty"`
	SectionIds         *Relation `xmlrpc:"section_ids,omitempty"`
	SelectedArticleIds *Relation `xmlrpc:"selected_article_ids,omitempty"`
	TagIds             *Relation `xmlrpc:"tag_ids,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ArticleSearchs represents array of article.search model.
type ArticleSearchs []ArticleSearch

// ArticleSearchModel is the odoo model name.
const ArticleSearchModel = "article.search"

// Many2One convert ArticleSearch to *Many2One.
func (as *ArticleSearch) Many2One() *Many2One {
	return NewMany2One(as.Id.Get(), "")
}

// CreateArticleSearch creates a new article.search model and returns its id.
func (c *Client) CreateArticleSearch(as *ArticleSearch) (int64, error) {
	ids, err := c.CreateArticleSearchs([]*ArticleSearch{as})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateArticleSearch creates a new article.search model and returns its id.
func (c *Client) CreateArticleSearchs(ass []*ArticleSearch) ([]int64, error) {
	var vv []interface{}
	for _, v := range ass {
		vv = append(vv, v)
	}
	return c.Create(ArticleSearchModel, vv, nil)
}

// UpdateArticleSearch updates an existing article.search record.
func (c *Client) UpdateArticleSearch(as *ArticleSearch) error {
	return c.UpdateArticleSearchs([]int64{as.Id.Get()}, as)
}

// UpdateArticleSearchs updates existing article.search records.
// All records (represented by ids) will be updated by as values.
func (c *Client) UpdateArticleSearchs(ids []int64, as *ArticleSearch) error {
	return c.Update(ArticleSearchModel, ids, as, nil)
}

// DeleteArticleSearch deletes an existing article.search record.
func (c *Client) DeleteArticleSearch(id int64) error {
	return c.DeleteArticleSearchs([]int64{id})
}

// DeleteArticleSearchs deletes existing article.search records.
func (c *Client) DeleteArticleSearchs(ids []int64) error {
	return c.Delete(ArticleSearchModel, ids)
}

// GetArticleSearch gets article.search existing record.
func (c *Client) GetArticleSearch(id int64) (*ArticleSearch, error) {
	ass, err := c.GetArticleSearchs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ass)[0]), nil
}

// GetArticleSearchs gets article.search existing records.
func (c *Client) GetArticleSearchs(ids []int64) (*ArticleSearchs, error) {
	ass := &ArticleSearchs{}
	if err := c.Read(ArticleSearchModel, ids, nil, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindArticleSearch finds article.search record by querying it with criteria.
func (c *Client) FindArticleSearch(criteria *Criteria) (*ArticleSearch, error) {
	ass := &ArticleSearchs{}
	if err := c.SearchRead(ArticleSearchModel, criteria, NewOptions().Limit(1), ass); err != nil {
		return nil, err
	}
	return &((*ass)[0]), nil
}

// FindArticleSearchs finds article.search records by querying it
// and filtering it with criteria and options.
func (c *Client) FindArticleSearchs(criteria *Criteria, options *Options) (*ArticleSearchs, error) {
	ass := &ArticleSearchs{}
	if err := c.SearchRead(ArticleSearchModel, criteria, options, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindArticleSearchIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindArticleSearchIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ArticleSearchModel, criteria, options)
}

// FindArticleSearchId finds record id by querying it with criteria.
func (c *Client) FindArticleSearchId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ArticleSearchModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
