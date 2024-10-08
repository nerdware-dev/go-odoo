package odoo

// DocumentsShare represents documents.share model.
type DocumentsShare struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                   *String    `xmlrpc:"access_token,omitempty"`
	Action                        *Selection `xmlrpc:"action,omitempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omitempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omitempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omitempty"`
	ActivityOption                *Bool      `xmlrpc:"activity_option,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AliasContact                  *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                 *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                   *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId            *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasId                       *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasModelId                  *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                     *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId            *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId           *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId                   *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	CanUpload                     *Bool      `xmlrpc:"can_upload,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateDeadline                  *Time      `xmlrpc:"date_deadline,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty"`
	DocumentIds                   *Relation  `xmlrpc:"document_ids,omitempty"`
	Domain                        *String    `xmlrpc:"domain,omitempty"`
	EmailDrop                     *Bool      `xmlrpc:"email_drop,omitempty"`
	FolderId                      *Many2One  `xmlrpc:"folder_id,omitempty"`
	FullUrl                       *String    `xmlrpc:"full_url,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty"`
	OwnerId                       *Many2One  `xmlrpc:"owner_id,omitempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omitempty"`
	State                         *Selection `xmlrpc:"state,omitempty"`
	TagIds                        *Relation  `xmlrpc:"tag_ids,omitempty"`
	Type                          *Selection `xmlrpc:"type,omitempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DocumentsShares represents array of documents.share model.
type DocumentsShares []DocumentsShare

// DocumentsShareModel is the odoo model name.
const DocumentsShareModel = "documents.share"

// Many2One convert DocumentsShare to *Many2One.
func (ds *DocumentsShare) Many2One() *Many2One {
	return NewMany2One(ds.Id.Get(), "")
}

// CreateDocumentsShare creates a new documents.share model and returns its id.
func (c *Client) CreateDocumentsShare(ds *DocumentsShare) (int64, error) {
	ids, err := c.CreateDocumentsShares([]*DocumentsShare{ds})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsShare creates a new documents.share model and returns its id.
func (c *Client) CreateDocumentsShares(dss []*DocumentsShare) ([]int64, error) {
	var vv []interface{}
	for _, v := range dss {
		vv = append(vv, v)
	}
	return c.Create(DocumentsShareModel, vv, nil)
}

// UpdateDocumentsShare updates an existing documents.share record.
func (c *Client) UpdateDocumentsShare(ds *DocumentsShare) error {
	return c.UpdateDocumentsShares([]int64{ds.Id.Get()}, ds)
}

// UpdateDocumentsShares updates existing documents.share records.
// All records (represented by ids) will be updated by ds values.
func (c *Client) UpdateDocumentsShares(ids []int64, ds *DocumentsShare) error {
	return c.Update(DocumentsShareModel, ids, ds, nil)
}

// DeleteDocumentsShare deletes an existing documents.share record.
func (c *Client) DeleteDocumentsShare(id int64) error {
	return c.DeleteDocumentsShares([]int64{id})
}

// DeleteDocumentsShares deletes existing documents.share records.
func (c *Client) DeleteDocumentsShares(ids []int64) error {
	return c.Delete(DocumentsShareModel, ids)
}

// GetDocumentsShare gets documents.share existing record.
func (c *Client) GetDocumentsShare(id int64) (*DocumentsShare, error) {
	dss, err := c.GetDocumentsShares([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dss)[0]), nil
}

// GetDocumentsShares gets documents.share existing records.
func (c *Client) GetDocumentsShares(ids []int64) (*DocumentsShares, error) {
	dss := &DocumentsShares{}
	if err := c.Read(DocumentsShareModel, ids, nil, dss); err != nil {
		return nil, err
	}
	return dss, nil
}

// FindDocumentsShare finds documents.share record by querying it with criteria.
func (c *Client) FindDocumentsShare(criteria *Criteria) (*DocumentsShare, error) {
	dss := &DocumentsShares{}
	if err := c.SearchRead(DocumentsShareModel, criteria, NewOptions().Limit(1), dss); err != nil {
		return nil, err
	}
	return &((*dss)[0]), nil
}

// FindDocumentsShares finds documents.share records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsShares(criteria *Criteria, options *Options) (*DocumentsShares, error) {
	dss := &DocumentsShares{}
	if err := c.SearchRead(DocumentsShareModel, criteria, options, dss); err != nil {
		return nil, err
	}
	return dss, nil
}

// FindDocumentsShareIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsShareIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsShareModel, criteria, options)
}

// FindDocumentsShareId finds record id by querying it with criteria.
func (c *Client) FindDocumentsShareId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsShareModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
