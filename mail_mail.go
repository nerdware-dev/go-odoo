package odoo

// MailMail represents mail.mail model.
type MailMail struct {
	AccountAuditLogAccountId   *Many2One  `xmlrpc:"account_audit_log_account_id,omitempty" json:"account_audit_log_account_id,omitempty"`
	AccountAuditLogActivated   *Bool      `xmlrpc:"account_audit_log_activated,omitempty" json:"account_audit_log_activated,omitempty"`
	AccountAuditLogCompanyId   *Many2One  `xmlrpc:"account_audit_log_company_id,omitempty" json:"account_audit_log_company_id,omitempty"`
	AccountAuditLogMoveId      *Many2One  `xmlrpc:"account_audit_log_move_id,omitempty" json:"account_audit_log_move_id,omitempty"`
	AccountAuditLogPartnerId   *Many2One  `xmlrpc:"account_audit_log_partner_id,omitempty" json:"account_audit_log_partner_id,omitempty"`
	AccountAuditLogPreview     *String    `xmlrpc:"account_audit_log_preview,omitempty" json:"account_audit_log_preview,omitempty"`
	AccountAuditLogTaxId       *Many2One  `xmlrpc:"account_audit_log_tax_id,omitempty" json:"account_audit_log_tax_id,omitempty"`
	AttachmentIds              *Relation  `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	AuthorAvatar               *String    `xmlrpc:"author_avatar,omitempty" json:"author_avatar,omitempty"`
	AuthorGuestId              *Many2One  `xmlrpc:"author_guest_id,omitempty" json:"author_guest_id,omitempty"`
	AuthorId                   *Many2One  `xmlrpc:"author_id,omitempty" json:"author_id,omitempty"`
	AutoDelete                 *Bool      `xmlrpc:"auto_delete,omitempty" json:"auto_delete,omitempty"`
	Body                       *String    `xmlrpc:"body,omitempty" json:"body,omitempty"`
	BodyContent                *String    `xmlrpc:"body_content,omitempty" json:"body_content,omitempty"`
	BodyHtml                   *String    `xmlrpc:"body_html,omitempty" json:"body_html,omitempty"`
	ChildIds                   *Relation  `xmlrpc:"child_ids,omitempty" json:"child_ids,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	Date                       *Time      `xmlrpc:"date,omitempty" json:"date,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmailAddSignature          *Bool      `xmlrpc:"email_add_signature,omitempty" json:"email_add_signature,omitempty"`
	EmailCc                    *String    `xmlrpc:"email_cc,omitempty" json:"email_cc,omitempty"`
	EmailFrom                  *String    `xmlrpc:"email_from,omitempty" json:"email_from,omitempty"`
	EmailLayoutXmlid           *String    `xmlrpc:"email_layout_xmlid,omitempty" json:"email_layout_xmlid,omitempty"`
	EmailTo                    *String    `xmlrpc:"email_to,omitempty" json:"email_to,omitempty"`
	FailureReason              *String    `xmlrpc:"failure_reason,omitempty" json:"failure_reason,omitempty"`
	FailureType                *Selection `xmlrpc:"failure_type,omitempty" json:"failure_type,omitempty"`
	FetchmailServerId          *Many2One  `xmlrpc:"fetchmail_server_id,omitempty" json:"fetchmail_server_id,omitempty"`
	HasError                   *Bool      `xmlrpc:"has_error,omitempty" json:"has_error,omitempty"`
	HasSmsError                *Bool      `xmlrpc:"has_sms_error,omitempty" json:"has_sms_error,omitempty"`
	Headers                    *String    `xmlrpc:"headers,omitempty" json:"headers,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsCurrentUserOrGuestAuthor *Bool      `xmlrpc:"is_current_user_or_guest_author,omitempty" json:"is_current_user_or_guest_author,omitempty"`
	IsInternal                 *Bool      `xmlrpc:"is_internal,omitempty" json:"is_internal,omitempty"`
	IsNotification             *Bool      `xmlrpc:"is_notification,omitempty" json:"is_notification,omitempty"`
	LetterIds                  *Relation  `xmlrpc:"letter_ids,omitempty" json:"letter_ids,omitempty"`
	LinkPreviewIds             *Relation  `xmlrpc:"link_preview_ids,omitempty" json:"link_preview_ids,omitempty"`
	MailActivityTypeId         *Many2One  `xmlrpc:"mail_activity_type_id,omitempty" json:"mail_activity_type_id,omitempty"`
	MailIds                    *Relation  `xmlrpc:"mail_ids,omitempty" json:"mail_ids,omitempty"`
	MailMessageId              *Many2One  `xmlrpc:"mail_message_id,omitempty" json:"mail_message_id,omitempty"`
	MailMessageIdInt           *Int       `xmlrpc:"mail_message_id_int,omitempty" json:"mail_message_id_int,omitempty"`
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
	RatingId                   *Many2One  `xmlrpc:"rating_id,omitempty" json:"rating_id,omitempty"`
	RatingIds                  *Relation  `xmlrpc:"rating_ids,omitempty" json:"rating_ids,omitempty"`
	RatingValue                *Float     `xmlrpc:"rating_value,omitempty" json:"rating_value,omitempty"`
	ReactionIds                *Relation  `xmlrpc:"reaction_ids,omitempty" json:"reaction_ids,omitempty"`
	RecipientIds               *Relation  `xmlrpc:"recipient_ids,omitempty" json:"recipient_ids,omitempty"`
	RecordAliasDomainId        *Many2One  `xmlrpc:"record_alias_domain_id,omitempty" json:"record_alias_domain_id,omitempty"`
	RecordCompanyId            *Many2One  `xmlrpc:"record_company_id,omitempty" json:"record_company_id,omitempty"`
	RecordName                 *String    `xmlrpc:"record_name,omitempty" json:"record_name,omitempty"`
	References                 *String    `xmlrpc:"references,omitempty" json:"references,omitempty"`
	ReplyTo                    *String    `xmlrpc:"reply_to,omitempty" json:"reply_to,omitempty"`
	ReplyToForceNew            *Bool      `xmlrpc:"reply_to_force_new,omitempty" json:"reply_to_force_new,omitempty"`
	ResId                      *Many2One  `xmlrpc:"res_id,omitempty" json:"res_id,omitempty"`
	RestrictedAttachmentCount  *Int       `xmlrpc:"restricted_attachment_count,omitempty" json:"restricted_attachment_count,omitempty"`
	ScheduledDate              *Time      `xmlrpc:"scheduled_date,omitempty" json:"scheduled_date,omitempty"`
	SnailmailError             *Bool      `xmlrpc:"snailmail_error,omitempty" json:"snailmail_error,omitempty"`
	Starred                    *Bool      `xmlrpc:"starred,omitempty" json:"starred,omitempty"`
	StarredPartnerIds          *Relation  `xmlrpc:"starred_partner_ids,omitempty" json:"starred_partner_ids,omitempty"`
	State                      *Selection `xmlrpc:"state,omitempty" json:"state,omitempty"`
	Subject                    *String    `xmlrpc:"subject,omitempty" json:"subject,omitempty"`
	SubtypeId                  *Many2One  `xmlrpc:"subtype_id,omitempty" json:"subtype_id,omitempty"`
	TrackingValueIds           *Relation  `xmlrpc:"tracking_value_ids,omitempty" json:"tracking_value_ids,omitempty"`
	UnrestrictedAttachmentIds  *Relation  `xmlrpc:"unrestricted_attachment_ids,omitempty" json:"unrestricted_attachment_ids,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// MailMails represents array of mail.mail model.
type MailMails []MailMail

// MailMailModel is the odoo model name.
const MailMailModel = "mail.mail"

// Many2One convert MailMail to *Many2One.
func (mm *MailMail) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMail creates a new mail.mail model and returns its id.
func (c *Client) CreateMailMail(mm *MailMail) (int64, error) {
	ids, err := c.CreateMailMails([]*MailMail{mm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMail creates a new mail.mail model and returns its id.
func (c *Client) CreateMailMails(mms []*MailMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range mms {
		vv = append(vv, v)
	}
	return c.Create(MailMailModel, vv, nil)
}

// UpdateMailMail updates an existing mail.mail record.
func (c *Client) UpdateMailMail(mm *MailMail) error {
	return c.UpdateMailMails([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMails updates existing mail.mail records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMails(ids []int64, mm *MailMail) error {
	return c.Update(MailMailModel, ids, mm, nil)
}

// DeleteMailMail deletes an existing mail.mail record.
func (c *Client) DeleteMailMail(id int64) error {
	return c.DeleteMailMails([]int64{id})
}

// DeleteMailMails deletes existing mail.mail records.
func (c *Client) DeleteMailMails(ids []int64) error {
	return c.Delete(MailMailModel, ids)
}

// GetMailMail gets mail.mail existing record.
func (c *Client) GetMailMail(id int64) (*MailMail, error) {
	mms, err := c.GetMailMails([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mms)[0]), nil
}

// GetMailMails gets mail.mail existing records.
func (c *Client) GetMailMails(ids []int64) (*MailMails, error) {
	mms := &MailMails{}
	if err := c.Read(MailMailModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMail finds mail.mail record by querying it with criteria.
func (c *Client) FindMailMail(criteria *Criteria) (*MailMail, error) {
	mms := &MailMails{}
	if err := c.SearchRead(MailMailModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	return &((*mms)[0]), nil
}

// FindMailMails finds mail.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMails(criteria *Criteria, options *Options) (*MailMails, error) {
	mms := &MailMails{}
	if err := c.SearchRead(MailMailModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMailModel, criteria, options)
}

// FindMailMailId finds record id by querying it with criteria.
func (c *Client) FindMailMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
