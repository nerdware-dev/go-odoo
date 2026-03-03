package odoo

// DocumentsDocument represents documents.document model.
type DocumentsDocument struct {
	Active                      *Bool      `xmlrpc:"active,omitempty" json:"active,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty" json:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty" json:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty" json:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty" json:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty" json:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty" json:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty" json:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty" json:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty" json:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty" json:"activity_user_id,omitempty"`
	AttachmentId                *Many2One  `xmlrpc:"attachment_id,omitempty" json:"attachment_id,omitempty"`
	AttachmentName              *String    `xmlrpc:"attachment_name,omitempty" json:"attachment_name,omitempty"`
	AttachmentType              *Selection `xmlrpc:"attachment_type,omitempty" json:"attachment_type,omitempty"`
	AvailableRuleIds            *Relation  `xmlrpc:"available_rule_ids,omitempty" json:"available_rule_ids,omitempty"`
	Checksum                    *String    `xmlrpc:"checksum,omitempty" json:"checksum,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty" json:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateShareId               *Many2One  `xmlrpc:"create_share_id,omitempty" json:"create_share_id,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Datas                       *String    `xmlrpc:"datas,omitempty" json:"datas,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty" json:"email_cc,omitempty"`
	FavoritedIds                *Relation  `xmlrpc:"favorited_ids,omitempty" json:"favorited_ids,omitempty"`
	FileExtension               *String    `xmlrpc:"file_extension,omitempty" json:"file_extension,omitempty"`
	FileSize                    *Int       `xmlrpc:"file_size,omitempty" json:"file_size,omitempty"`
	FolderId                    *Many2One  `xmlrpc:"folder_id,omitempty" json:"folder_id,omitempty"`
	GroupIds                    *Relation  `xmlrpc:"group_ids,omitempty" json:"group_ids,omitempty"`
	Handler                     *Selection `xmlrpc:"handler,omitempty" json:"handler,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty" json:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IndexContent                *String    `xmlrpc:"index_content,omitempty" json:"index_content,omitempty"`
	IsEditableAttachment        *Bool      `xmlrpc:"is_editable_attachment,omitempty" json:"is_editable_attachment,omitempty"`
	IsFavorited                 *Bool      `xmlrpc:"is_favorited,omitempty" json:"is_favorited,omitempty"`
	IsLocked                    *Bool      `xmlrpc:"is_locked,omitempty" json:"is_locked,omitempty"`
	IsMultipage                 *Bool      `xmlrpc:"is_multipage,omitempty" json:"is_multipage,omitempty"`
	IsShared                    *Bool      `xmlrpc:"is_shared,omitempty" json:"is_shared,omitempty"`
	LockUid                     *Many2One  `xmlrpc:"lock_uid,omitempty" json:"lock_uid,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty" json:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty" json:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty" json:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty" json:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty" json:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty" json:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty" json:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty" json:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty" json:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty" json:"message_partner_ids,omitempty"`
	Mimetype                    *String    `xmlrpc:"mimetype,omitempty" json:"mimetype,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty" json:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty" json:"name,omitempty"`
	OwnerId                     *Many2One  `xmlrpc:"owner_id,omitempty" json:"owner_id,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty" json:"partner_id,omitempty"`
	PreviousAttachmentIds       *Relation  `xmlrpc:"previous_attachment_ids,omitempty" json:"previous_attachment_ids,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty" json:"product_id,omitempty"`
	ProductTemplateId           *Many2One  `xmlrpc:"product_template_id,omitempty" json:"product_template_id,omitempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omitempty" json:"project_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	Raw                         *String    `xmlrpc:"raw,omitempty" json:"raw,omitempty"`
	RequestActivityId           *Many2One  `xmlrpc:"request_activity_id,omitempty" json:"request_activity_id,omitempty"`
	ResId                       *Many2One  `xmlrpc:"res_id,omitempty" json:"res_id,omitempty"`
	ResModel                    *String    `xmlrpc:"res_model,omitempty" json:"res_model,omitempty"`
	ResModelName                *String    `xmlrpc:"res_model_name,omitempty" json:"res_model_name,omitempty"`
	ResName                     *String    `xmlrpc:"res_name,omitempty" json:"res_name,omitempty"`
	ServerRevisionId            *String    `xmlrpc:"server_revision_id,omitempty" json:"server_revision_id,omitempty"`
	SpreadsheetBinaryData       *String    `xmlrpc:"spreadsheet_binary_data,omitempty" json:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData             *String    `xmlrpc:"spreadsheet_data,omitempty" json:"spreadsheet_data,omitempty"`
	SpreadsheetRevisionIds      *Relation  `xmlrpc:"spreadsheet_revision_ids,omitempty" json:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot         *String    `xmlrpc:"spreadsheet_snapshot,omitempty" json:"spreadsheet_snapshot,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty" json:"tag_ids,omitempty"`
	TaskId                      *Many2One  `xmlrpc:"task_id,omitempty" json:"task_id,omitempty"`
	Thumbnail                   *String    `xmlrpc:"thumbnail,omitempty" json:"thumbnail,omitempty"`
	ThumbnailStatus             *Selection `xmlrpc:"thumbnail_status,omitempty" json:"thumbnail_status,omitempty"`
	Type                        *Selection `xmlrpc:"type,omitempty" json:"type,omitempty"`
	Url                         *String    `xmlrpc:"url,omitempty" json:"url,omitempty"`
	UrlPreviewImage             *String    `xmlrpc:"url_preview_image,omitempty" json:"url_preview_image,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty" json:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// DocumentsDocuments represents array of documents.document model.
type DocumentsDocuments []DocumentsDocument

// DocumentsDocumentModel is the odoo model name.
const DocumentsDocumentModel = "documents.document"

// Many2One convert DocumentsDocument to *Many2One.
func (dd *DocumentsDocument) Many2One() *Many2One {
	return NewMany2One(dd.Id.Get(), "")
}

// CreateDocumentsDocument creates a new documents.document model and returns its id.
func (c *Client) CreateDocumentsDocument(dd *DocumentsDocument) (int64, error) {
	ids, err := c.CreateDocumentsDocuments([]*DocumentsDocument{dd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsDocument creates a new documents.document model and returns its id.
func (c *Client) CreateDocumentsDocuments(dds []*DocumentsDocument) ([]int64, error) {
	var vv []interface{}
	for _, v := range dds {
		vv = append(vv, v)
	}
	return c.Create(DocumentsDocumentModel, vv, nil)
}

// UpdateDocumentsDocument updates an existing documents.document record.
func (c *Client) UpdateDocumentsDocument(dd *DocumentsDocument) error {
	return c.UpdateDocumentsDocuments([]int64{dd.Id.Get()}, dd)
}

// UpdateDocumentsDocuments updates existing documents.document records.
// All records (represented by ids) will be updated by dd values.
func (c *Client) UpdateDocumentsDocuments(ids []int64, dd *DocumentsDocument) error {
	return c.Update(DocumentsDocumentModel, ids, dd, nil)
}

// DeleteDocumentsDocument deletes an existing documents.document record.
func (c *Client) DeleteDocumentsDocument(id int64) error {
	return c.DeleteDocumentsDocuments([]int64{id})
}

// DeleteDocumentsDocuments deletes existing documents.document records.
func (c *Client) DeleteDocumentsDocuments(ids []int64) error {
	return c.Delete(DocumentsDocumentModel, ids)
}

// GetDocumentsDocument gets documents.document existing record.
func (c *Client) GetDocumentsDocument(id int64) (*DocumentsDocument, error) {
	dds, err := c.GetDocumentsDocuments([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*dds) == 0 {
		return nil, nil
	}
	return &((*dds)[0]), nil
}

// GetDocumentsDocuments gets documents.document existing records.
func (c *Client) GetDocumentsDocuments(ids []int64) (*DocumentsDocuments, error) {
	dds := &DocumentsDocuments{}
	if err := c.Read(DocumentsDocumentModel, ids, nil, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDocumentsDocument finds documents.document record by querying it with criteria.
func (c *Client) FindDocumentsDocument(criteria *Criteria) (*DocumentsDocument, error) {
	dds := &DocumentsDocuments{}
	if err := c.SearchRead(DocumentsDocumentModel, criteria, NewOptions().Limit(1), dds); err != nil {
		return nil, err
	}
	if len(*dds) == 0 {
		return nil, nil
	}
	return &((*dds)[0]), nil
}

// FindDocumentsDocuments finds documents.document records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsDocuments(criteria *Criteria, options *Options) (*DocumentsDocuments, error) {
	dds := &DocumentsDocuments{}
	if err := c.SearchRead(DocumentsDocumentModel, criteria, options, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDocumentsDocumentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsDocumentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsDocumentModel, criteria, options)
}

// FindDocumentsDocumentId finds record id by querying it with criteria.
func (c *Client) FindDocumentsDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
