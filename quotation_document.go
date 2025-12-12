package odoo

// QuotationDocument represents quotation.document model.
type QuotationDocument struct {
	AccessToken          *String    `xmlrpc:"access_token,omitempty" json:"access_token,omitempty"`
	Active               *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	Checksum             *String    `xmlrpc:"checksum,omitempty" json:"checksum,omitempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Datas                *String    `xmlrpc:"datas,omitempty" json:"datas,omitempty"`
	DbDatas              *String    `xmlrpc:"db_datas,omitempty" json:"db_datas,omitempty"`
	Description          *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	DocumentType         *Selection `xmlrpc:"document_type,omitempty" json:"document_type,omitempty"`
	FileSize             *Int       `xmlrpc:"file_size,omitempty" json:"file_size,omitempty"`
	FormFieldIds         *Relation  `xmlrpc:"form_field_ids,omitempty" json:"form_field_ids,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	ImageHeight          *Int       `xmlrpc:"image_height,omitempty" json:"image_height,omitempty"`
	ImageSrc             *String    `xmlrpc:"image_src,omitempty" json:"image_src,omitempty"`
	ImageWidth           *Int       `xmlrpc:"image_width,omitempty" json:"image_width,omitempty"`
	IndexContent         *String    `xmlrpc:"index_content,omitempty" json:"index_content,omitempty"`
	IrAttachmentId       *Many2One  `xmlrpc:"ir_attachment_id,omitempty" json:"ir_attachment_id,omitempty"`
	LocalUrl             *String    `xmlrpc:"local_url,omitempty" json:"local_url,omitempty"`
	Mimetype             *String    `xmlrpc:"mimetype,omitempty" json:"mimetype,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OriginalId           *Many2One  `xmlrpc:"original_id,omitempty" json:"original_id,omitempty"`
	Public               *Bool      `xmlrpc:"public,omitempty" json:"public,omitempty"`
	QuotationTemplateIds *Relation  `xmlrpc:"quotation_template_ids,omitempty" json:"quotation_template_ids,omitempty"`
	Raw                  *String    `xmlrpc:"raw,omitempty" json:"raw,omitempty"`
	ResField             *String    `xmlrpc:"res_field,omitempty" json:"res_field,omitempty"`
	ResId                *Many2One  `xmlrpc:"res_id,omitempty" json:"res_id,omitempty"`
	ResModel             *String    `xmlrpc:"res_model,omitempty" json:"res_model,omitempty"`
	ResName              *String    `xmlrpc:"res_name,omitempty" json:"res_name,omitempty"`
	Sequence             *Int       `xmlrpc:"sequence,omitempty" json:"sequence,omitempty"`
	StoreFname           *String    `xmlrpc:"store_fname,omitempty" json:"store_fname,omitempty"`
	Type                 *Selection `xmlrpc:"type,omitempty" json:"type,omitempty"`
	Url                  *String    `xmlrpc:"url,omitempty" json:"url,omitempty"`
	VoiceIds             *Relation  `xmlrpc:"voice_ids,omitempty" json:"voice_ids,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// QuotationDocuments represents array of quotation.document model.
type QuotationDocuments []QuotationDocument

// QuotationDocumentModel is the odoo model name.
const QuotationDocumentModel = "quotation.document"

// Many2One convert QuotationDocument to *Many2One.
func (qd *QuotationDocument) Many2One() *Many2One {
	return NewMany2One(qd.Id.Get(), "")
}

// CreateQuotationDocument creates a new quotation.document model and returns its id.
func (c *Client) CreateQuotationDocument(qd *QuotationDocument) (int64, error) {
	ids, err := c.CreateQuotationDocuments([]*QuotationDocument{qd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQuotationDocument creates a new quotation.document model and returns its id.
func (c *Client) CreateQuotationDocuments(qds []*QuotationDocument) ([]int64, error) {
	var vv []interface{}
	for _, v := range qds {
		vv = append(vv, v)
	}
	return c.Create(QuotationDocumentModel, vv, nil)
}

// UpdateQuotationDocument updates an existing quotation.document record.
func (c *Client) UpdateQuotationDocument(qd *QuotationDocument) error {
	return c.UpdateQuotationDocuments([]int64{qd.Id.Get()}, qd)
}

// UpdateQuotationDocuments updates existing quotation.document records.
// All records (represented by ids) will be updated by qd values.
func (c *Client) UpdateQuotationDocuments(ids []int64, qd *QuotationDocument) error {
	return c.Update(QuotationDocumentModel, ids, qd, nil)
}

// DeleteQuotationDocument deletes an existing quotation.document record.
func (c *Client) DeleteQuotationDocument(id int64) error {
	return c.DeleteQuotationDocuments([]int64{id})
}

// DeleteQuotationDocuments deletes existing quotation.document records.
func (c *Client) DeleteQuotationDocuments(ids []int64) error {
	return c.Delete(QuotationDocumentModel, ids)
}

// GetQuotationDocument gets quotation.document existing record.
func (c *Client) GetQuotationDocument(id int64) (*QuotationDocument, error) {
	qds, err := c.GetQuotationDocuments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qds)[0]), nil
}

// GetQuotationDocuments gets quotation.document existing records.
func (c *Client) GetQuotationDocuments(ids []int64) (*QuotationDocuments, error) {
	qds := &QuotationDocuments{}
	if err := c.Read(QuotationDocumentModel, ids, nil, qds); err != nil {
		return nil, err
	}
	return qds, nil
}

// FindQuotationDocument finds quotation.document record by querying it with criteria.
func (c *Client) FindQuotationDocument(criteria *Criteria) (*QuotationDocument, error) {
	qds := &QuotationDocuments{}
	if err := c.SearchRead(QuotationDocumentModel, criteria, NewOptions().Limit(1), qds); err != nil {
		return nil, err
	}
	return &((*qds)[0]), nil
}

// FindQuotationDocuments finds quotation.document records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQuotationDocuments(criteria *Criteria, options *Options) (*QuotationDocuments, error) {
	qds := &QuotationDocuments{}
	if err := c.SearchRead(QuotationDocumentModel, criteria, options, qds); err != nil {
		return nil, err
	}
	return qds, nil
}

// FindQuotationDocumentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQuotationDocumentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QuotationDocumentModel, criteria, options)
}

// FindQuotationDocumentId finds record id by querying it with criteria.
func (c *Client) FindQuotationDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QuotationDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
