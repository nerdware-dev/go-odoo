package odoo

// KnowsystemArticleRevision represents knowsystem.article.revision model.
type KnowsystemArticleRevision struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	ArticleId         *Many2One `xmlrpc:"article_id,omitempty"`
	AttachmentIds     *Relation `xmlrpc:"attachment_ids,omitempty"`
	AttachmentsChange *Bool     `xmlrpc:"attachments_change,omitempty"`
	AuthorId          *Many2One `xmlrpc:"author_id,omitempty"`
	ChangeDatetime    *Time     `xmlrpc:"change_datetime,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	Description       *String   `xmlrpc:"description,omitempty"`
	DescriptionChange *Int      `xmlrpc:"description_change,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	Name              *String   `xmlrpc:"name,omitempty"`
	NameChange        *Bool     `xmlrpc:"name_change,omitempty"`
	SectionChange     *Bool     `xmlrpc:"section_change,omitempty"`
	SectionId         *Many2One `xmlrpc:"section_id,omitempty"`
	TagIds            *Relation `xmlrpc:"tag_ids,omitempty"`
	TagsChange        *Bool     `xmlrpc:"tags_change,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowsystemArticleRevisions represents array of knowsystem.article.revision model.
type KnowsystemArticleRevisions []KnowsystemArticleRevision

// KnowsystemArticleRevisionModel is the odoo model name.
const KnowsystemArticleRevisionModel = "knowsystem.article.revision"

// Many2One convert KnowsystemArticleRevision to *Many2One.
func (kar *KnowsystemArticleRevision) Many2One() *Many2One {
	return NewMany2One(kar.Id.Get(), "")
}

// CreateKnowsystemArticleRevision creates a new knowsystem.article.revision model and returns its id.
func (c *Client) CreateKnowsystemArticleRevision(kar *KnowsystemArticleRevision) (int64, error) {
	ids, err := c.CreateKnowsystemArticleRevisions([]*KnowsystemArticleRevision{kar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowsystemArticleRevision creates a new knowsystem.article.revision model and returns its id.
func (c *Client) CreateKnowsystemArticleRevisions(kars []*KnowsystemArticleRevision) ([]int64, error) {
	var vv []interface{}
	for _, v := range kars {
		vv = append(vv, v)
	}
	return c.Create(KnowsystemArticleRevisionModel, vv, nil)
}

// UpdateKnowsystemArticleRevision updates an existing knowsystem.article.revision record.
func (c *Client) UpdateKnowsystemArticleRevision(kar *KnowsystemArticleRevision) error {
	return c.UpdateKnowsystemArticleRevisions([]int64{kar.Id.Get()}, kar)
}

// UpdateKnowsystemArticleRevisions updates existing knowsystem.article.revision records.
// All records (represented by ids) will be updated by kar values.
func (c *Client) UpdateKnowsystemArticleRevisions(ids []int64, kar *KnowsystemArticleRevision) error {
	return c.Update(KnowsystemArticleRevisionModel, ids, kar, nil)
}

// DeleteKnowsystemArticleRevision deletes an existing knowsystem.article.revision record.
func (c *Client) DeleteKnowsystemArticleRevision(id int64) error {
	return c.DeleteKnowsystemArticleRevisions([]int64{id})
}

// DeleteKnowsystemArticleRevisions deletes existing knowsystem.article.revision records.
func (c *Client) DeleteKnowsystemArticleRevisions(ids []int64) error {
	return c.Delete(KnowsystemArticleRevisionModel, ids)
}

// GetKnowsystemArticleRevision gets knowsystem.article.revision existing record.
func (c *Client) GetKnowsystemArticleRevision(id int64) (*KnowsystemArticleRevision, error) {
	kars, err := c.GetKnowsystemArticleRevisions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kars)[0]), nil
}

// GetKnowsystemArticleRevisions gets knowsystem.article.revision existing records.
func (c *Client) GetKnowsystemArticleRevisions(ids []int64) (*KnowsystemArticleRevisions, error) {
	kars := &KnowsystemArticleRevisions{}
	if err := c.Read(KnowsystemArticleRevisionModel, ids, nil, kars); err != nil {
		return nil, err
	}
	return kars, nil
}

// FindKnowsystemArticleRevision finds knowsystem.article.revision record by querying it with criteria.
func (c *Client) FindKnowsystemArticleRevision(criteria *Criteria) (*KnowsystemArticleRevision, error) {
	kars := &KnowsystemArticleRevisions{}
	if err := c.SearchRead(KnowsystemArticleRevisionModel, criteria, NewOptions().Limit(1), kars); err != nil {
		return nil, err
	}
	return &((*kars)[0]), nil
}

// FindKnowsystemArticleRevisions finds knowsystem.article.revision records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleRevisions(criteria *Criteria, options *Options) (*KnowsystemArticleRevisions, error) {
	kars := &KnowsystemArticleRevisions{}
	if err := c.SearchRead(KnowsystemArticleRevisionModel, criteria, options, kars); err != nil {
		return nil, err
	}
	return kars, nil
}

// FindKnowsystemArticleRevisionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowsystemArticleRevisionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowsystemArticleRevisionModel, criteria, options)
}

// FindKnowsystemArticleRevisionId finds record id by querying it with criteria.
func (c *Client) FindKnowsystemArticleRevisionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowsystemArticleRevisionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
