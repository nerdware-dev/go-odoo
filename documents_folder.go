package odoo

// DocumentsFolder represents documents.folder model.
type DocumentsFolder struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	ActionCount       *Int      `xmlrpc:"action_count,omitempty"`
	ChildrenFolderIds *Relation `xmlrpc:"children_folder_ids,omitempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	Description       *String   `xmlrpc:"description,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	DocumentCount     *Int      `xmlrpc:"document_count,omitempty"`
	DocumentIds       *Relation `xmlrpc:"document_ids,omitempty"`
	FacetIds          *Relation `xmlrpc:"facet_ids,omitempty"`
	GroupIds          *Relation `xmlrpc:"group_ids,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	Name              *String   `xmlrpc:"name,omitempty"`
	ParentFolderId    *Many2One `xmlrpc:"parent_folder_id,omitempty"`
	ReadGroupIds      *Relation `xmlrpc:"read_group_ids,omitempty"`
	Sequence          *Int      `xmlrpc:"sequence,omitempty"`
	ShareLinkIds      *Relation `xmlrpc:"share_link_ids,omitempty"`
	UserSpecific      *Bool     `xmlrpc:"user_specific,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DocumentsFolders represents array of documents.folder model.
type DocumentsFolders []DocumentsFolder

// DocumentsFolderModel is the odoo model name.
const DocumentsFolderModel = "documents.folder"

// Many2One convert DocumentsFolder to *Many2One.
func (df *DocumentsFolder) Many2One() *Many2One {
	return NewMany2One(df.Id.Get(), "")
}

// CreateDocumentsFolder creates a new documents.folder model and returns its id.
func (c *Client) CreateDocumentsFolder(df *DocumentsFolder) (int64, error) {
	ids, err := c.CreateDocumentsFolders([]*DocumentsFolder{df})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsFolder creates a new documents.folder model and returns its id.
func (c *Client) CreateDocumentsFolders(dfs []*DocumentsFolder) ([]int64, error) {
	var vv []interface{}
	for _, v := range dfs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsFolderModel, vv, nil)
}

// UpdateDocumentsFolder updates an existing documents.folder record.
func (c *Client) UpdateDocumentsFolder(df *DocumentsFolder) error {
	return c.UpdateDocumentsFolders([]int64{df.Id.Get()}, df)
}

// UpdateDocumentsFolders updates existing documents.folder records.
// All records (represented by ids) will be updated by df values.
func (c *Client) UpdateDocumentsFolders(ids []int64, df *DocumentsFolder) error {
	return c.Update(DocumentsFolderModel, ids, df, nil)
}

// DeleteDocumentsFolder deletes an existing documents.folder record.
func (c *Client) DeleteDocumentsFolder(id int64) error {
	return c.DeleteDocumentsFolders([]int64{id})
}

// DeleteDocumentsFolders deletes existing documents.folder records.
func (c *Client) DeleteDocumentsFolders(ids []int64) error {
	return c.Delete(DocumentsFolderModel, ids)
}

// GetDocumentsFolder gets documents.folder existing record.
func (c *Client) GetDocumentsFolder(id int64) (*DocumentsFolder, error) {
	dfs, err := c.GetDocumentsFolders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dfs)[0]), nil
}

// GetDocumentsFolders gets documents.folder existing records.
func (c *Client) GetDocumentsFolders(ids []int64) (*DocumentsFolders, error) {
	dfs := &DocumentsFolders{}
	if err := c.Read(DocumentsFolderModel, ids, nil, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDocumentsFolder finds documents.folder record by querying it with criteria.
func (c *Client) FindDocumentsFolder(criteria *Criteria) (*DocumentsFolder, error) {
	dfs := &DocumentsFolders{}
	if err := c.SearchRead(DocumentsFolderModel, criteria, NewOptions().Limit(1), dfs); err != nil {
		return nil, err
	}
	return &((*dfs)[0]), nil
}

// FindDocumentsFolders finds documents.folder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFolders(criteria *Criteria, options *Options) (*DocumentsFolders, error) {
	dfs := &DocumentsFolders{}
	if err := c.SearchRead(DocumentsFolderModel, criteria, options, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDocumentsFolderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFolderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsFolderModel, criteria, options)
}

// FindDocumentsFolderId finds record id by querying it with criteria.
func (c *Client) FindDocumentsFolderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsFolderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
