package odoo

// MailComposeMessage represents mail.compose.message model.
type MailComposeMessage struct {
	AttachmentIds        *Relation  `xmlrpc:"attachment_ids,omitempty" json:"attachment_ids,omitempty"`
	AuthorId             *Many2One  `xmlrpc:"author_id,omitempty" json:"author_id,omitempty"`
	AutoDelete           *Bool      `xmlrpc:"auto_delete,omitempty" json:"auto_delete,omitempty"`
	AutoDeleteKeepLog    *Bool      `xmlrpc:"auto_delete_keep_log,omitempty" json:"auto_delete_keep_log,omitempty"`
	Body                 *String    `xmlrpc:"body,omitempty" json:"body,omitempty"`
	BodyHasTemplateValue *Bool      `xmlrpc:"body_has_template_value,omitempty" json:"body_has_template_value,omitempty"`
	CanEditBody          *Bool      `xmlrpc:"can_edit_body,omitempty" json:"can_edit_body,omitempty"`
	CompositionBatch     *Bool      `xmlrpc:"composition_batch,omitempty" json:"composition_batch,omitempty"`
	CompositionMode      *Selection `xmlrpc:"composition_mode,omitempty" json:"composition_mode,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty" json:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty" json:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty" json:"display_name,omitempty"`
	EmailAddSignature    *Bool      `xmlrpc:"email_add_signature,omitempty" json:"email_add_signature,omitempty"`
	EmailFrom            *String    `xmlrpc:"email_from,omitempty" json:"email_from,omitempty"`
	EmailLayoutXmlid     *String    `xmlrpc:"email_layout_xmlid,omitempty" json:"email_layout_xmlid,omitempty"`
	ForceSend            *Bool      `xmlrpc:"force_send,omitempty" json:"force_send,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty" json:"id,omitempty"`
	IsMailTemplateEditor *Bool      `xmlrpc:"is_mail_template_editor,omitempty" json:"is_mail_template_editor,omitempty"`
	Lang                 *String    `xmlrpc:"lang,omitempty" json:"lang,omitempty"`
	MailActivityTypeId   *Many2One  `xmlrpc:"mail_activity_type_id,omitempty" json:"mail_activity_type_id,omitempty"`
	MailServerId         *Many2One  `xmlrpc:"mail_server_id,omitempty" json:"mail_server_id,omitempty"`
	MessageType          *Selection `xmlrpc:"message_type,omitempty" json:"message_type,omitempty"`
	Model                *String    `xmlrpc:"model,omitempty" json:"model,omitempty"`
	ModelIsThread        *Bool      `xmlrpc:"model_is_thread,omitempty" json:"model_is_thread,omitempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omitempty" json:"parent_id,omitempty"`
	PartnerIds           *Relation  `xmlrpc:"partner_ids,omitempty" json:"partner_ids,omitempty"`
	RecordAliasDomainId  *Many2One  `xmlrpc:"record_alias_domain_id,omitempty" json:"record_alias_domain_id,omitempty"`
	RecordCompanyId      *Many2One  `xmlrpc:"record_company_id,omitempty" json:"record_company_id,omitempty"`
	RecordName           *String    `xmlrpc:"record_name,omitempty" json:"record_name,omitempty"`
	RenderModel          *String    `xmlrpc:"render_model,omitempty" json:"render_model,omitempty"`
	ReplyTo              *String    `xmlrpc:"reply_to,omitempty" json:"reply_to,omitempty"`
	ReplyToForceNew      *Bool      `xmlrpc:"reply_to_force_new,omitempty" json:"reply_to_force_new,omitempty"`
	ReplyToMode          *Selection `xmlrpc:"reply_to_mode,omitempty" json:"reply_to_mode,omitempty"`
	ResDomain            *String    `xmlrpc:"res_domain,omitempty" json:"res_domain,omitempty"`
	ResDomainUserId      *Many2One  `xmlrpc:"res_domain_user_id,omitempty" json:"res_domain_user_id,omitempty"`
	ResIds               *String    `xmlrpc:"res_ids,omitempty" json:"res_ids,omitempty"`
	ScheduledDate        *String    `xmlrpc:"scheduled_date,omitempty" json:"scheduled_date,omitempty"`
	Subject              *String    `xmlrpc:"subject,omitempty" json:"subject,omitempty"`
	SubtypeId            *Many2One  `xmlrpc:"subtype_id,omitempty" json:"subtype_id,omitempty"`
	SubtypeIsLog         *Bool      `xmlrpc:"subtype_is_log,omitempty" json:"subtype_is_log,omitempty"`
	TemplateId           *Many2One  `xmlrpc:"template_id,omitempty" json:"template_id,omitempty"`
	TemplateName         *String    `xmlrpc:"template_name,omitempty" json:"template_name,omitempty"`
	UseExclusionList     *Bool      `xmlrpc:"use_exclusion_list,omitempty" json:"use_exclusion_list,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty" json:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty" json:"write_uid,omitempty"`
}

// MailComposeMessages represents array of mail.compose.message model.
type MailComposeMessages []MailComposeMessage

// MailComposeMessageModel is the odoo model name.
const MailComposeMessageModel = "mail.compose.message"

// Many2One convert MailComposeMessage to *Many2One.
func (mcm *MailComposeMessage) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailComposeMessage creates a new mail.compose.message model and returns its id.
func (c *Client) CreateMailComposeMessage(mcm *MailComposeMessage) (int64, error) {
	ids, err := c.CreateMailComposeMessages([]*MailComposeMessage{mcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailComposeMessage creates a new mail.compose.message model and returns its id.
func (c *Client) CreateMailComposeMessages(mcms []*MailComposeMessage) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcms {
		vv = append(vv, v)
	}
	return c.Create(MailComposeMessageModel, vv, nil)
}

// UpdateMailComposeMessage updates an existing mail.compose.message record.
func (c *Client) UpdateMailComposeMessage(mcm *MailComposeMessage) error {
	return c.UpdateMailComposeMessages([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailComposeMessages updates existing mail.compose.message records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailComposeMessages(ids []int64, mcm *MailComposeMessage) error {
	return c.Update(MailComposeMessageModel, ids, mcm, nil)
}

// DeleteMailComposeMessage deletes an existing mail.compose.message record.
func (c *Client) DeleteMailComposeMessage(id int64) error {
	return c.DeleteMailComposeMessages([]int64{id})
}

// DeleteMailComposeMessages deletes existing mail.compose.message records.
func (c *Client) DeleteMailComposeMessages(ids []int64) error {
	return c.Delete(MailComposeMessageModel, ids)
}

// GetMailComposeMessage gets mail.compose.message existing record.
func (c *Client) GetMailComposeMessage(id int64) (*MailComposeMessage, error) {
	mcms, err := c.GetMailComposeMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// GetMailComposeMessages gets mail.compose.message existing records.
func (c *Client) GetMailComposeMessages(ids []int64) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.Read(MailComposeMessageModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessage finds mail.compose.message record by querying it with criteria.
func (c *Client) FindMailComposeMessage(criteria *Criteria) (*MailComposeMessage, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// FindMailComposeMessages finds mail.compose.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessages(criteria *Criteria, options *Options) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailComposeMessageModel, criteria, options)
}

// FindMailComposeMessageId finds record id by querying it with criteria.
func (c *Client) FindMailComposeMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailComposeMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
