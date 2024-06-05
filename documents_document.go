package odoo

import (
	"fmt"
)

// DocumentsDocument represents documents.document model.
type DocumentsDocument struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttachmentId                *Many2One  `xmlrpc:"attachment_id,omptempty"`
	AttachmentName              *String    `xmlrpc:"attachment_name,omptempty"`
	AttachmentType              *Selection `xmlrpc:"attachment_type,omptempty"`
	AvailableRuleIds            *Relation  `xmlrpc:"available_rule_ids,omptempty"`
	Checksum                    *String    `xmlrpc:"checksum,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateShareId               *Many2One  `xmlrpc:"create_share_id,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Datas                       *String    `xmlrpc:"datas,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omptempty"`
	FavoritedIds                *Relation  `xmlrpc:"favorited_ids,omptempty"`
	FileSize                    *Int       `xmlrpc:"file_size,omptempty"`
	FolderId                    *Many2One  `xmlrpc:"folder_id,omptempty"`
	GroupIds                    *Relation  `xmlrpc:"group_ids,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IndexContent                *String    `xmlrpc:"index_content,omptempty"`
	IsLocked                    *Bool      `xmlrpc:"is_locked,omptempty"`
	LockUid                     *Many2One  `xmlrpc:"lock_uid,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mimetype                    *String    `xmlrpc:"mimetype,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	OwnerId                     *Many2One  `xmlrpc:"owner_id,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	RequestActivityId           *Many2One  `xmlrpc:"request_activity_id,omptempty"`
	ResId                       *Int       `xmlrpc:"res_id,omptempty"`
	ResModel                    *String    `xmlrpc:"res_model,omptempty"`
	ResModelName                *String    `xmlrpc:"res_model_name,omptempty"`
	ResName                     *String    `xmlrpc:"res_name,omptempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omptempty"`
	Thumbnail                   *String    `xmlrpc:"thumbnail,omptempty"`
	Type                        *Selection `xmlrpc:"type,omptempty"`
	Url                         *String    `xmlrpc:"url,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(DocumentsDocumentModel, vv)
}

// UpdateDocumentsDocument updates an existing documents.document record.
func (c *Client) UpdateDocumentsDocument(dd *DocumentsDocument) error {
	return c.UpdateDocumentsDocuments([]int64{dd.Id.Get()}, dd)
}

// UpdateDocumentsDocuments updates existing documents.document records.
// All records (represented by ids) will be updated by dd values.
func (c *Client) UpdateDocumentsDocuments(ids []int64, dd *DocumentsDocument) error {
	return c.Update(DocumentsDocumentModel, ids, dd)
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
	if dds != nil && len(*dds) > 0 {
		return &((*dds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.document not found", id)
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
	if dds != nil && len(*dds) > 0 {
		return &((*dds)[0]), nil
	}
	return nil, fmt.Errorf("documents.document was not found with criteria %v", criteria)
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
	ids, err := c.Search(DocumentsDocumentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsDocumentId finds record id by querying it with criteria.
func (c *Client) FindDocumentsDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.document was not found with criteria %v and options %v", criteria, options)
}
