package odoo

// DocumentsDocument represents documents.document model.
type DocumentsDocument struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttachmentId                *Many2One  `xmlrpc:"attachment_id,omitempty"`
	AttachmentName              *String    `xmlrpc:"attachment_name,omitempty"`
	AttachmentType              *Selection `xmlrpc:"attachment_type,omitempty"`
	AvailableRuleIds            *Relation  `xmlrpc:"available_rule_ids,omitempty"`
	Checksum                    *String    `xmlrpc:"checksum,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateShareId               *Many2One  `xmlrpc:"create_share_id,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Datas                       *String    `xmlrpc:"datas,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty"`
	FavoritedIds                *Relation  `xmlrpc:"favorited_ids,omitempty"`
	FileSize                    *Int       `xmlrpc:"file_size,omitempty"`
	FolderId                    *Many2One  `xmlrpc:"folder_id,omitempty"`
	GroupIds                    *Relation  `xmlrpc:"group_ids,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IndexContent                *String    `xmlrpc:"index_content,omitempty"`
	IsLocked                    *Bool      `xmlrpc:"is_locked,omitempty"`
	LockUid                     *Many2One  `xmlrpc:"lock_uid,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Mimetype                    *String    `xmlrpc:"mimetype,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	OwnerId                     *Many2One  `xmlrpc:"owner_id,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	RequestActivityId           *Many2One  `xmlrpc:"request_activity_id,omitempty"`
	ResId                       *Int       `xmlrpc:"res_id,omitempty"`
	ResModel                    *String    `xmlrpc:"res_model,omitempty"`
	ResModelName                *String    `xmlrpc:"res_model_name,omitempty"`
	ResName                     *String    `xmlrpc:"res_name,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	Thumbnail                   *String    `xmlrpc:"thumbnail,omitempty"`
	Type                        *Selection `xmlrpc:"type,omitempty"`
	Url                         *String    `xmlrpc:"url,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return ids[0], nil
}
