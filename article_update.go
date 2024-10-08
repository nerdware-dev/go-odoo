package odoo

// ArticleUpdate represents article.update model.
type ArticleUpdate struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	Activate       *Selection `xmlrpc:"activate,omitempty"`
	Articles       *String    `xmlrpc:"articles,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	SectionId      *Many2One  `xmlrpc:"section_id,omitempty"`
	ToAddTagIds    *Relation  `xmlrpc:"to_add_tag_ids,omitempty"`
	ToRemoveTagIds *Relation  `xmlrpc:"to_remove_tag_ids,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ArticleUpdates represents array of article.update model.
type ArticleUpdates []ArticleUpdate

// ArticleUpdateModel is the odoo model name.
const ArticleUpdateModel = "article.update"

// Many2One convert ArticleUpdate to *Many2One.
func (au *ArticleUpdate) Many2One() *Many2One {
	return NewMany2One(au.Id.Get(), "")
}

// CreateArticleUpdate creates a new article.update model and returns its id.
func (c *Client) CreateArticleUpdate(au *ArticleUpdate) (int64, error) {
	ids, err := c.CreateArticleUpdates([]*ArticleUpdate{au})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateArticleUpdate creates a new article.update model and returns its id.
func (c *Client) CreateArticleUpdates(aus []*ArticleUpdate) ([]int64, error) {
	var vv []interface{}
	for _, v := range aus {
		vv = append(vv, v)
	}
	return c.Create(ArticleUpdateModel, vv, nil)
}

// UpdateArticleUpdate updates an existing article.update record.
func (c *Client) UpdateArticleUpdate(au *ArticleUpdate) error {
	return c.UpdateArticleUpdates([]int64{au.Id.Get()}, au)
}

// UpdateArticleUpdates updates existing article.update records.
// All records (represented by ids) will be updated by au values.
func (c *Client) UpdateArticleUpdates(ids []int64, au *ArticleUpdate) error {
	return c.Update(ArticleUpdateModel, ids, au, nil)
}

// DeleteArticleUpdate deletes an existing article.update record.
func (c *Client) DeleteArticleUpdate(id int64) error {
	return c.DeleteArticleUpdates([]int64{id})
}

// DeleteArticleUpdates deletes existing article.update records.
func (c *Client) DeleteArticleUpdates(ids []int64) error {
	return c.Delete(ArticleUpdateModel, ids)
}

// GetArticleUpdate gets article.update existing record.
func (c *Client) GetArticleUpdate(id int64) (*ArticleUpdate, error) {
	aus, err := c.GetArticleUpdates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aus)[0]), nil
}

// GetArticleUpdates gets article.update existing records.
func (c *Client) GetArticleUpdates(ids []int64) (*ArticleUpdates, error) {
	aus := &ArticleUpdates{}
	if err := c.Read(ArticleUpdateModel, ids, nil, aus); err != nil {
		return nil, err
	}
	return aus, nil
}

// FindArticleUpdate finds article.update record by querying it with criteria.
func (c *Client) FindArticleUpdate(criteria *Criteria) (*ArticleUpdate, error) {
	aus := &ArticleUpdates{}
	if err := c.SearchRead(ArticleUpdateModel, criteria, NewOptions().Limit(1), aus); err != nil {
		return nil, err
	}
	return &((*aus)[0]), nil
}

// FindArticleUpdates finds article.update records by querying it
// and filtering it with criteria and options.
func (c *Client) FindArticleUpdates(criteria *Criteria, options *Options) (*ArticleUpdates, error) {
	aus := &ArticleUpdates{}
	if err := c.SearchRead(ArticleUpdateModel, criteria, options, aus); err != nil {
		return nil, err
	}
	return aus, nil
}

// FindArticleUpdateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindArticleUpdateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ArticleUpdateModel, criteria, options)
}

// FindArticleUpdateId finds record id by querying it with criteria.
func (c *Client) FindArticleUpdateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ArticleUpdateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
