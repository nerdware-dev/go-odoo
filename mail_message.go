package odoo

// MailMessage represents mail.message model.
type MailMessage struct {
	AccountAuditLogAccountId   *Many2One  `xmlrpc:"account_audit_log_account_id,omitempty" json:"account_audit_log_account_id,omitempty"`
	AccountAuditLogCompanyId   *Many2One  `xmlrpc:"account_audit_log_company_id,omitempty" json:"account_audit_log_company_id,omitempty"`
	AccountAuditLogDisplayName *String    `xmlrpc:"account_audit_log_display_name,omitempty" json:"account_audit_log_display_name,omitempty"`
	AccountAuditLogMoveId      *Many2One  `xmlrpc:"account_audit_log_move_id,omitempty" json:"account_audit_log_move_id,omitempty"`
	AccountAuditLogPartnerId   *Many2One  `xmlrpc:"account_audit_log_partner_id,omitempty" json:"account_audit_log_partner_id,omitempty"`
	AccountAuditLogPreview     *String    `xmlrpc:"account_audit_log_preview,omitempty" json:"account_audit_log_preview,omitempty"`
	AccountAuditLogTaxId       *Many2One  `xmlrpc:"account_audit_log_tax_id,omitempty" json:"account_audit_log_tax_id,omitempty"`
	AttachmentIds              *Relation  `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	AuthorAvatar               *String    `xmlrpc:"author_avatar,omitempty" json:"author_avatar,omitempty"`
	AuthorGuestId              *Many2One  `xmlrpc:"author_guest_id,omitempty" json:"author_guest_id,omitempty"`
	AuthorId                   *Many2One  `xmlrpc:"author_id,omitempty" json:"author_id,omitempty"`
	Body                       *String    `xmlrpc:"body,omitempty" json:"body,omitempty"`
	ChildIds                   *Relation  `xmlrpc:"child_ids,omitempty" json:"child_ids,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Date                       *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	Description                *String    `xmlrpc:"description,omitempty" json:"description,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmailAddSignature          *Bool      `xmlrpc:"email_add_signature,omitempty" json:"email_add_signature,omitempty"`
	EmailFrom                  *String    `xmlrpc:"email_from,omitempty" json:"email_from,omitempty"`
	EmailLayoutXmlid           *String    `xmlrpc:"email_layout_xmlid,omitempty" json:"email_layout_xmlid,omitempty"`
	HasError                   *Bool      `xmlrpc:"has_error,omitempty" json:"has_error,omitempty"`
	HasSmsError                *Bool      `xmlrpc:"has_sms_error,omitempty" json:"has_sms_error,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsCurrentUserOrGuestAuthor *Bool      `xmlrpc:"is_current_user_or_guest_author,omitempty" json:"is_current_user_or_guest_author,omitempty"`
	IsInternal                 *Bool      `xmlrpc:"is_internal,omitempty" json:"is_internal,omitempty"`
	LetterIds                  *Relation  `xmlrpc:"letter_ids,omitempty" json:"letter_ids,omitempty"`
	LinkPreviewIds             *Relation  `xmlrpc:"link_preview_ids,omitempty" json:"link_preview_ids,omitempty"`
	MailActivityTypeId         *Many2One  `xmlrpc:"mail_activity_type_id,omitempty" json:"mail_activity_type_id,omitempty"`
	MailIds                    *Relation  `xmlrpc:"mail_ids,omitempty" json:"mail_ids,omitempty"`
	MailServerId               *Many2One  `xmlrpc:"mail_server_id,omitempty" json:"mail_server_id,omitempty"`
	MessageId                  *String    `xmlrpc:"message_id,omitempty" json:"message_id,omitempty"`
	MessageType                *Selection `xmlrpc:"message_type,omitempty" json:"message_type,omitempty"`
	Model                      *String    `xmlrpc:"model,omitempty" json:"model,omitempty"`
	Needaction                 *Bool      `xmlrpc:"needaction,omitempty" json:"needaction,omitempty"`
	NotificationIds            *Relation  `xmlrpc:"notification_ids,omitempty" json:"notification_ids,omitempty"`
	NotifiedPartnerIds         *Relation  `xmlrpc:"notified_partner_ids,omitempty" json:"notified_partner_ids,omitempty"`
	ParentId                   *Many2One  `xmlrpc:"parent_id,omitempty" json:"parent_id,omitempty"`
	PartnerIds                 *Relation  `xmlrpc:"partner_ids,omitempty" json:"partner_ids,omitempty"`
	PinnedAt                   *Time      `xmlrpc:"pinned_at,omitempty" json:"pinned_at,omitempty"`
	Preview                    *String    `xmlrpc:"preview,omitempty" json:"preview,omitempty"`
	RatingIds                  *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RatingValue                *Float     `xmlrpc:"rating_value,omitempty" json:"rating_value,omitempty"`
	ReactionIds                *Relation  `xmlrpc:"reaction_ids,omitempty" json:"reaction_ids,omitempty"`
	RecordAliasDomainId        *Many2One  `xmlrpc:"record_alias_domain_id,omitempty" json:"record_alias_domain_id,omitempty"`
	RecordCompanyId            *Many2One  `xmlrpc:"record_company_id,omitempty" json:"record_company_id,omitempty"`
	RecordName                 *String    `xmlrpc:"record_name,omitempty" json:"record_name,omitempty"`
	ReplyTo                    *String    `xmlrpc:"reply_to,omitempty" json:"reply_to,omitempty"`
	ReplyToForceNew            *Bool      `xmlrpc:"reply_to_force_new,omitempty" json:"reply_to_force_new,omitempty"`
	ResId                      *Many2One  `xmlrpc:"res_id,omitempty" json:"res_id,omitempty"`
	ShowAuditLog               *Bool      `xmlrpc:"show_audit_log,omitempty" json:"show_audit_log,omitempty"`
	SnailmailError             *Bool      `xmlrpc:"snailmail_error,omitempty" json:"snailmail_error,omitempty"`
	Starred                    *Bool      `xmlrpc:"starred,omitempty" json:"starred,omitempty"`
	StarredPartnerIds          *Relation  `xmlrpc:"starred_partner_ids,omitempty" json:"starred_partner_ids,omitempty"`
	Subject                    *String    `xmlrpc:"subject,omitempty" json:"subject,omitempty"`
	SubtypeId                  *Many2One  `xmlrpc:"subtype_id,omitempty" json:"subtype_id,omitempty"`
	TrackingValueIds           *Relation  `xmlrpc:"tracking_value_ids,omitempty" json:"tracking_value_ids,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// MailMessages represents array of mail.message model.
type MailMessages []MailMessage

// MailMessageModel is the odoo model name.
const MailMessageModel = "mail.message"

// Many2One convert MailMessage to *Many2One.
func (mm *MailMessage) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMessage creates a new mail.message model and returns its id.
func (c *Client) CreateMailMessage(mm *MailMessage) (int64, error) {
	ids, err := c.CreateMailMessages([]*MailMessage{mm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMessage creates a new mail.message model and returns its id.
func (c *Client) CreateMailMessages(mms []*MailMessage) ([]int64, error) {
	var vv []interface{}
	for _, v := range mms {
		vv = append(vv, v)
	}
	return c.Create(MailMessageModel, vv, nil)
}

// UpdateMailMessage updates an existing mail.message record.
func (c *Client) UpdateMailMessage(mm *MailMessage) error {
	return c.UpdateMailMessages([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMessages updates existing mail.message records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMessages(ids []int64, mm *MailMessage) error {
	return c.Update(MailMessageModel, ids, mm, nil)
}

// DeleteMailMessage deletes an existing mail.message record.
func (c *Client) DeleteMailMessage(id int64) error {
	return c.DeleteMailMessages([]int64{id})
}

// DeleteMailMessages deletes existing mail.message records.
func (c *Client) DeleteMailMessages(ids []int64) error {
	return c.Delete(MailMessageModel, ids)
}

// GetMailMessage gets mail.message existing record.
func (c *Client) GetMailMessage(id int64) (*MailMessage, error) {
	mms, err := c.GetMailMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	if len(*mms) == 0 {
		return nil, nil
	}
	return &((*mms)[0]), nil
}

// GetMailMessages gets mail.message existing records.
func (c *Client) GetMailMessages(ids []int64) (*MailMessages, error) {
	mms := &MailMessages{}
	if err := c.Read(MailMessageModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMessage finds mail.message record by querying it with criteria.
func (c *Client) FindMailMessage(criteria *Criteria) (*MailMessage, error) {
	mms := &MailMessages{}
	if err := c.SearchRead(MailMessageModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if len(*mms) == 0 {
		return nil, nil
	}
	return &((*mms)[0]), nil
}

// FindMailMessages finds mail.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessages(criteria *Criteria, options *Options) (*MailMessages, error) {
	mms := &MailMessages{}
	if err := c.SearchRead(MailMessageModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMessageModel, criteria, options)
}

// FindMailMessageId finds record id by querying it with criteria.
func (c *Client) FindMailMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}
