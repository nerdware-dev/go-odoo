package odoo

// DocumentsFacet represents documents.facet model.
type DocumentsFacet struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FolderId    *Many2One `xmlrpc:"folder_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	TagIds      *Relation `xmlrpc:"tag_ids,omitempty"`
	Tooltip     *String   `xmlrpc:"tooltip,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DocumentsFacets represents array of documents.facet model.
type DocumentsFacets []DocumentsFacet

// DocumentsFacetModel is the odoo model name.
const DocumentsFacetModel = "documents.facet"

// Many2One convert DocumentsFacet to *Many2One.
func (df *DocumentsFacet) Many2One() *Many2One {
	return NewMany2One(df.Id.Get(), "")
}

// CreateDocumentsFacet creates a new documents.facet model and returns its id.
func (c *Client) CreateDocumentsFacet(df *DocumentsFacet) (int64, error) {
	ids, err := c.CreateDocumentsFacets([]*DocumentsFacet{df})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsFacet creates a new documents.facet model and returns its id.
func (c *Client) CreateDocumentsFacets(dfs []*DocumentsFacet) ([]int64, error) {
	var vv []interface{}
	for _, v := range dfs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsFacetModel, vv, nil)
}

// UpdateDocumentsFacet updates an existing documents.facet record.
func (c *Client) UpdateDocumentsFacet(df *DocumentsFacet) error {
	return c.UpdateDocumentsFacets([]int64{df.Id.Get()}, df)
}

// UpdateDocumentsFacets updates existing documents.facet records.
// All records (represented by ids) will be updated by df values.
func (c *Client) UpdateDocumentsFacets(ids []int64, df *DocumentsFacet) error {
	return c.Update(DocumentsFacetModel, ids, df, nil)
}

// DeleteDocumentsFacet deletes an existing documents.facet record.
func (c *Client) DeleteDocumentsFacet(id int64) error {
	return c.DeleteDocumentsFacets([]int64{id})
}

// DeleteDocumentsFacets deletes existing documents.facet records.
func (c *Client) DeleteDocumentsFacets(ids []int64) error {
	return c.Delete(DocumentsFacetModel, ids)
}

// GetDocumentsFacet gets documents.facet existing record.
func (c *Client) GetDocumentsFacet(id int64) (*DocumentsFacet, error) {
	dfs, err := c.GetDocumentsFacets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dfs)[0]), nil
}

// GetDocumentsFacets gets documents.facet existing records.
func (c *Client) GetDocumentsFacets(ids []int64) (*DocumentsFacets, error) {
	dfs := &DocumentsFacets{}
	if err := c.Read(DocumentsFacetModel, ids, nil, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDocumentsFacet finds documents.facet record by querying it with criteria.
func (c *Client) FindDocumentsFacet(criteria *Criteria) (*DocumentsFacet, error) {
	dfs := &DocumentsFacets{}
	if err := c.SearchRead(DocumentsFacetModel, criteria, NewOptions().Limit(1), dfs); err != nil {
		return nil, err
	}
	return &((*dfs)[0]), nil
}

// FindDocumentsFacets finds documents.facet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFacets(criteria *Criteria, options *Options) (*DocumentsFacets, error) {
	dfs := &DocumentsFacets{}
	if err := c.SearchRead(DocumentsFacetModel, criteria, options, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDocumentsFacetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFacetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsFacetModel, criteria, options)
}

// FindDocumentsFacetId finds record id by querying it with criteria.
func (c *Client) FindDocumentsFacetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsFacetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
