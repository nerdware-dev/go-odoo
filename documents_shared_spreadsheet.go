package odoo

// DocumentsSharedSpreadsheet represents documents.shared.spreadsheet model.
type DocumentsSharedSpreadsheet struct {
	CreateDate             *Time     `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentId             *Many2One `xmlrpc:"document_id,omitempty" json:"document_id,omitempty"`
	ExcelExport            *String   `xmlrpc:"excel_export,omitempty" json:"excel_export,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty" json:"id,omitempty"`
	ServerRevisionId       *String   `xmlrpc:"server_revision_id,omitempty" json:"server_revision_id,omitempty"`
	ShareId                *Many2One `xmlrpc:"share_id,omitempty" json:"share_id,omitempty"`
	SpreadsheetBinaryData  *String   `xmlrpc:"spreadsheet_binary_data,omitempty" json:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData        *String   `xmlrpc:"spreadsheet_data,omitempty" json:"spreadsheet_data,omitempty"`
	SpreadsheetRevisionIds *Relation `xmlrpc:"spreadsheet_revision_ids,omitempty" json:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot    *String   `xmlrpc:"spreadsheet_snapshot,omitempty" json:"spreadsheet_snapshot,omitempty"`
	Thumbnail              *String   `xmlrpc:"thumbnail,omitempty" json:"thumbnail,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// DocumentsSharedSpreadsheets represents array of documents.shared.spreadsheet model.
type DocumentsSharedSpreadsheets []DocumentsSharedSpreadsheet

// DocumentsSharedSpreadsheetModel is the odoo model name.
const DocumentsSharedSpreadsheetModel = "documents.shared.spreadsheet"

// Many2One convert DocumentsSharedSpreadsheet to *Many2One.
func (dss *DocumentsSharedSpreadsheet) Many2One() *Many2One {
	return NewMany2One(dss.Id.Get(), "")
}

// CreateDocumentsSharedSpreadsheet creates a new documents.shared.spreadsheet model and returns its id.
func (c *Client) CreateDocumentsSharedSpreadsheet(dss *DocumentsSharedSpreadsheet) (int64, error) {
	ids, err := c.CreateDocumentsSharedSpreadsheets([]*DocumentsSharedSpreadsheet{dss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsSharedSpreadsheet creates a new documents.shared.spreadsheet model and returns its id.
func (c *Client) CreateDocumentsSharedSpreadsheets(dsss []*DocumentsSharedSpreadsheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range dsss {
		vv = append(vv, v)
	}
	return c.Create(DocumentsSharedSpreadsheetModel, vv, nil)
}

// UpdateDocumentsSharedSpreadsheet updates an existing documents.shared.spreadsheet record.
func (c *Client) UpdateDocumentsSharedSpreadsheet(dss *DocumentsSharedSpreadsheet) error {
	return c.UpdateDocumentsSharedSpreadsheets([]int64{dss.Id.Get()}, dss)
}

// UpdateDocumentsSharedSpreadsheets updates existing documents.shared.spreadsheet records.
// All records (represented by ids) will be updated by dss values.
func (c *Client) UpdateDocumentsSharedSpreadsheets(ids []int64, dss *DocumentsSharedSpreadsheet) error {
	return c.Update(DocumentsSharedSpreadsheetModel, ids, dss, nil)
}

// DeleteDocumentsSharedSpreadsheet deletes an existing documents.shared.spreadsheet record.
func (c *Client) DeleteDocumentsSharedSpreadsheet(id int64) error {
	return c.DeleteDocumentsSharedSpreadsheets([]int64{id})
}

// DeleteDocumentsSharedSpreadsheets deletes existing documents.shared.spreadsheet records.
func (c *Client) DeleteDocumentsSharedSpreadsheets(ids []int64) error {
	return c.Delete(DocumentsSharedSpreadsheetModel, ids)
}

// GetDocumentsSharedSpreadsheet gets documents.shared.spreadsheet existing record.
func (c *Client) GetDocumentsSharedSpreadsheet(id int64) (*DocumentsSharedSpreadsheet, error) {
	dsss, err := c.GetDocumentsSharedSpreadsheets([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*dsss) == 0 {
		return nil, nil
	}
	return &((*dsss)[0]), nil
}

// GetDocumentsSharedSpreadsheets gets documents.shared.spreadsheet existing records.
func (c *Client) GetDocumentsSharedSpreadsheets(ids []int64) (*DocumentsSharedSpreadsheets, error) {
	dsss := &DocumentsSharedSpreadsheets{}
	if err := c.Read(DocumentsSharedSpreadsheetModel, ids, nil, dsss); err != nil {
		return nil, err
	}
	return dsss, nil
}

// FindDocumentsSharedSpreadsheet finds documents.shared.spreadsheet record by querying it with criteria.
func (c *Client) FindDocumentsSharedSpreadsheet(criteria *Criteria) (*DocumentsSharedSpreadsheet, error) {
	dsss := &DocumentsSharedSpreadsheets{}
	if err := c.SearchRead(DocumentsSharedSpreadsheetModel, criteria, NewOptions().Limit(1), dsss); err != nil {
		return nil, err
	}
	if len(*dsss) == 0 {
		return nil, nil
	}
	return &((*dsss)[0]), nil
}

// FindDocumentsSharedSpreadsheets finds documents.shared.spreadsheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsSharedSpreadsheets(criteria *Criteria, options *Options) (*DocumentsSharedSpreadsheets, error) {
	dsss := &DocumentsSharedSpreadsheets{}
	if err := c.SearchRead(DocumentsSharedSpreadsheetModel, criteria, options, dsss); err != nil {
		return nil, err
	}
	return dsss, nil
}

// FindDocumentsSharedSpreadsheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsSharedSpreadsheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsSharedSpreadsheetModel, criteria, options)
}

// FindDocumentsSharedSpreadsheetId finds record id by querying it with criteria.
func (c *Client) FindDocumentsSharedSpreadsheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsSharedSpreadsheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
